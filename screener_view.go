// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

// ViewType represents the general view in which results are displayed
type ViewType = string

// ColumnType represents the various column sets available for screens
type ColumnType = string

// TabType represents which tab (filters, settings, or none) is visible on screen
type TabType = string

// View Types
const (
	// Basic Scraper
	OverviewView    ViewType = "110"
	ValuationView   ViewType = "120"
	OwnershipView   ViewType = "130"
	PerformanceView ViewType = "140"
	CustomView      ViewType = "150"
	FinancialView   ViewType = "160"
	TechnicalView   ViewType = "170"

	ChartsView             ViewType = "210"
	BasicDefaultView       ViewType = "310"
	BasicNewsView          ViewType = "320"
	BasicDescriptionView   ViewType = "330"
	BasicSnapshotView      ViewType = "340"
	BasicTAView            ViewType = "350"
	TickersView            ViewType = "410"
	BulkTickersDefaultView ViewType = "510"
	BulkTickersFullView    ViewType = "520"
)

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
