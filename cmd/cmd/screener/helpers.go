// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package screener

import (
	"fmt"
	"github.com/d3an/finviz"
	"github.com/d3an/finviz/screener"
	"os"
	"strings"
)

func er(msg interface{}) {
	fmt.Println("Error: ", msg)
	os.Exit(1)
}

func extractFilterInput(filterArg string) (filterName string, filterValues []string, err error) {
	firstSplit := strings.Split(filterArg, ":")
	if splitCount := len(firstSplit); splitCount != 2 {
		return "", nil, MalformedFilterError(fmt.Sprintf("Filter argument: \"%v\" is not well-formed. Proper form: \"Filter:Value1,Value2\" OR \"filter:value\"", filterArg))
	}
	return firstSplit[0], strings.Split(firstSplit[1], ","), nil
}

// ViewFactory consumes a view query string and returns the associated ViewInterface
func ViewFactory(viewQuery string) (finviz.ViewInterface, error) {
	switch strings.ToLower(viewQuery) {
	default:
		return nil, finviz.InvalidViewError(fmt.Sprintf("view \"%v\" is not supported", viewQuery))
	case "overview", "110":
		return &screener.OverviewScreenerView{}, nil
	case "valuation", "120":
		return &screener.ValuationScreenerView{}, nil
	case "ownership", "130":
		return &screener.OwnershipScreenerView{}, nil
	case "performance", "140":
		return &screener.PerformanceScreenerView{}, nil
	case "custom", "150":
		return &screener.CustomScreenerView{}, nil
	case "financial", "160":
		return &screener.FinancialScreenerView{}, nil
	case "technical", "170":
		return &screener.TechnicalScreenerView{}, nil
	case "charts", "210":
		return &screener.ChartsScreenerView{}, nil
	case "basic", "310":
		return &screener.BasicScreenerView{}, nil
	case "news", "320":
		return &screener.NewsScreenerView{}, nil
	case "description", "330":
		return &screener.DescriptionScreenerView{}, nil
	case "snapshot", "340":
		return &screener.SnapshotScreenerView{}, nil
	case "ta", "350":
		return &screener.TAScreenerView{}, nil
	case "tickers", "410":
		return &screener.TickersScreenerView{}, nil
	case "bulk", "510":
		return &screener.BulkScreenerView{}, nil
	case "bulkfull", "520":
		return &screener.BulkFullScreenerView{}, nil
	}
}
