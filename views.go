// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func (v *ViewType) getURLView() string {
	return fmt.Sprintf("v=%v", v.ViewID)
}

// Scrape scrapes the DefaultView (100 series) html document for the screen's ticker results
func (v *ViewType) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

// Scrape scrapes the BasicView (310) html document for the screen's ticker results
func (b *BasicView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

	return generateRows(headers, tickerDataSlice)
}

// Scrape scrapes the NewsView (320) html document for the screen's ticker results
func (n *NewsView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

	return generateRows(headers, tickerDataSlice)
}

// Scrape scrapes the DescriptionView (330) html document for the screen's ticker results
func (d *DescriptionView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

	return generateRows(headers, tickerDataSlice)
}

// Scrape scrapes the SnapshotView (340) html document for the screen's ticker results
func (s *SnapshotView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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
				switch j {
				case 0:
					headers, rawTickerData = basicNewsViewHelper(childNode, headers, rawTickerData)
				case 1:
					headers, rawTickerData = basicDescriptionViewHelper(childNode, headers, rawTickerData)
				case 2:
					headers, rawTickerData = basicInsiderTradingViewHelper(childNode, headers, rawTickerData)
				}
			}
			tickerDataSlice = append(tickerDataSlice, rawTickerData)
		}
	})

	return generateRows(headers, tickerDataSlice)
}

// Scrape scrapes the TAView (350) html document for the screen's ticker results
func (t *TAView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

	return generateRows(headers, tickerDataSlice)
}

// Scrape scrapes the TickersView (410) html document for the screen's ticker results (max 1000 tickers)
func (t *TickersView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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
	return generateRows(headers, tickerDataSlice)
}

// Scrape scrapes the BulkView (510) html document for the screen's ticker results (max 500 tickers)
func (b *BulkView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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
	return generateRows(headers, tickers)
}

// Scrape scrapes the BulkFullView (520) html document for the screen's ticker results (max 500 tickers)
func (b *BulkFullView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
	var tickers []map[string]interface{}
	var tickerDataSlice [][]map[string]interface{}
	var rawTickerData map[string]interface{}
	rootNode := doc.Find("#screener-content").Find("tbody").Eq(3).Find("tr")

	rootNode.Children().Each(func(i int, columnNode *goquery.Selection) {
		var columnDataSlice []map[string]interface{}
		columnNode.Children().Find("table > tbody").Children().Each(func(j int, rowNode *goquery.Selection) {
			rawTickerData = make(map[string]interface{})
			rowNode.Children().Each(func(k int, itemNode *goquery.Selection) {
				switch k {
				case 0:
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
				case 1:
					rawTickerData["Price"] = strings.TrimSpace(itemNode.Text())
				case 2:
					rawTickerData["Change"] = itemNode.Children().Eq(0).Text()
				case 3:
					rawTickerData["Volume"] = strings.TrimSpace(itemNode.Text())
				case 4:
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
	return generateRows(headers, tickers)
}
