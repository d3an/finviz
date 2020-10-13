// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package screener

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/d3an/finviz"
	"strings"
)

// ScreenerView represents the default view for the Screener app
type ScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *ScreenerView) GenerateURL(_ *map[string]interface{}) (string, error) {
	n := finviz.View{}
	url, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v/screener.ashx", url), nil
}

// Scrape scrapes the DefaultView (100 series) html document for the screen's ticker results
func (v *ScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

// OverviewScreenerView (110)
type OverviewScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *OverviewScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=110%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs)), nil
}

// Scrape scrapes the DefaultView (100 series) html document for the screen's ticker results
func (v *OverviewScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
	n := ScreenerView{}
	return n.Scrape(doc)
}

// ValuationScreenerView (120)
type ValuationScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *ValuationScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=120%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs)), nil
}

// Scrape scrapes the DefaultView (100 series) html document for the screen's ticker results
func (v *ValuationScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
	n := ScreenerView{}
	return n.Scrape(doc)
}

// OwnershipScreenerView (130)
type OwnershipScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *OwnershipScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=130%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs)), nil
}

// Scrape scrapes the DefaultView (100 series) html document for the screen's ticker results
func (v *OwnershipScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
	n := ScreenerView{}
	return n.Scrape(doc)
}

// PerformanceScreenerView (140)
type PerformanceScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *PerformanceScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=140%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs)), nil
}

// Scrape scrapes the DefaultView (100 series) html document for the screen's ticker results
func (v *PerformanceScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
	n := ScreenerView{}
	return n.Scrape(doc)
}

// CustomScreenerView (150)
type CustomScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *CustomScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	customColumns, err := getCustomColumnsValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=150%v%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs), customColumns), nil
}

// Scrape scrapes the DefaultView (100 series) html document for the screen's ticker results
func (v *CustomScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
	n := ScreenerView{}
	return n.Scrape(doc)
}

// FinancialScreenerView (160)
type FinancialScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *FinancialScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=160%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs)), nil
}

// Scrape scrapes the DefaultView (100 series) html document for the screen's ticker results
func (v *FinancialScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
	n := ScreenerView{}
	return n.Scrape(doc)
}

// TechnicalScreenerView (170)
type TechnicalScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *TechnicalScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=170%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs)), nil
}

// Scrape scrapes the DefaultView (100 series) html document for the screen's ticker results
func (v *TechnicalScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
	n := ScreenerView{}
	return n.Scrape(doc)
}

// ChartsScreenerView (210)
type ChartsScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *ChartsScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	chartStyle, err := getChartStylingValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=210%v%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs), chartStyle), nil
}

// Scrape scrapes the ChartsScreenerView (210) html document for the screen's ticker results
func (v *ChartsScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
	var tickerDataSlice []map[string]interface{}
	var headers []string

	doc.Find("#screener-content").Find("tbody").Children().Eq(4).Find("tbody").Find("td").Children().Each(func(i int, spanNode *goquery.Selection) {
		var rawTickerData = make(map[string]interface{})
		if titleAttr := spanNode.AttrOr("title", ""); titleAttr != "" {
			body := strings.Split(strings.Split(strings.Split(titleAttr, "body=[")[2], "]")[0], "<b>")[1]
			rawTickerData["Company"] = strings.Split(body, "</b>")[0]
			body = strings.Split(body, "<br>")[1]
			rawTickerData["Industry"] = strings.Split(body, " | ")[0]
			rawTickerData["Country"] = strings.Split(body, " | ")[1]
			rawTickerData["Market Cap"] = strings.Split(body, " | ")[2]
		}
		rawTickerData["Ticker"] = strings.Split(strings.Split(spanNode.Find("a").AttrOr("href", ""), "quote.ashx?t=")[1], "&")[0]
		rawTickerData["Chart"] = spanNode.Find("img").AttrOr("src", "")
		tickerDataSlice = append(tickerDataSlice, rawTickerData)
	})

	headers = []string{"Ticker", "Chart", "Company", "Industry", "Country", "Market Cap"}
	return generateRows(headers, tickerDataSlice)
}

// BasicScreenerView (310)
type BasicScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *BasicScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	chartStyle, err := getChartStylingValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=310%v%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs), chartStyle), nil
}

// Scrape scrapes the BasicScreenerView (310) html document for the screen's ticker results
func (v *BasicScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

// NewsScreenerView (320)
type NewsScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *NewsScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	chartStyle, err := getChartStylingValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=320%v%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs), chartStyle), nil
}

// Scrape scrapes the NewsScreenerView (320) html document for the screen's ticker results
func (v *NewsScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

// DescriptionScreenerView (330)
type DescriptionScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *DescriptionScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	chartStyle, err := getChartStylingValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=330%v%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs), chartStyle), nil
}

// Scrape scrapes the DescriptionScreenerView (330) html document for the screen's ticker results
func (v *DescriptionScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

// SnapshotScreenerView (340)
type SnapshotScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *SnapshotScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	chartStyle, err := getChartStylingValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=340%v%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs), chartStyle), nil
}

// Scrape scrapes the SnapshotScreenerView (340) html document for the screen's ticker results
func (v *SnapshotScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

// TAScreenerView (350)
type TAScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *TAScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	chartStyle, err := getChartStylingValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=350%v%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs), chartStyle), nil
}

// Scrape scrapes the TAScreenerView (350) html document for the screen's ticker results
func (v *TAScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

// TickersScreenerView (410)
type TickersScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *TickersScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=410%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs)), nil
}

// Scrape scrapes the TickersScreenerView (410) html document for the screen's ticker results (max 1000 tickers)
func (v *TickersScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

// BulkScreenerView (510)
type BulkScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *BulkScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=510%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs)), nil
}

// Scrape scrapes the BulkScreenerView (510) html document for the screen's ticker results (max 500 tickers)
func (v *BulkScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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

// BulkFullScreenerView (520)
type BulkFullScreenerView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *BulkFullScreenerView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := ScreenerView{}
	screenerURL, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	filters, err := getFiltersValue(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=520%v%v%v%v", screenerURL, getSignalValue(viewArgs),
		filters, getTickersValue(viewArgs), getOrderValue(viewArgs)), nil
}

// Scrape scrapes the BulkFullScreenerView (520) html document for the screen's ticker results (max 500 tickers)
func (v *BulkFullScreenerView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
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
