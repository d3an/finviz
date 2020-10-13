// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package news

import (
	"github.com/d3an/finviz"
	"github.com/go-gota/gota/dataframe"
	"net/http"
)

// GetNewsData returns a DataFrame containing recent news data
func GetNewsData(c *http.Client, v finviz.ViewInterface) (*dataframe.DataFrame, error) {
	url, err := v.GenerateURL(nil)
	if err != nil {
		return nil, err
	}

	html, err := finviz.MakeGetRequest(c, url)
	if err != nil {
		return nil, err
	}

	doc, err := finviz.GenerateDocument(html)
	if err != nil {
		return nil, err
	}

	results, err := v.Scrape(doc)
	if err != nil {
		return nil, err
	}

	df := dataframe.LoadRecords(results)
	return &df, nil
}

// NewsSourceAttributeLookup provides an interface for identifying news sources based on their CSS attributes
var NewsSourceAttributeLookup = map[string]string{
	"is-1":   "MarketWatch",
	"is-2":   "WSJ",
	"is-3":   "Reuters",
	"is-4":   "Yahoo Finance",
	"is-5":   "CNN",
	"is-6":   "The New York Times",
	"is-7":   "Bloomberg",
	"is-9":   "BBC",
	"is-10":  "CNBC",
	"is-11":  "Fox Business",
	"is-102": "Mish's Global Economic Trend Analysis",
	"is-105": "Trader Feed",
	"is-113": "Howard Lindzon",
	"is-114": "Seeking Alpha",
	"is-121": "The Disciplined Investor",
	"is-123": "Fallond Stock Picks",
	"is-132": "Zero Hedge",
	"is-133": "market folly",
	"is-136": "Daily Reckoning",
	"is-137": "Vantage Point Trading",
	"is-141": "Abnormal Returns",
	"is-142": "Calculated Risk",
}
