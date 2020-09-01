// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"github.com/dnaeon/go-vcr/recorder"
	"testing"
)

// TestChartViewInterface_SetTimeFrame tests the correctness of the ChartViewInterface.SetTimeFrame method
func TestChartViewInterface_SetTimeFrame(t *testing.T) {
	failingTestInputs := []struct {
		curChartStyle string
		newTimeFrame  string
	}{
		{
			"technical",
			"weekly",
		},
		{
			"technical",
			"",
		},
		{
			"technical",
			"monthly",
		},
		{
			"technical",
			"5min",
		},
		{
			"line",
			"15min",
		},
		{
			"line",
			"30min",
		},
		{
			"line",
			"",
		},
		{
			"candle",
			"1min",
		},
		{
			"candle",
			"",
		},
		// When Elite is supported, these should pass:
		{
			"line",
			"1min",
		},
		{
			"line",
			"5min",
		},
		{
			"candle",
			"5min",
		},
		{
			"candle",
			"15min",
		},
		{
			"candle",
			"30min",
		},
	}

	passingTestInputs := []struct {
		curChartStyle string
		newTimeFrame  string
	}{
		{
			"technical",
			"daily",
		},
		{
			"line",
			"weekly",
		},
		{
			"candle",
			"monthly",
		},
	}

	for _, testInput := range failingTestInputs {
		view := ChartsView{ChartViewType{"charts", testInput.curChartStyle, "daily"}}
		err := view.SetTimeFrame(testInput.newTimeFrame)
		if err == nil {
			t.Fail()
			t.Logf("Expected SetChartStyle to fail when new TimeFrame is \"%v\" and existing ChartStyle is \"%v\"", testInput.newTimeFrame, testInput.curChartStyle)
		}
	}

	for _, testInput := range passingTestInputs {
		view := ChartsView{ChartViewType{"charts", testInput.curChartStyle, "daily"}}
		err := view.SetTimeFrame(testInput.newTimeFrame)
		if err != nil {
			t.Fail()
			t.Logf("Expected SetTimeFrame to pass when new TimeFrame is \"%v\" and existing ChartStyle is \"%v\"", testInput.newTimeFrame, testInput.curChartStyle)
		}
	}
}

// TestChartViewInterface_SetChartStyle tests the correctness of the ChartViewInterface.SetChartStyle method
func TestChartViewInterface_SetChartStyle(t *testing.T) {
	failingTestInputs := []struct {
		newChartStyle string
		curTimeFrame  string
	}{
		{
			"invalid",
			"daily",
		},
		{
			"",
			"daily",
		},
		{
			"technical",
			"weekly",
		},
		{
			"line",
			"15min",
		},
		{
			"line",
			"30min",
		},
		{
			"candle",
			"1min",
		},
	}

	passingTestInputs := []struct {
		newChartStyle string
		curTimeFrame  string
	}{
		{
			"technical",
			"daily",
		},
		{
			"line",
			"weekly",
		},
		{
			"candle",
			"monthly",
		},
		{
			"line",
			"1min",
		},
		{
			"line",
			"5min",
		},
		{
			"candle",
			"5min",
		},
		{
			"candle",
			"15min",
		},
		{
			"candle",
			"30min",
		},
	}

	for _, testInput := range failingTestInputs {
		view := ChartsView{ChartViewType{"charts", "technical", testInput.curTimeFrame}}
		err := view.SetChartStyle(testInput.newChartStyle)
		if err == nil {
			t.Fail()
			t.Logf("Expected SetTimeFrame to fail when new ChartStyle is \"%v\" and existing TimeFrame is \"%v\"", testInput.newChartStyle, testInput.curTimeFrame)
		}
	}

	for _, testInput := range passingTestInputs {
		view := ChartsView{ChartViewType{"charts", "technical", testInput.curTimeFrame}}
		err := view.SetChartStyle(testInput.newChartStyle)
		if err != nil {
			t.Fail()
			t.Logf("Expected SetTimeFrame to pass when new ChartStyle is \"%v\" and existing TimeFrame is \"%v\"", testInput.newChartStyle, testInput.curTimeFrame)
		}
	}
}

