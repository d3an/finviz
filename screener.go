// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"fmt"
	"sort"
	"strconv"

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
	View          string
	CustomColumns []string
}

func getCustomColumns(columns []string, view ViewInterface) (string, ViewInterface, error) {
	columnsLen := len(columns)
	if columnsLen == 0 {
		return "", view, nil
	}

	var orderedColumns []string
	for i := 0; i < columnsLen; i++ {
		if _, err := strconv.Atoi(columns[i]); err == nil {
			orderedColumns = append(orderedColumns, columns[i])
		} else if val, ok := CustomColumnLookup[columns[i]]; ok {
			orderedColumns = append(orderedColumns, val)
		} else {
			return "", view, fmt.Errorf("%v is not a valid custom column", columns[i])
		}
	}

	view, _ = GetViewFactory("custom")
	return fmt.Sprintf("&c=%v", strings.Join(orderedColumns, ",")), view, nil
}

func getFilterList(filters []FilterInterface) (string, error) {
	filterSize := len(filters)
	if filterSize == 0 {
		return "", nil
	}

	var validFilters []FilterInterface
	var filterKeys []string

	for i := 0; i < filterSize; i++ {
		if filters[i].GetValue() == "" {
			return "", NoValuesError(fmt.Sprintf("%v filter was initialized without a value.", filters[i].GetName()))
		}
		if filterArrayContains(validFilters, filters[i]) {
			return "", DuplicateFilterError(fmt.Sprintf("%v filter was declared more than once.", filters[i].GetName()))
		}
		validFilters = append(validFilters, filters[i])
		filterKeys = append(filterKeys, filters[i].GetURLKey())
	}

	filterList := strings.Join(filterKeys, ",")
	return fmt.Sprintf("&f=%v", filterList), nil
}

func getSignal(signal SignalType) string {
	if signal == "" {
		return ""
	}
	return fmt.Sprintf("&s=%v", signal)
}

func getSortOrder(generalOrder GeneralOrderType, signal SignalType, specificOrder SpecificOrderType) string {
	// To sort by Signal, the Signal field must be non-empty
	if specificOrder == Signal && signal == "" && generalOrder == "" {
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

	sort.Strings(tickers)
	tickerList := strings.Join(tickers, ",")
	return fmt.Sprintf("&t=%v", tickerList)
}

// GenerateURL consumes valid inputs to the screen and generates a corresponding valid URL
func GenerateURL(input *ScreenInput, view ViewInterface) (string, error) {
	signal := getSignal(input.Signal)
	filterList, err := getFilterList(input.Filters)
	if err != nil {
		return "", err
	}

	sortOrder := getSortOrder(input.GeneralOrder, input.Signal, input.SpecificOrder)
	tickerList := getTickerList(input.Tickers)
	customColumns, view, err := getCustomColumns(input.CustomColumns, view)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v?%v%v%v%v%v%v", APIURL, view.getURLView(), signal, filterList, tickerList, sortOrder, customColumns), nil
}

// RunScreen consumes a client and screen input to produce a dataframe of results
func RunScreen(c *http.Client, input *ScreenInput) (*dataframe.DataFrame, error) {
	view, err := GetViewFactory(input.View)
	if err != nil {
		return nil, err
	}

	url, err := GenerateURL(input, view)
	if err != nil {
		return nil, err
	}
	fmt.Println(url)

	html, err := MakeGetRequest(c, url)
	if err != nil {
		return nil, err
	}

	df, err := GetStockDataframe(html, view)
	if err != nil {
		return nil, err
	}

	return df, nil
}
