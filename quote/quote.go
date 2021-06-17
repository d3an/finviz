package quote

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-gota/gota/dataframe"
	"github.com/pkg/errors"

	"github.com/d3an/finviz/utils"
)

const (
	APIURL = "https://finviz.com/quote.ashx"
)

var (
	once     sync.Once
	instance *Client
)

type Config struct {
	userAgent string
	recorder  *recorder.Recorder
}

type Client struct {
	*http.Client
	config Config
}

func New(config *Config) *Client {
	once.Do(func() {
		transport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 30 * time.Second,
		}
		client := &http.Client{
			Timeout:   30 * time.Second,
			Transport: transport,
		}
		if config != nil {
			instance = &Client{Client: client, config: *config}
		}
		instance = &Client{
			Client: client,
			config: Config{userAgent: uarand.GetRandom()},
		}
	})

	return instance
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", c.config.userAgent)
	return c.Client.Do(req)
}

func GenerateURL(ticker string) (string, error) {
	return fmt.Sprintf("%s?t=%s&ty=c&p=d&b=1", APIURL, strings.ToUpper(ticker)), nil
}

type response struct {
	Result *map[string]interface{}
	Error  error
}

func (c *Client) GetQuotes(tickers []string) (*dataframe.DataFrame, error) {
	var wg sync.WaitGroup
	resultCount := len(tickers)
	results := make([]chan response, resultCount)
	for i := range results {
		results[i] = make(chan response, 1)
	}

	for i, ticker := range tickers {
		wg.Add(1)
		go c.getData(ticker, &wg, &results[i])
		wg.Wait()
	}

	var scrapeResults []map[string]interface{}
	for i := 0; i < resultCount; i++ {
		r := <-results[i]
		if r.Error != nil {
			return nil, errors.Wrapf(r.Error, "error received while scraping quote for '%s'", tickers[i])
		}
		scrapeResults = append(scrapeResults, *r.Result)
	}

	return processScrapeResults(scrapeResults)
}

func processScrapeResults(results []map[string]interface{}) (*dataframe.DataFrame, error) {
	quoteHeaders := []string{"Ticker", "Company", "Industry", "Sector", "Country", "Index", "Market Cap", "Price", "Change", "Volume", "Income", "Sales", "Book/sh", "Cash/sh", "Dividend", "Dividend %", "Employees", "Optionable", "Shortable", "Recom", "P/E", "Forward P/E", "PEG", "P/S", "P/B", "P/C", "P/FCF", "Quick Ratio", "Current Ratio", "Debt/Eq", "LT Debt/Eq", "EPS (ttm)", "EPS next Y", "EPS next Q", "EPS this Y", "EPS growth next Y", "EPS next 5Y", "EPS past 5Y", "Sales past 5Y", "Sales Q/Q", "EPS Q/Q", "Earnings", "Insider Own", "Insider Trans", "Inst Own", "Inst Trans", "ROA", "ROE", "ROI", "Gross Margin", "Oper. Margin", "Profit Margin", "Payout", "Shs Outstand", "Shs Float", "Short Float", "Short Ratio", "Target Price", "52W Range", "52W High", "52W Low", "RSI (14)", "SMA20", "SMA50", "SMA200", "Rel Volume", "Avg Volume", "Perf Week", "Perf Month", "Perf Quarter", "Perf Half Y", "Perf Year", "Perf YTD", "Beta", "ATR", "Volatility (Week)", "Volatility (Month)", "Prev Close", "Chart", "Analyst Recommendations", "News", "Description", "Insider Trading"}
	orderedRows, err := utils.GenerateRows(quoteHeaders, results)
	if err != nil {
		return nil, fmt.Errorf("error failed to generate rows from quote KVP map")
	}
	df := dataframe.LoadRecords(orderedRows)
	return &df, nil
}

func (c *Client) getData(ticker string, wg *sync.WaitGroup, result *chan response) {
	defer wg.Done()
	defer close(*result)

	url, err := GenerateURL(ticker)
	if err != nil {
		*result <- response{Error: err}
		return
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		*result <- response{Error: err}
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		*result <- response{Error: err}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		*result <- response{Error: err}
		return
	}

	if resp.StatusCode != http.StatusOK {
		*result <- response{Error: fmt.Errorf("error getting url: '%s', status code: '%d', body: '%s'", url, resp.StatusCode, string(body))}
		return
	}

	doc, err := utils.GenerateDocument(body)
	if err != nil {
		*result <- response{Error: err}
		return
	}

	dict, err := Scrape(doc)
	if err != nil {
		*result <- response{Error: err}
		return
	}

	*result <- response{Result: dict}
}

