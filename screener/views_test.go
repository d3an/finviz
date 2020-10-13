// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package screener

import (
	"fmt"
	"github.com/d3an/finviz"
	"github.com/dnaeon/go-vcr/recorder"
	"testing"
)

func TestScreenerView_GenerateURL(t *testing.T) {
	testInputs := []struct {
		viewArgs          map[string]interface{}
		viewInterfaceType finviz.ViewInterface
		expectedURL       string
	}{
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&OverviewScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=110&s=ta_unusualvolume&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&ValuationScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=120&s=ta_unusualvolume&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&OwnershipScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=130&s=ta_unusualvolume&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&PerformanceScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=140&s=ta_unusualvolume&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
				"custom_columns": []string{"0", "1", "3", "Company", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "shares Outstanding", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "volume", "68", "69", "IPO Date", "70"},
			},
			&CustomScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=150&s=ta_unusualvolume&f=exch_nyse&o=-volume&c=0,1,3,2,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&FinancialScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=160&s=ta_unusualvolume&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&TechnicalScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=170&s=ta_unusualvolume&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
				"chart_type":     Candle,
				"timeframe":      Weekly,
			},
			&ChartsScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=210&s=ta_unusualvolume&f=exch_nyse&o=-volume&ta=0&p=w", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
				"chart_type":     Line,
				"timeframe":      Monthly,
			},
			&ChartsScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=210&s=ta_unusualvolume&f=exch_nyse&o=-volume&ty=l&ta=0&p=m", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&BasicScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=310&s=ta_unusualvolume&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&NewsScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=320&s=ta_unusualvolume&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&DescriptionScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=330&s=ta_unusualvolume&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&SnapshotScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=340&s=ta_unusualvolume&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&TAScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=350&s=ta_unusualvolume&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": AllStocks,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&TickersScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=410&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": AllStocks,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&BulkScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=510&f=exch_nyse&o=-volume", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"signal": AllStocks,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&BulkFullScreenerView{},
			fmt.Sprintf("%v/screener.ashx?v=520&f=exch_nyse&o=-volume", finviz.APIURL),
		},
	}

	for _, testInput := range testInputs {
		if url, err := testInput.viewInterfaceType.GenerateURL(&testInput.viewArgs); err != nil {
			t.Errorf("GenerateURL failed. Error: %v", err)
		} else if url != testInput.expectedURL {
			t.Fail()
			t.Logf("URL Generation failed. Expected: \"%v\", Received: \"%v\"", testInput.expectedURL, url)
		}
	}
}

