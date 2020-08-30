// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"

	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// NewClient generates a new client instance
func NewClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}

/*
// newTestingClient generates a new testing client instance that uses go-vcr
func newTestingClient(r *recorder.Recorder) *http.Client {
	return &http.Client{
		Timeout:   30 * time.Second,
		Transport: r,
	}
}
*/

// MakeGetRequest is used to get a byte array of the screen given a valid URL
func MakeGetRequest(c *http.Client, url string) ([]byte, error) {
	// Set up GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Make GET request
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle unsuccessful GET requests
	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return b, StatusCodeError(fmt.Sprintf("Received HTTP response status %v: %v", resp.StatusCode, resp.Status))
	}

	// Convert the response body to a string
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return html, nil
}

func generateDocument(html interface{}) (doc *goquery.Document, err error) {
	switch html := html.(type) {
	default:
		return nil, fmt.Errorf("HTML object is not of type string or []byte or io.ReadCloser")
	case string:
		html = strings.ReplaceAll(html, "\\r", "")
		html = strings.ReplaceAll(html, "\\n", "")
		html = strings.ReplaceAll(html, "\\\"", "\"")

		html = strings.Map(func(r rune) rune {
			if r == '\n' || r == '\t' {
				return ' '
			}
			return r
		}, html)
		doc, err = goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
		if err != nil {
			return nil, err
		}
	case []byte:
		doc, err = goquery.NewDocumentFromReader(bytes.NewReader(html))
		if err != nil {
			return nil, err
		}
	case io.ReadCloser:
		byteArray, err := ioutil.ReadAll(html)
		if err != nil {
			return nil, err
		}
		return generateDocument(byteArray)
	}
	return doc, nil
}

