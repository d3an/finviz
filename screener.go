// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"fmt"

	"github.com/go-gota/gota/dataframe"

	"net/http"
	"strings"
)

// APIURL is the base URL for the screener. Elite not supported yet.
const APIURL = "https://finviz.com/screener.ashx"

// ScreenInput represents the data passed to the screen
type ScreenInput struct {
	Signal        SignalType
	GeneralOrder  GeneralOrderType
	SpecificOrder SpecificOrderType
	Tickers       []string
	Filters       []FilterInterface
	View          ViewType
	CustomColumns []string
}

func getScreenerView(viewType ViewType) string {
	if viewType == "" {
		viewType = "1"
	}

	return fmt.Sprintf("v=%v", viewType)
}

func getFilterList(filters []FilterInterface) string {
	filterSize := len(filters)
	if filterSize == 0 {
		return ""
	}

	var filterKeys []string
	for i := 0; i < filterSize; i++ {
		filterKeys = append(filterKeys, filters[i].GetURLKey())
	}

	filterList := strings.Join(filterKeys, ",")
	return fmt.Sprintf("&f=%v", filterList)
}

func getSignal(signal SignalType) string {
	if signal == "" {
		return ""
	}
	return fmt.Sprintf("&s=%v", signal)
}

func getSortOrder(generalOrder GeneralOrderType, signal SignalType, specificOrder SpecificOrderType) string {
	if specificOrder == Signal && signal == "" && generalOrder == Ascending {
		return ""
	} else if specificOrder == Signal && signal == "" {
		return fmt.Sprintf("&o=%v", generalOrder)
	}

	if specificOrder == "" && generalOrder == "" {
		return ""
	}

	return fmt.Sprintf("&o=%v%v", generalOrder, specificOrder)
}

func getTickerList(tickers []string) string {
	tickersSize := len(tickers)
	if tickersSize == 0 {
		return ""
	}

	for i := 0; i < tickersSize; i++ {
		tickers[i] = strings.ToUpper(tickers[i])
	}

	tickerList := strings.Join(tickers, ",")
	return fmt.Sprintf("&t=%v", tickerList)
}

// GenerateURL consumes valid inputs to the screen and generates a corresponding valid URL
func GenerateURL(input ScreenInput) string {
	screenerView := getScreenerView(input.View)
	signal := getSignal(input.Signal)
	filterList := getFilterList(input.Filters)
	sortOrder := getSortOrder(input.GeneralOrder, input.Signal, input.SpecificOrder)
	tickerList := getTickerList(input.Tickers)

	return fmt.Sprintf("%v?%v%v%v%v%v", APIURL, screenerView, signal, filterList, tickerList, sortOrder)
}

// RunScreen consumes a client and screen input to produce a dataframe of results
func RunScreen(c *http.Client, input ScreenInput) (*dataframe.DataFrame, error) {
	url := GenerateURL(input)

	html, err := MakeGetRequest(c, url)
	if err != nil {
		return nil, err
	}

	df, err := GetStockDataframe(html, input.View)
	if err != nil {
		return nil, err
	}

	return df, nil
}
