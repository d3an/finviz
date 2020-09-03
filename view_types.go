// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import "github.com/PuerkitoBio/goquery"

// ViewInterface introduces methods to interact with FinViz views
type ViewInterface interface {
	Scrape(document *goquery.Document) ([][]string, error)
	getURLComponent() string
}

// ChartViewInterface introduces methods to interact with FinViz views that contain charts
type ChartViewInterface interface {
	ViewInterface
	GetChartStyle() ChartStyle
	GetTimeFrame() TimeFrame
	SetChartStyle(string) error
	SetTimeFrame(string) error
}

// ChartStyle defines the types of charts available in the FinViz screener
type ChartStyle = string

// ChartStyle constant definitions
const (
	Candle    ChartStyle = "candle"
	Line      ChartStyle = "line"
	Technical ChartStyle = "technical"
)

// TimeFrame defines the period length of charts available in the FinViz screener
type TimeFrame = string

// TimeFrame constant definitions
const (
	Min1    TimeFrame = "1min"
	Min5    TimeFrame = "5min"
	Min15   TimeFrame = "15min"
	Min30   TimeFrame = "30min"
	Daily   TimeFrame = "daily"
	Weekly  TimeFrame = "weekly"
	Monthly TimeFrame = "monthly"
)

// ViewType handles the ID of each FinViz view
type ViewType struct {
	ViewID string
}

// ChartViewType handles the ID, ChartStyle, and the TimeFrame of each FinViz chart view.
type ChartViewType struct {
	ViewID     string
	chartStyle ChartStyle
	timeFrame  TimeFrame
}

// OverviewView (110)
type OverviewView struct {
	ViewType
}

// ValuationView (120)
type ValuationView struct {
	ViewType
}

// OwnershipView (130)
type OwnershipView struct {
	ViewType
}

// PerformanceView (140)
type PerformanceView struct {
	ViewType
}

// CustomView (150)
type CustomView struct {
	ViewType
}

// FinancialView (160)
type FinancialView struct {
	ViewType
}

// TechnicalView (170)
type TechnicalView struct {
	ViewType
}

// ChartsView (210)
type ChartsView struct {
	ChartViewType
}

// BasicView (310)
type BasicView struct {
	ChartViewType
}

// NewsView (320)
type NewsView struct {
	ChartViewType
}

// DescriptionView (330)
type DescriptionView struct {
	ChartViewType
}

// SnapshotView (340)
type SnapshotView struct {
	ChartViewType
}

// TAView (350)
type TAView struct {
	ChartViewType
}

// TickersView (410)
type TickersView struct {
	ViewType
}

// BulkView (510)
type BulkView struct {
	ViewType
}

// BulkFullView (520)
type BulkFullView struct {
	ViewType
}

// Handle charts customizations (chart type, tick frequency)

// CustomColumnLookup provides an interface for selecting different columns for the CustomView
var CustomColumnLookup = map[string]string{
	"no":                                "0",
	"ticker":                            "1",
	"company":                           "2",
	"sector":                            "3",
	"industry":                          "4",
	"country":                           "5",
	"market cap":                        "6",
	"p/e":                               "7",
	"pe":                                "7",
	"forward p/e":                       "8",
	"forward pe":                        "8",
	"peg":                               "9",
	"p/s":                               "10",
	"ps":                                "10",
	"p/b":                               "11",
	"pb":                                "11",
	"p/cash":                            "12",
	"p/c":                               "12",
	"pc":                                "12",
	"p/free cash flow":                  "13",
	"p/fcf":                             "13",
	"pfcf":                              "13",
	"dividend yield":                    "14",
	"payout ratio":                      "15",
	"eps":                               "16",
	"eps growth this year":              "17",
	"eps growth next year":              "18",
	"eps growth past 5 years":           "19",
	"eps growth next 5 years":           "20",
	"sales growth past 5 years":         "21",
	"eps growth qtr over qtr":           "22",
	"eps growth quarter over quarter":   "22",
	"sales growth qtr over qtr":         "23",
	"sales growth quarter over quarter": "23",
	"shares outstanding":                "24",
	"so":                                "24",
	"shares float":                      "25",
	"float":                             "25",
	"insider ownership":                 "26",
	"insider transactions":              "27",
	"institutional ownership":           "28",
	"institutional transactions":        "29",
	"float short":                       "30",
	"short selling":                     "30",
	"short ratio":                       "31",
	"return on assets":                  "32",
	"roa":                               "32",
	"return on equity":                  "33",
	"roe":                               "33",
	"return on investment":              "34",
	"roi":                               "34",
	"current ratio":                     "35",
	"quick ratio":                       "36",
	"long term debt/equity":             "37",
	"long-term debt/equity":             "37",
	"lt d/e":                            "37",
	"lt de":                             "37",
	"total debt/equity":                 "38",
	"debt/equity":                       "38",
	"d/e":                               "38",
	"de":                                "38",
	"gross margin":                      "39",
	"gm":                                "39",
	"operating margin":                  "40",
	"om":                                "40",
	"net profit margin":                 "41",
	"npm":                               "41",
	"performance (week)":                "42",
	"performance (month)":               "43",
	"performance (quarter)":             "44",
	"performance (half year)":           "45",
	"performance (year)":                "46",
	"performance (yeartodate)":          "47",
	"performance (ytd)":                 "47",
	"beta":                              "48",
	"average true range":                "49",
	"atr":                               "49",
	"volatility (week)":                 "50",
	"volatility w":                      "50",
	"volatility (month)":                "51",
	"volatility m":                      "51",
	"20-day simple moving average":      "52",
	"sma20":                             "52",
	"50-day simple moving average":      "53",
	"sma50":                             "53",
	"200-day simple moving average":     "54",
	"sma200":                            "54",
	"50-day high":                       "55",
	"50-day low":                        "56",
	"52-week high":                      "57",
	"52-week low":                       "58",
	"rsi":                               "59",
	"relative strength index":           "59",
	"change from open":                  "60",
	"gap":                               "61",
	"analyst recommendation":            "62",
	"analyst recom":                     "62",
	"recommendation":                    "62",
	"recom":                             "62",
	"average volume":                    "63",
	"avgvol":                            "63",
	"relative volume":                   "64",
	"relvol":                            "64",
	"price":                             "65",
	"change":                            "66",
	"volume":                            "67",
	"earnings date":                     "68",
	"target price":                      "69",
	"ipo date":                          "70",
}