/*
// basicDefaultViewScraper scrapes the 310 Screen document for the query's results
func basicDefaultViewScraper(html interface{}) (rows [][]string, err error) {
	// generate doc
	doc, err := generateDocument(html)
	if err != nil {
		return nil, err
	}

	// var container declaration
	var tickerDataSlice []map[string]interface{}
	var headers []string

	rootNode := doc.Find("div > table > tbody").Children()
	firstDataNodeIndex := 4
	lastDataNodeIndex := rootNode.Size() - 3

	rootNode.Slice(firstDataNodeIndex, lastDataNodeIndex).Each(func(i int, childNode *goquery.Selection) {
		var rawTickerData = make(map[string]interface{})
		headers, rawTickerData = basicDefaultViewHelper(childNode, headers, rawTickerData)
		tickerDataSlice = append(tickerDataSlice, rawTickerData)
	})

	return generateRowsHelper(headers, tickerDataSlice)
}

// basicNewsViewScraper scrapes the 320 Screen document for the query's results
func basicNewsViewScraper(html interface{}) (rows [][]string, err error) {
	// generate doc
	doc, err := generateDocument(html)
	if err != nil {
		return nil, err
	}

	// var container declaration
	var tickerDataSlice []map[string]interface{}
	var headers []string
	var rawTickerData map[string]interface{}

	rootNode := doc.Find("div > table > tbody").Children()
	firstDataNodeIndex := 4
	lastDataNodeIndex := rootNode.Size() - 3

	rootNode.Slice(firstDataNodeIndex, lastDataNodeIndex).Each(func(i int, childNode *goquery.Selection) {
		if i%3 == 0 {
			rawTickerData = make(map[string]interface{})
			headers, rawTickerData = basicDefaultViewHelper(childNode, headers, rawTickerData)

		} else if i%3 == 1 {
			headers, rawTickerData = basicNewsViewHelper(childNode, headers, rawTickerData)
			tickerDataSlice = append(tickerDataSlice, rawTickerData)
		}
	})

	return generateRowsHelper(headers, tickerDataSlice)
}

// basicDescriptionViewScraper scrapes the 330 Screen document for the query's results
func basicDescriptionViewScraper(html interface{}) (rows [][]string, err error) {
	// generate doc
	doc, err := generateDocument(html)
	if err != nil {
		return nil, err
	}

	// var container declaration
	var tickerDataSlice []map[string]interface{}
	var headers []string
	var rawTickerData map[string]interface{}

	rootNode := doc.Find("div > table > tbody").Children()
	firstDataNodeIndex := 4
	lastDataNodeIndex := rootNode.Size() - 3

	rootNode.Slice(firstDataNodeIndex, lastDataNodeIndex).Each(func(i int, childNode *goquery.Selection) {
		if i%3 == 0 {
			rawTickerData = make(map[string]interface{})
			headers, rawTickerData = basicDefaultViewHelper(childNode, headers, rawTickerData)
		} else if i%3 == 1 {
			headers, rawTickerData = basicDescriptionViewHelper(childNode, headers, rawTickerData)
			tickerDataSlice = append(tickerDataSlice, rawTickerData)
		}
	})

	return generateRowsHelper(headers, tickerDataSlice)
}

// basicSnapshotViewScraper scrapes the 340 Screen document for the query's results
func basicSnapshotViewScraper(html interface{}) (rows [][]string, err error) {
	// generate doc
	doc, err := generateDocument(html)
	if err != nil {
		return nil, err
	}

	// var container declaration
	var tickerDataSlice []map[string]interface{}
	var headers []string
	var rawTickerData map[string]interface{}

	rootNode := doc.Find("div > table > tbody").Children()
	firstDataNodeIndex := 4
	lastDataNodeIndex := rootNode.Size() - 3

	rootNode.Slice(firstDataNodeIndex, lastDataNodeIndex).Each(func(i int, childNode *goquery.Selection) {
		if i%3 == 0 {
			rawTickerData = make(map[string]interface{})
			headers, rawTickerData = basicDefaultViewHelper(childNode, headers, rawTickerData)
		} else if i%3 == 1 {
			snapshotNodeLen := childNode.Find("td > table > tbody").Size()
			for j := 0; j < snapshotNodeLen; j++ {
				if j == 0 {
					headers, rawTickerData = basicNewsViewHelper(childNode, headers, rawTickerData)
				} else if j == 1 {
					headers, rawTickerData = basicDescriptionViewHelper(childNode, headers, rawTickerData)
				} else if j == 2 {
					headers, rawTickerData = basicInsiderTradingViewHelper(childNode, headers, rawTickerData)
				}
			}
			tickerDataSlice = append(tickerDataSlice, rawTickerData)
		}
	})

	return generateRowsHelper(headers, tickerDataSlice)
}

// tickersViewScraper scrapes a 410 Screen document for the screen's ticker results (max 1000 tickers)
func tickersViewScraper(html interface{}) (rows [][]string, err error) {
	// generate doc
	doc, err := generateDocument(html)
	if err != nil {
		return nil, err
	}

	var tickerDataSlice []map[string]interface{}
	var rawTickerData map[string]interface{}

	doc.Find("#screener-content").Find("tbody").Eq(3).Find("tr").Children().Children().Each(func(i int, tickerNode *goquery.Selection) {
		rawTickerData = make(map[string]interface{})
		if titleAttr := tickerNode.AttrOr("title", ""); titleAttr != "" {
			body := strings.Split(strings.Split(titleAttr, "body=[")[2], "]")[0]
			rawTickerData["Chart"] = strings.Split(strings.Split(body, "<img src='")[1], "'>")[0]
			body = strings.Split(strings.Split(body, "<img src='")[1], "'>")[1]
			rawTickerData["Company"] = strings.Split(strings.Split(body, "<b>")[1], "</b>")[0]
			body = strings.Split(strings.Split(strings.Split(body, "<b>")[1], "</b>")[1], "<br>\u00a0")[1]
			rawTickerData["Industry"] = strings.Split(body, " | ")[0]
			rawTickerData["Country"] = strings.Split(body, " | ")[1]
			rawTickerData["Market Cap"] = strings.Split(body, " | ")[2]
			rawTickerData["Change"] = strings.Split(strings.Split(body, " | ")[3], "Change: ")[1]
		}
		rawTickerData["Ticker"] = strings.TrimSpace(tickerNode.Text())
		tickerDataSlice = append(tickerDataSlice, rawTickerData)
	})

	headers := []string{"Ticker", "Change", "Chart", "Company", "Industry", "Country", "Market Cap"}
	return generateRowsHelper(headers, tickerDataSlice)
}

// bulkTickersDefaultViewScraper scrapes a 510 Screen document for the screen's ticker results (max 500 tickers)
func bulkTickersDefaultViewScraper(html interface{}) (rows [][]string, err error) {
	doc, err := generateDocument(html)
	if err != nil {
		return nil, err
	}

	var tickers []map[string]interface{}
	var tickerDataSlice [][]map[string]interface{}
	var rawTickerData map[string]interface{}
	rootNode := doc.Find("#screener-content").Find("tbody").Eq(3).Find("tr")

	rootNode.Children().Each(func(i int, columnNode *goquery.Selection) {
		var columnDataSlice []map[string]interface{}
		columnNode.Children().Find("table > tbody").Children().Each(func(j int, rowNode *goquery.Selection) {
			rawTickerData = make(map[string]interface{})
			rowNode.Children().Each(func(k int, itemNode *goquery.Selection) {
				if k == 0 {
					if titleAttr := itemNode.AttrOr("title", ""); titleAttr != "" {
						body := strings.Split(strings.Split(titleAttr, "body=[")[2], "]")[0]
						rawTickerData["Chart"] = strings.Split(strings.Split(body, "<img src='")[1], "'>")[0]
						body = strings.Split(strings.Split(body, "<img src='")[1], "'>")[1]
						rawTickerData["Company"] = strings.Split(strings.Split(body, "<b>")[1], "</b>")[0]
						body = strings.Split(strings.Split(strings.Split(body, "<b>")[1], "</b>")[1], "<br>\u00a0")[1]
						rawTickerData["Industry"] = strings.Split(body, " | ")[0]
						rawTickerData["Country"] = strings.Split(body, " | ")[1]
						rawTickerData["Market Cap"] = strings.Split(body, " | ")[2]
					}
					rawTickerData["Ticker"] = itemNode.Text()
				} else if k == 1 {
					rawTickerData["Change"] = strings.TrimSpace(itemNode.Text())
				}
			})
			columnDataSlice = append(columnDataSlice, rawTickerData)
		})
		tickerDataSlice = append(tickerDataSlice, columnDataSlice)
	})

	columnCount := len(tickerDataSlice)
	rowCount := len(tickerDataSlice[0])
	for j := 0; j < rowCount; j++ {
		for i := 0; i < columnCount; i++ {
			if j <= len(tickerDataSlice[i])-1 {
				tickers = append(tickers, tickerDataSlice[i][j])
			}
		}
	}

	headers := []string{"Ticker", "Change", "Chart", "Company", "Industry", "Country", "Market Cap"}
	return generateRowsHelper(headers, tickers)
}

// bulkTickersFullViewScraper scrapes a 520 Screen document for the screen's ticker results (max 500 tickers)
func bulkTickersFullViewScraper(html interface{}) (rows [][]string, err error) {
	doc, err := generateDocument(html)
	if err != nil {
		return nil, err
	}

	var tickers []map[string]interface{}
	var tickerDataSlice [][]map[string]interface{}
	var rawTickerData map[string]interface{}
	rootNode := doc.Find("#screener-content").Find("tbody").Eq(3).Find("tr")

	rootNode.Children().Each(func(i int, columnNode *goquery.Selection) {
		var columnDataSlice []map[string]interface{}
		columnNode.Children().Find("table > tbody").Children().Each(func(j int, rowNode *goquery.Selection) {
			rawTickerData = make(map[string]interface{})
			rowNode.Children().Each(func(k int, itemNode *goquery.Selection) {
				if k == 0 {
					if titleAttr := itemNode.AttrOr("title", ""); titleAttr != "" {
						body := strings.Split(strings.Split(titleAttr, "body=[")[2], "]")[0]
						rawTickerData["Chart"] = strings.Split(strings.Split(body, "<img src='")[1], "'>")[0]
						body = strings.Split(strings.Split(body, "<img src='")[1], "'>")[1]
						rawTickerData["Company"] = strings.Split(strings.Split(body, "<b>")[1], "</b>")[0]
						body = strings.Split(strings.Split(strings.Split(body, "<b>")[1], "</b>")[1], "<br>\u00a0")[1]
						rawTickerData["Industry"] = strings.Split(body, " | ")[0]
						rawTickerData["Country"] = strings.Split(body, " | ")[1]
						rawTickerData["Market Cap"] = strings.Split(body, " | ")[2]
					}
					rawTickerData["Ticker"] = itemNode.Text()
				} else if k == 1 {
					rawTickerData["Price"] = strings.TrimSpace(itemNode.Text())
				} else if k == 2 {
					rawTickerData["Change"] = itemNode.Children().Eq(0).Text()
				} else if k == 3 {
					rawTickerData["Volume"] = strings.TrimSpace(itemNode.Text())
				} else if k == 4 {
					rawTickerData["Relative Volume"] = itemNode.Find("small").Text()
				}
			})
			columnDataSlice = append(columnDataSlice, rawTickerData)
		})
		tickerDataSlice = append(tickerDataSlice, columnDataSlice)
	})

	columnCount := len(tickerDataSlice)
	rowCount := len(tickerDataSlice[0])
	for j := 0; j < rowCount; j++ {
		for i := 0; i < columnCount; i++ {
			if j <= len(tickerDataSlice[i])-1 {
				tickers = append(tickers, tickerDataSlice[i][j])
			}
		}
	}

	headers := []string{"Ticker", "Change", "Price", "Volume", "Relative Volume", "Chart", "Company", "Industry", "Country", "Market Cap"}
	return generateRowsHelper(headers, tickers)
}
*/

