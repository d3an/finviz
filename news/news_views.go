// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package news

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/d3an/finviz"
	"strings"
)

// NewsView is the root view for the News app
type NewsView struct{}

// TimeNewsView is the view for news sorted by time
type TimeNewsView struct{}

// SourceNewsView is the view for news sorted by source
type SourceNewsView struct{}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *NewsView) GenerateURL(_ *map[string]interface{}) (string, error) {
	n := finviz.View{}
	url, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v/news.ashx", url), nil
}

// Scrape throws an error if used
func (v *NewsView) Scrape(_ *goquery.Document) ([][]string, error) {
	n := finviz.View{}
	return n.Scrape(nil)
}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *TimeNewsView) GenerateURL(_ *map[string]interface{}) (string, error) {
	n := NewsView{}
	return n.GenerateURL(nil)
}

// Scrape parses news data sorted by time
func (v *TimeNewsView) Scrape(doc *goquery.Document) ([][]string, error) {
	var newsDataSlice []map[string]interface{}

	doc.Find("#news > div").Children().Eq(1).Find("tbody").Eq(0).Children().Eq(1).Children().Each(func(i int, newsColumn *goquery.Selection) {
		var newsType string

		switch i {
		case 0:
			newsType = "news"
		case 2:
			newsType = "blog"
		}

		if i != 1 {
			newsColumn.Find("tbody").Eq(0).Children().Each(func(j int, newsItem *goquery.Selection) {
				var rawNewsData = make(map[string]interface{})
				rawNewsData["Article Date"] = newsItem.Children().Eq(1).Text()
				rawNewsData["Article Title"] = newsItem.Children().Eq(2).Children().Eq(0).Text()
				rawNewsData["Article URL"] = newsItem.Children().Eq(2).Children().Eq(0).AttrOr("href", "")
				if classes := newsItem.Children().Eq(0).AttrOr("class", ""); classes != "" {
					if val, exists := NewsSourceAttributeLookup[strings.Split(classes, " ")[1]]; exists {
						rawNewsData["Source Name"] = val
					} else {
						rawNewsData["Source Name"] = "Unknown"
					}
				}
				rawNewsData["News Type"] = newsType
				newsDataSlice = append(newsDataSlice, rawNewsData)
			})
		}
	})

	headers := []string{"Article Date", "Article Title", "Article URL", "Source Name", "News Type"}
	return finviz.GenerateRows(headers, newsDataSlice)
}

// GenerateURL consumes valid inputs to the screen and generates a valid URL
func (v *SourceNewsView) GenerateURL(_ *map[string]interface{}) (string, error) {
	n := &NewsView{}
	url, err := n.GenerateURL(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?v=2", url), nil
}

// Scrape parses news data sorted by source
func (v *SourceNewsView) Scrape(doc *goquery.Document) ([][]string, error) {
	var newsDataSlice []map[string]interface{}

	for tableIndex := 2; tableIndex <= 4; tableIndex++ {
		var newsType string

		switch tableIndex {
		case 2:
			newsType = "news"
		case 4:
			newsType = "blog"
		}

		if tableIndex != 3 {
			doc.Find("#news > div").Children().Eq(tableIndex).Find("tr").Children().Each(func(i int, newsColumn *goquery.Selection) {
				if align := newsColumn.AttrOr("align", ""); align != "" {
					newsColumn.Children().Each(func(j int, newsSource *goquery.Selection) {
						var sourceName string
						var sourceURL string
						newsSource.Find("tbody").Eq(0).Children().Each(func(k int, item *goquery.Selection) {
							if k == 1 {
								sourceItem := item.Find("tr").Children().Eq(1).Find("a")
								sourceName = sourceItem.Text()
								sourceURL = sourceItem.AttrOr("href", "")
							} else if k > 1 {
								var rawNewsData = make(map[string]interface{})
								rawNewsData["Article Date"] = item.Children().Eq(0).Text()
								rawNewsData["Article Title"] = item.Children().Eq(1).Find("a").Text()
								rawNewsData["Article URL"] = item.Children().Eq(1).Children().Eq(0).AttrOr("href", "")
								rawNewsData["Source Name"] = sourceName
								rawNewsData["Source URL"] = sourceURL
								rawNewsData["News Type"] = newsType
								newsDataSlice = append(newsDataSlice, rawNewsData)
							}
						})
					})
				}
			})
		}
	}

	headers := []string{"Article Date", "Article Title", "Article URL", "Source Name", "Source URL", "News Type"}
	return finviz.GenerateRows(headers, newsDataSlice)
}