func TestScreenerView_GetData(t *testing.T) {
	testInputs := []struct {
		cassettePath                 string
		viewArgs                     map[string]interface{}
		viewInterfaceType            finviz.ViewInterface
		expectedDataFrameRowCount    int
		expectedDataFrameColumnCount int
		expectedDataFrameColumnNames []string
	}{
		{
			"cassettes/overview_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&OverviewScreenerView{},
			20,
			11,
			[]string{"No.", "Ticker", "Company", "Sector", "Industry", "Country", "Market Cap", "P/E", "Price", "Change", "Volume"},
		},
		{
			"cassettes/valuation_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&ValuationScreenerView{},
			20,
			18,
			[]string{"No.", "Ticker", "Market Cap", "P/E", "Fwd P/E", "PEG", "P/S", "P/B", "P/C", "P/FCF", "EPS this Y", "EPS next Y", "EPS past 5Y", "EPS next 5Y", "Sales past 5Y", "Price", "Change", "Volume"},
		},
		{
			"cassettes/ownership_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&OwnershipScreenerView{},
			20,
			15,
			[]string{"No.", "Ticker", "Market Cap", "Outstanding", "Float", "Insider Own", "Insider Trans", "Inst Own", "Inst Trans", "Float Short", "Short Ratio", "Avg Volume", "Price", "Change", "Volume"},
		},
		{
			"cassettes/performance_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&PerformanceScreenerView{},
			20,
			16,
			[]string{"No.", "Ticker", "Perf Week", "Perf Month", "Perf Quart", "Perf Half", "Perf Year", "Perf YTD", "Volatility W", "Volatility M", "Recom", "Avg Volume", "Rel Volume", "Price", "Change", "Volume"},
		},
		{
			"cassettes/custom_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
				"custom_columns": []string{"0", "1", "Sector", "2", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "shares Outstanding", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "volume", "68", "69", "IPO Date", "70"},
			},
			&CustomScreenerView{},
			20,
			71,
			[]string{"No.", "Ticker", "Sector", "Company", "Industry", "Country", "Market Cap", "P/E", "Fwd P/E", "PEG", "P/S", "P/B", "P/C", "P/FCF", "Dividend", "Payout Ratio", "EPS", "EPS this Y", "EPS next Y", "EPS past 5Y", "EPS next 5Y", "Sales past 5Y", "EPS Q/Q", "Sales Q/Q", "Outstanding", "Float", "Insider Own", "Insider Trans", "Inst Own", "Inst Trans", "Float Short", "Short Ratio", "ROA", "ROE", "ROI", "Curr R", "Quick R", "LTDebt/Eq", "Debt/Eq", "Gross M", "Oper M", "Profit M", "Perf Week", "Perf Month", "Perf Quart", "Perf Half", "Perf Year", "Perf YTD", "Beta", "ATR", "Volatility W", "Volatility M", "SMA20", "SMA50", "SMA200", "50D High", "50D Low", "52W High", "52W Low", "RSI", "from Open", "Gap", "Recom", "Avg Volume", "Rel Volume", "Price", "Change", "Volume", "Earnings", "Target Price", "IPO Date"},
		},
		{
			"cassettes/financial_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&FinancialScreenerView{},
			20,
			18,
			[]string{"No.", "Ticker", "Market Cap", "Dividend", "ROA", "ROE", "ROI", "Curr R", "Quick R", "LTDebt/Eq", "Debt/Eq", "Gross M", "Oper M", "Profit M", "Earnings", "Price", "Change", "Volume"},
		},
		{
			"cassettes/technical_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&TechnicalScreenerView{},
			20,
			15,
			[]string{"No.", "Ticker", "Beta", "ATR", "SMA20", "SMA50", "SMA200", "52W High", "52W Low", "RSI", "Price", "Change", "from Open", "Gap", "Volume"},
		},
		{
			"cassettes/charts_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&ChartsScreenerView{},
			12,
			6,
			[]string{"Ticker", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
		{
			"cassettes/charts_candle_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
				"chart_type":     Candle,
				"timeframe":      Weekly,
			},
			&ChartsScreenerView{},
			24,
			6,
			[]string{"Ticker", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
		{
			"cassettes/charts_line_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
				"chart_type":     Line,
				"timeframe":      Monthly,
			},
			&ChartsScreenerView{},
			24,
			6,
			[]string{"Ticker", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
		{
			"cassettes/basic_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&BasicScreenerView{},
			10,
			29,
			[]string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range"},
		},
		{
			"cassettes/news_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&NewsScreenerView{},
			10,
			30,
			[]string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "News"},
		},
		{
			"cassettes/description_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&DescriptionScreenerView{},
			10,
			30,
			[]string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "Description"},
		},
		{
			"cassettes/snapshot_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&SnapshotScreenerView{},
			10,
			32,
			[]string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "News", "Description", "Insider Trading"},
		},
		{
			"cassettes/ta_screener_view",
			map[string]interface{}{
				"signal": UnusualVolume,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&TAScreenerView{},
			10,
			29,
			[]string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "Perf Week", "Beta", "Perf Month", "ATR", "Perf Quarter", "Volatility W", "Perf Half Y", "Volatility M", "Perf Year", "SMA20", "Perf YTD", "SMA50", "RSI (14)", "SMA200", "Change Open", "52W High", "Gap", "52W Low", "Rel Volume", "Short Float", "Avg Volume", "Candlestick", "52W Range"},
		},
		{
			"cassettes/tickers_screener_view",
			map[string]interface{}{
				"signal": AllStocks,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&TickersScreenerView{},
			1000,
			7,
			[]string{"Ticker", "Change", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
		{
			"cassettes/bulk_screener_view",
			map[string]interface{}{
				"signal": AllStocks,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&BulkScreenerView{},
			500,
			7,
			[]string{"Ticker", "Change", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
		{
			"cassettes/bulk_full_screener_view",
			map[string]interface{}{
				"signal": AllStocks,
				"filters": []FilterInterface{
					ExchangeFilter(NYSE),
				},
				"tickers":        []string{},
				"general_order":  Descending,
				"specific_order": Volume,
			},
			&BulkFullScreenerView{},
			500,
			10,
			[]string{"Ticker", "Change", "Price", "Volume", "Relative Volume", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
	}

	for _, testInput := range testInputs {
		r, err := recorder.New(testInput.cassettePath)
		if err != nil {
			t.Error(err)
		}

		client := finviz.NewTestingClient(r)

		// Scraping Test
		df, err := GetScreenerData(client, testInput.viewInterfaceType, &testInput.viewArgs)
		if err != nil {
			t.Errorf("GetData function failed. Error: %v", err)
		}

		// Check row count
		if df.Nrow() != testInput.expectedDataFrameRowCount {
			t.Fail()
			t.Logf("Expected %v rows, received %v rows", testInput.expectedDataFrameRowCount, df.Nrow())
		}

		// Check column count
		if df.Ncol() != testInput.expectedDataFrameColumnCount {
			t.Fail()
			t.Logf("Expected %v columns, received %v columns", testInput.expectedDataFrameColumnCount, df.Ncol())
		}

		// Check column names
		receivedColumns := df.Names()
		for i := 0; i < testInput.expectedDataFrameColumnCount; i++ {
			if receivedColumns[i] != testInput.expectedDataFrameColumnNames[i] {
				t.Fail()
				t.Logf("Expected column: \"%v\", received column: \"%v\"", testInput.expectedDataFrameColumnNames[i], receivedColumns[i])
			}
		}

		if err := r.Stop(); err != nil {
			t.Error(err)
		}
	}
}