// defaultViewScraper scrapes an HTML document for the screen's ticker results (max 20 tickers)
func defaultViewScraper(html interface{}) (rows [][]string, err error) {
	doc, err := generateDocument(html)
	if err != nil {
		return nil, err
	}

	// Only collects data from the v=111 view (Note: maximum 20 stocks are returned)
	doc.Find("#screener-content").Find("tbody").Eq(3).Each(func(i int, childNode *goquery.Selection) {
		childNode.Children().Each(func(i int, childNode *goquery.Selection) {
			var row []string
			childNode.Children().Each(func(i int, childNode *goquery.Selection) {
				row = append(row, childNode.Text())
			})
			rows = append(rows, row)
		})
	})

	return rows, nil
}

// GetStockDataframe consumes an html instance and returns a dataframe of stocks returned
func GetStockDataframe(html interface{}, viewType ViewType) (*dataframe.DataFrame, error) {
	var results [][]string
	var err error

	/*
		switch viewType {
		default:
			results, err = defaultViewScraper(html)
		case OverviewView:
			results, err = defaultViewScraper(html)
		case ValuationView:
			results, err = defaultViewScraper(html)
		case OwnershipView:
			results, err = defaultViewScraper(html)
		case PerformanceView:
			results, err = defaultViewScraper(html)
		case CustomView:
			results, err = defaultViewScraper(html)
		case FinancialView:
			results, err = defaultViewScraper(html)
		case TechnicalView:
			results, err = defaultViewScraper(html)
		case ChartsView:
			results, err = defaultViewScraper(html)
		case BasicDefaultView:
			results, err = basicDefaultViewScraper(html)
		case BasicNewsView:
			results, err = basicNewsViewScraper(html)
		case BasicDescriptionView:
			results, err = basicDescriptionViewScraper(html)
		case BasicSnapshotView:
			results, err = basicSnapshotViewScraper(html)
		case BasicTAView:
			results, err = basicDefaultViewScraper(html)
		case TickersView:
			results, err = tickersViewScraper(html)
		case BulkTickersDefaultView:
			results, err = bulkTickersDefaultViewScraper(html)
		case BulkTickersFullView:
			results, err = bulkTickersFullViewScraper(html)
		}
	*/
	results, err = defaultViewScraper(html)
	if err != nil {
		return nil, err
	}
	df := dataframe.LoadRecords(results)
	return &df, nil
}

// GetOrderedTickerList is a helper function for viewing tickers from a screen dataframe
func GetOrderedTickerList(df *dataframe.DataFrame) []string {
	var tickers []string
	column := df.Col("Ticker")
	colLen := column.Len()

	for i := 0; i < colLen; i++ {
		tickers = append(tickers, column.Val(i).(string))
	}

	return tickers
}
