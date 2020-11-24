// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package quote

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/d3an/finviz"
	"strings"
)

// QuoteView represents the default view for the Quote app
type QuoteView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *QuoteView) GenerateURL(viewArgs *map[string]interface{}) (string, error) {
	n := finviz.View{}
	url, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	// Single Ticker
	ticker, err := getTickerString(viewArgs)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v/quote.ashx?t=%v&ty=c&p=d&b=1", url, strings.ToUpper(ticker)), nil
}

// Scrape scrapes the ticker quote view (quote.ashx) html document for the screen's ticker results
func (v *QuoteView) Scrape(doc *goquery.Document) (rows [][]string, err error) {
	basicData := doc.Find("body > table").Eq(2).Find("tbody").Eq(0)

	titleData := basicData.Children().Eq(5).Find("tbody").Eq(0).Children().Eq(1).Find("tbody").Eq(0)
	tableData := basicData.Children().Eq(6).Find("tbody").Eq(0)

	extraData := doc.Find("body > table").Eq(3).Find("tbody").Eq(0).ChildrenFiltered("tr")

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
				} else {
					rawTickerData[key] = cellNode.Find("b").Text()
				}
			}
		})
	})

	// Analyst Recommendations
	if len(extraData.Eq(3).Find("#news-table").Nodes) == 0 {
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
	insiderData := extraData.Eq(len(extraData.Nodes) - 4)
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
	}

	var headers []string
	for key := range rawTickerData {
		headers = append(headers, key)
	}

	return finviz.GenerateRows(headers, []map[string]interface{}{rawTickerData})
}

// MapScrape scrapes FinViz views to a KVP map
func (v *QuoteView) MapScrape(doc *goquery.Document) (*map[string]interface{}, error) {
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
					/*} else if key == "EPS next Y" {
					if _, exists := rawTickerData["EPS next Y"]; exists {
						rawTickerData["EPS growth next Y"] = cellNode.Find("b").Text()
					} else {
						rawTickerData[key] = cellNode.Find("b").Text()
					}*/
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