// TestViewInterface_Scrape tests the scrape functions for all supported/available views
// If FinViz changes their view templates, fixtures must be regenerated
func TestViewInterface_Scrape(t *testing.T) {
	testInputs := []struct {
		cassetteName string
		screenInput  ScreenInput
		columnNames  []string
		nRows        int
		nCols        int
	}{
		{
			"fixtures/finviz_screener_view_110",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "overview",
			},
			[]string{"No.", "Ticker", "Company", "Sector", "Industry", "Country", "Market Cap", "P/E", "Price", "Change", "Volume"},
			20,
			11,
		},
		{
			"fixtures/finviz_screener_view_120",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "valuation",
			},
			[]string{"No.", "Ticker", "Market Cap", "P/E", "Fwd P/E", "PEG", "P/S", "P/B", "P/C", "P/FCF", "EPS this Y", "EPS next Y", "EPS past 5Y", "EPS next 5Y", "Sales past 5Y", "Price", "Change", "Volume"},
			20,
			18,
		},
		{
			"fixtures/finviz_screener_view_130",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "ownership",
			},
			[]string{"No.", "Ticker", "Market Cap", "Outstanding", "Float", "Insider Own", "Insider Trans", "Inst Own", "Inst Trans", "Float Short", "Short Ratio", "Avg Volume", "Price", "Change", "Volume"},
			20,
			15,
		},
		{
			"fixtures/finviz_screener_view_140",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "performance",
			},
			[]string{"No.", "Ticker", "Perf Week", "Perf Month", "Perf Quart", "Perf Half", "Perf Year", "Perf YTD", "Volatility W", "Volatility M", "Recom", "Avg Volume", "Rel Volume", "Price", "Change", "Volume"},
			20,
			16,
		},
		{
			"fixtures/finviz_screener_view_150",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "custom",
				CustomColumns: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70"},
			},
			[]string{"No.", "Ticker", "Company", "Sector", "Industry", "Country", "Market Cap", "P/E", "Fwd P/E", "PEG", "P/S", "P/B", "P/C", "P/FCF", "Dividend", "Payout Ratio", "EPS", "EPS this Y", "EPS next Y", "EPS past 5Y", "EPS next 5Y", "Sales past 5Y", "EPS Q/Q", "Sales Q/Q", "Outstanding", "Float", "Insider Own", "Insider Trans", "Inst Own", "Inst Trans", "Float Short", "Short Ratio", "ROA", "ROE", "ROI", "Curr R", "Quick R", "LTDebt/Eq", "Debt/Eq", "Gross M", "Oper M", "Profit M", "Perf Week", "Perf Month", "Perf Quart", "Perf Half", "Perf Year", "Perf YTD", "Beta", "ATR", "Volatility W", "Volatility M", "SMA20", "SMA50", "SMA200", "50D High", "50D Low", "52W High", "52W Low", "RSI", "from Open", "Gap", "Recom", "Avg Volume", "Rel Volume", "Price", "Change", "Volume", "Earnings", "Target Price", "IPO Date"},
			20,
			71,
		},
		{
			"fixtures/finviz_screener_view_160",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "financial",
			},
			[]string{"No.", "Ticker", "Market Cap", "Dividend", "ROA", "ROE", "ROI", "Curr R", "Quick R", "LTDebt/Eq", "Debt/Eq", "Gross M", "Oper M", "Profit M", "Earnings", "Price", "Change", "Volume"},
			20,
			18,
		},
		{
			"fixtures/finviz_screener_view_170",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "technical",
			},
			[]string{"No.", "Ticker", "Beta", "ATR", "SMA20", "SMA50", "SMA200", "52W High", "52W Low", "RSI", "Price", "Change", "from Open", "Gap", "Volume"},
			20,
			15,
		},
		{
			"fixtures/finviz_screener_view_210_technical",
			ScreenInput{
				Signal:           AllStocks,
				GeneralOrder:     Descending,
				SpecificOrder:    Change,
				View:             "charts",
				CustomChartStyle: "technical",
			},
			[]string{"Ticker", "Chart", "Company", "Industry", "Country", "Market Cap"},
			12,
			6,
		},
		{
			"fixtures/finviz_screener_view_210_line",
			ScreenInput{
				Signal:           AllStocks,
				GeneralOrder:     Descending,
				SpecificOrder:    Change,
				View:             "charts",
				CustomChartStyle: "line",
			},
			[]string{"Ticker", "Chart", "Company", "Industry", "Country", "Market Cap"},
			24,
			6,
		},
		{
			"fixtures/finviz_screener_view_210_candle",
			ScreenInput{
				Signal:           AllStocks,
				GeneralOrder:     Descending,
				SpecificOrder:    Change,
				View:             "charts",
				CustomChartStyle: "candle",
			},
			[]string{"Ticker", "Chart", "Company", "Industry", "Country", "Market Cap"},
			24,
			6,
		},
		{
			"fixtures/finviz_screener_view_310",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "basic",
			},
			[]string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range"},
			10,
			29,
		},
		{
			"fixtures/finviz_screener_view_320",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "news",
			},
			[]string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "News"},
			10,
			30,
		},
		{
			"fixtures/finviz_screener_view_330",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "description",
			},
			[]string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "Description"},
			10,
			30,
		},
		{
			"fixtures/finviz_screener_view_340",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "snapshot",
			},
			[]string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "News", "Description", "Insider Trading"},
			10,
			32,
		},
		{
			"fixtures/finviz_screener_view_350",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "ta",
			},
			[]string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "Perf Week", "Beta", "Perf Month", "ATR", "Perf Quarter", "Volatility W", "Perf Half Y", "Volatility M", "Perf Year", "SMA20", "Perf YTD", "SMA50", "RSI (14)", "SMA200", "Change Open", "52W High", "Gap", "52W Low", "Rel Volume", "Short Float", "Avg Volume", "Candlestick", "52W Range"},
			10,
			29,
		},
		{
			"fixtures/finviz_screener_view_410",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "tickers",
			},
			[]string{"Ticker", "Change", "Chart", "Company", "Industry", "Country", "Market Cap"},
			1000,
			7,
		},
		{
			"fixtures/finviz_screener_view_510",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "bulk",
			},
			[]string{"Ticker", "Change", "Chart", "Company", "Industry", "Country", "Market Cap"},
			500,
			7,
		},
		{
			"fixtures/finviz_screener_view_520",
			ScreenInput{
				Signal:        AllStocks,
				GeneralOrder:  Descending,
				SpecificOrder: Change,
				View:          "bulkfull",
			},
			[]string{"Ticker", "Change", "Price", "Volume", "Relative Volume", "Chart", "Company", "Industry", "Country", "Market Cap"},
			500,
			10,
		},
	}

	for _, testInput := range testInputs {
		r, err := recorder.New(testInput.cassetteName)
		if err != nil {
			t.Error(err)
		}

		client := newTestingClient(r)

		df, err := RunScreen(client, &testInput.screenInput)
		if err != nil {
			t.Error(err)
		}

		// Check column count
		if df.Ncol() != testInput.nCols {
			t.Fail()
			t.Logf("Expected %v columns, received %v columns", testInput.nCols, df.Nrow())
		}

		// Check column names
		receivedColumns := df.Names()
		for i := 0; i < testInput.nCols; i++ {
			if receivedColumns[i] != testInput.columnNames[i] {
				t.Fail()
				t.Logf("Expected column named \"%v\", received column named \"%v\"", testInput.columnNames[i], receivedColumns[i])
			}
		}

		// Check row count
		if df.Nrow() != testInput.nRows {
			t.Fail()
			t.Logf("Expected %v rows, received %v rows", testInput.nRows, df.Nrow())
		}

		if err := r.Stop(); err != nil {
			t.Error(err)
		}
	}
}