// Scrape scrapes FinViz views to a KVP map
func Scrape(doc *goquery.Document) (*map[string]interface{}, error) {
	basicData := doc.Find("body > table").Eq(2).Find("tbody").Eq(0)

	titleData := basicData.Children().Eq(5).Find("tbody").Eq(0).Children().Eq(1).Find("tbody").Eq(0)
	tableData := basicData.Children().Eq(6).Find("tbody").Eq(0)

	extraSection := doc.Find("body > table").Eq(3).Find("tbody").Eq(0).ChildrenFiltered("tr")

	rawTickerData := make(map[string]interface{})

	// Title Data
	rawTickerData["Chart"] = doc.Find("#chart0").Eq(0).AttrOr("src", "")
	rawTickerData["Ticker"] = doc.Find("#ticker").Eq(0).Text()
	rawTickerData["Company"] = titleData.Children().Eq(1).Find("b").Eq(0).Text()
	rawTickerData["Sector"] = titleData.Children().Eq(2).Find("a").Eq(0).Text()
	rawTickerData["Industry"] = titleData.Children().Eq(2).Find("a").Eq(1).Text()
	rawTickerData["Country"] = titleData.Children().Eq(2).Find("a").Eq(2).Text()

	// Table Data
	tableData.Children().Each(func(i int, rowNode *goquery.Selection) {
		var key string
		rowNode.Children().Each(func(j int, cellNode *goquery.Selection) {
			if j%2 == 0 {
				key = cellNode.Text()
			} else {
				if key == "Volatility" {
					data := strings.Split(cellNode.Find("small").Text(), " ")
					rawTickerData["Volatility (Week)"] = data[0]
					rawTickerData["Volatility (Month)"] = data[1]
				} else if key == "Index" {
					rawTickerData[key] = strings.Join(strings.Split(cellNode.Find("small").Text(), " "), ",")
				} else if key == "52W Range" {
					rawTickerData[key] = cellNode.Find("small").Text()
				} else if key == "EPS next Y" {
					if _, exists := rawTickerData[key]; exists {
						rawTickerData["EPS growth next Y"] = cellNode.Find("b").Text()
					} else {
						rawTickerData[key] = cellNode.Find("b").Text()
					}
				} else {
					rawTickerData[key] = cellNode.Find("b").Text()
				}
			}
		})
	})

	// Analyst Recommendations
	if len(extraSection.Eq(3).Find("#news-table").Nodes) == 0 {
		var analystRecommendations []map[string]string
		analystRecomData := doc.Find(".fullview-ratings-outer").Find("tbody").Eq(0)
		analystRecomData.Children().Each(func(i int, rowNode *goquery.Selection) {
			itemNodes := rowNode.Find("tbody").Find("tr").Children()
			analystRecommendations = append(analystRecommendations, map[string]string{
				"Date":         itemNodes.Eq(0).Text(),
				"Action":       itemNodes.Eq(1).Find("b").Text(),
				"Brokerage":    itemNodes.Eq(2).Text(),
				"Rating":       itemNodes.Eq(3).Text(),
				"Price Target": itemNodes.Eq(4).Text(),
			})
		})
		rawTickerData["Analyst Recommendations"] = analystRecommendations
	} else {
		rawTickerData["Analyst Recommendations"] = ""
	}

	// News
	newsData := doc.Find("#news-table").Find("tbody")
	var news []map[string]string
	var date string
	var datetime string
	var newsGain string
	newsData.Children().Each(func(i int, rowNode *goquery.Selection) {
		if rowNode.Find("td").Eq(0).AttrOr("style", "") != "" {
			date = rowNode.Find("td").Eq(0).Text()[0:9]
			datetime = rowNode.Find("td").Eq(0).Text()[0:17]
		} else {
			datetime = fmt.Sprintf("%s %s", date, rowNode.Find("td").Eq(0).Text()[0:7])
		}

		newsGain = ""
		if len(rowNode.Find("td").Eq(1).Find("span").Nodes) >= 2 {
			newsGain = rowNode.Find("td").Eq(1).Find("span").Eq(1).Text()
		}

		news = append(news, map[string]string{
			"Datetime":  datetime,
			"Link":      rowNode.Find("td").Eq(1).Find("a").AttrOr("href", ""),
			"Source":    rowNode.Find("td").Eq(1).Find("span").Eq(0).Text()[1:],
			"Title":     rowNode.Find("td").Eq(1).Find("a").Text(),
			"News Gain": newsGain,
		})
	})
	rawTickerData["News"] = news

	// Description
	if len(doc.Find(".fullview-profile").Nodes) == 1 {
		rawTickerData["Description"] = doc.Find(".fullview-profile").Text()
	} else {
		rawTickerData["Description"] = ""
	}

	// Income Statement
	// statement.ashx?t=____&s=IA
	// statement.ashx?t=____&s=IQ

	// Balance Sheet
	// statement.ashx?t=____&s=BA
	// statement.ashx?t=____&s=BQ

	// Cash Flow
	// statement.ashx?t=____&s=CA
	// statement.ashx?t=____&s=CQ

	// Insider Trading
	insiderData := extraSection.Eq(len(extraSection.Nodes) - 4)
	if len(insiderData.Find(".body-table").Nodes) > 0 {
		var insiderTrading []map[string]string
		insiderData.Find("tbody").Eq(0).Children().Each(func(i int, rowNode *goquery.Selection) {
			if i != 0 {
				insiderTrading = append(insiderTrading, map[string]string{
					"Owner":               rowNode.Children().Eq(0).Find("a").Text(),
					"Relationship":        rowNode.Children().Eq(1).Text(),
					"Date":                rowNode.Children().Eq(2).Text(),
					"Transaction":         rowNode.Children().Eq(3).Text(),
					"Cost":                rowNode.Children().Eq(4).Text(),
					"#Shares":             rowNode.Children().Eq(5).Text(),
					"Value ($)":           rowNode.Children().Eq(6).Text(),
					"#Shares Total":       rowNode.Children().Eq(7).Text(),
					"SEC Form 4 Datetime": rowNode.Children().Eq(8).Find("a").Text(),
					"SEC Form 4 Link":     rowNode.Children().Eq(8).Find("a").AttrOr("href", ""),
				})
			}
		})
		rawTickerData["Insider Trading"] = insiderTrading
	} else {
		rawTickerData["Insider Trading"] = ""
	}

	return &rawTickerData, nil
}
