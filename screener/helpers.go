// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package screener

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/d3an/finviz/utils"
)

func basicDefaultViewHelper(rootNode *goquery.Selection, headers []string, rawTickerData map[string]interface{}) (extraHeaders []string, extraRawTickerData map[string]interface{}) {
	rootNode.Find(".snapshot-table > tbody").Children().Each(func(j int, childNode *goquery.Selection) {
		if !utils.Contains(headers, childNode.Children().First().Text()) {
			headers = append(headers, childNode.Children().First().Text())
		}
		rawTickerData[childNode.Children().First().Text()] = childNode.Children().Last().Find("a").Text()
	})

	if !utils.Contains(headers, "Chart") {
		headers = append(headers, "Chart")
	}
	rawTickerData["Chart"] = rootNode.Find("img").AttrOr("src", "")

	rootNode.Find(".snapshot-table2 > tbody").Children().Each(func(j int, childNode *goquery.Selection) {
		if !utils.Contains(headers, childNode.Children().Eq(0).Text()) && !utils.Contains(headers, childNode.Children().Eq(2).Text()) {
			headers = append(append(headers, childNode.Children().Eq(0).Text()), childNode.Children().Eq(2).Text())
		}
		rawTickerData[childNode.Children().Eq(0).Text()] = childNode.Children().Eq(1).Text()
		rawTickerData[childNode.Children().Eq(2).Text()] = childNode.Children().Eq(3).Text()
	})

	return headers, rawTickerData
}

func basicNewsViewHelper(rootNode *goquery.Selection, headers []string, rawTickerData map[string]interface{}) (extraHeaders []string, extraRawTickerData map[string]interface{}) {
	if !utils.Contains(headers, "News") {
		headers = append(headers, "News")
	}
	var news []map[string]string
	rootNode.Find(".body-table-news > tbody").Children().Each(func(j int, childNode *goquery.Selection) {
		news = append(news, map[string]string{
			"Datetime": childNode.Children().Eq(0).Text(),
			"Link":     childNode.Children().Eq(1).Find("a").AttrOr("href", ""),
			"Source":   childNode.Children().Eq(1).Find("span").Text(),
			"Title":    childNode.Children().Eq(1).Find("a").Text(),
		})
	})
	rawTickerData["News"] = news
	return headers, rawTickerData
}

func basicDescriptionViewHelper(rootNode *goquery.Selection, headers []string, rawTickerData map[string]interface{}) (extraHeaders []string, extraRawTickerData map[string]interface{}) {
	if !utils.Contains(headers, "Description") {
		headers = append(headers, "Description")
	}
	description := rootNode.Find(".body-table-profile").Text()
	description = strings.TrimPrefix(description, "\"")
	description = strings.TrimSuffix(description, "\"")
	rawTickerData["Description"] = description
	return headers, rawTickerData
}

func basicInsiderTradingViewHelper(rootNode *goquery.Selection, headers []string, rawTickerData map[string]interface{}) (extraHeaders []string, extraRawTickerData map[string]interface{}) {
	if !utils.Contains(headers, "Insider Trading") {
		headers = append(headers, "Insider Trading")
	}
	var insiderTrading []map[string]string
	if insiderNode := rootNode.Find(".body-table > tbody"); insiderNode != nil {
		insiderNode.Children().Each(func(k int, childNode *goquery.Selection) {
			if k > 0 {
				insiderTrading = append(insiderTrading, map[string]string{
					"Owner":             childNode.Children().Eq(0).Find("a").Text(),
					"Owner Link":        fmt.Sprintf("/%v", childNode.Children().Eq(0).Find("a").AttrOr("href", "")),
					"Relationship":      childNode.Children().Eq(1).Text(),
					"Date":              childNode.Children().Eq(2).Text(),
					"Transaction":       childNode.Children().Eq(3).Text(),
					"Cost":              childNode.Children().Eq(4).Text(),
					"# of Shares":       childNode.Children().Eq(5).Text(),
					"Value ($)":         childNode.Children().Eq(6).Text(),
					"# of Shares Total": childNode.Children().Eq(7).Text(),
					"SEC Form 4 Date":   childNode.Children().Eq(8).Find("a").Text(),
					"SEC Form 4 Link":   childNode.Children().Eq(8).Find("a").AttrOr("href", ""),
				})
			}
		})
		rawTickerData["Insider Trading"] = insiderTrading
	}
	return headers, rawTickerData
}
