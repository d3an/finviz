// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package quote

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/cenkalti/backoff/v4"
	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-gota/gota/dataframe"

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

func (c *Client) RandomizeUserAgent() {
	c.config.userAgent = uarand.GetRandom()
}

func GenerateURL(ticker string) (string, error) {
	return fmt.Sprintf("%s?t=%s&ty=c&p=d&b=1", APIURL, strings.ToUpper(ticker)), nil
}

type response struct {
	Result  *map[string]interface{}
	Warning error
	Error   error
}

type Warning struct {
	Ticker string
	Error  error
}

type Error struct {
	Ticker string
	Error  error
}

type Results struct {
	Data     *dataframe.DataFrame
	Warnings []Warning
	Errors   []Error
}

func (c *Client) GetQuotes(tickers []string) (finalResults Results, err error) {
	var wg sync.WaitGroup
	resultCount := len(tickers)
	rawResults := make([]chan response, resultCount)
	for i := range rawResults {
		rawResults[i] = make(chan response, 1)
	}

	for i, ticker := range tickers {
		wg.Add(1)
		go c.getData(ticker, &wg, &rawResults[i])
		wg.Wait()
	}

	var scrapedResults []map[string]interface{}
	for i := 0; i < resultCount; i++ {
		r := <-rawResults[i]
		if r.Warning != nil {
			finalResults.Warnings = append(finalResults.Warnings, Warning{Ticker: tickers[i], Error: r.Warning})
			continue
		}
		if r.Error != nil {
			finalResults.Errors = append(finalResults.Errors, Error{Ticker: tickers[i], Error: r.Error})
			continue
		}
		scrapedResults = append(scrapedResults, *r.Result)
	}

	finalResults.Data, err = processScrapeResults(scrapedResults)
	return
}

func processScrapeResults(results []map[string]interface{}) (*dataframe.DataFrame, error) {
	quoteHeaders := []string{"Ticker", "Company", "Industry", "Sector", "Country", "Exchange", "Index", "Market Cap", "Price", "Change", "Volume", "Income", "Sales", "Book/sh", "Cash/sh", "Dividend", "Dividend %", "Employees", "Optionable", "Shortable", "Recom", "P/E", "Forward P/E", "PEG", "P/S", "P/B", "P/C", "P/FCF", "Quick Ratio", "Current Ratio", "Debt/Eq", "LT Debt/Eq", "EPS (ttm)", "EPS next Y", "EPS next Q", "EPS this Y", "EPS growth next Y", "EPS next 5Y", "EPS past 5Y", "Sales past 5Y", "Sales Q/Q", "EPS Q/Q", "Earnings", "Insider Own", "Insider Trans", "Inst Own", "Inst Trans", "ROA", "ROE", "ROI", "Gross Margin", "Oper. Margin", "Profit Margin", "Payout", "Shs Outstand", "Shs Float", "Short Float", "Short Ratio", "Target Price", "52W Range", "52W High", "52W Low", "RSI (14)", "SMA20", "SMA50", "SMA200", "Rel Volume", "Avg Volume", "Perf Week", "Perf Month", "Perf Quarter", "Perf Half Y", "Perf Year", "Perf YTD", "Beta", "ATR", "Volatility (Week)", "Volatility (Month)", "Prev Close", "Analyst Recommendations", "News", "Description", "Insider Trading"}
	orderedRows, err := utils.GenerateRows(quoteHeaders, results)
	if err != nil {
		return nil, fmt.Errorf("error failed to generate rows from quote KVP map")
	}
	df := dataframe.LoadRecords(orderedRows)
	return utils.CleanFinvizDataFrame(&df), nil
}

func (c *Client) getData(ticker string, wg *sync.WaitGroup, result *chan response) {
	defer wg.Done()
	defer close(*result)

	url, err := GenerateURL(ticker)
	if err != nil {
		*result <- response{Error: err}
		return
	}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		*result <- response{Error: err}
		return
	}

	var body []byte
	var warning error
	if err = backoff.RetryNotify(func() error {
		resp, err := c.Do(req)
		if err != nil {
			return backoff.Permanent(err)
		}
		defer resp.Body.Close()

		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return backoff.Permanent(err)
		}

		if resp.StatusCode == http.StatusNotFound {
			warning = fmt.Errorf("resource not found")
			return nil
		} else if resp.StatusCode == http.StatusForbidden && string(body) == "error code: 1010" {
			c.RandomizeUserAgent()
			return fmt.Errorf("blocked by Cloudflare, status code: '%d', body: '%s'", resp.StatusCode, string(body))
		} else if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to get url: '%s', status code: '%d', body: '%s'", url, resp.StatusCode, string(body))
		}

		if string(body) == "Too many requests." {
			return fmt.Errorf("request rate limit reached")
		}
		return nil
	}, backoff.NewExponentialBackOff(), func(err error, td time.Duration) {
		fmt.Printf("[ERROR]: %v\n", err)
		fmt.Printf("[WAIT_IN_SECONDS]: %v\n", 2*td.Seconds())
		time.Sleep(td)
	}); err != nil {
		*result <- response{Error: err}
		return
	} else if warning != nil {
		*result <- response{Warning: warning}
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
	data := make(map[string]interface{})
	doc.Find("tr[class=\"table-dark-row\"] > td").Each(func(column int, row *goquery.Selection) {
		if column%2 == 0 {
			switch row.Text() {
			default:
				data[row.Text()] = row.Next().Text()
			case "Index":
				if row.Next().Text() == "S&P 500" {
					data["Index"] = "S&P500"
				} else {
					data["Index"] = strings.Join(strings.Split(row.Next().Text(), " "), ",")
				}
			case "EPS next Y":
				if _, exists := data["EPS next Y"]; exists {
					data["EPS growth next Y"] = row.Next().Text()
				} else {
					data["EPS next Y"] = row.Next().Text()
				}
			case "Volatility":
				vols := strings.Split(row.Next().Text(), " ")
				data["Volatility (Week)"] = vols[0]
				data["Volatility (Month)"] = vols[1]
			}
		}
	})

	// Title Data
	mainSection := doc.Find(".fullview-title > tbody > tr")
	data["Ticker"] = doc.Find("#ticker").Text()
	exchange := doc.Find("#ticker").Next().Text()
	data["Exchange"] = exchange[1 : len(exchange)-1]
	data["Company"] = mainSection.Eq(1).Text()
	data["Sector"] = mainSection.Eq(2).Find("a").Eq(0).Text()
	data["Industry"] = mainSection.Eq(2).Find("a").Eq(1).Text()
	data["Country"] = mainSection.Eq(2).Find("a").Eq(2).Text()

	extraSection := doc.Find("body > table").Eq(2).Find("tbody").Eq(0).ChildrenFiltered("tr")

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
		data["Analyst Recommendations"] = analystRecommendations
	} else {
		data["Analyst Recommendations"] = ""
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
	data["News"] = news

	// Description
	if len(doc.Find(".fullview-profile").Nodes) == 1 {
		data["Description"] = doc.Find(".fullview-profile").Text()
	} else {
		data["Description"] = ""
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
	// insiderData := extraSection.Eq(len(extraSection.Nodes) - 4)
	insiderData := doc.Find(".insider-sale-row-2")
	if len(insiderData.Nodes) > 0 {
		var insiderTrading []map[string]string
		insiderData.Parent().Children().Each(func(i int, rowNode *goquery.Selection) {
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
		data["Insider Trading"] = insiderTrading
	} else {
		data["Insider Trading"] = ""
	}

	return &data, nil
}
