// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package screener

import (
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-gota/gota/series"
	"github.com/stretchr/testify/require"

	"github.com/d3an/finviz/utils"
)

func newTestClient(config *Config) *Client {
	if config != nil {
		return &Client{
			Client: &http.Client{
				Timeout:   30 * time.Second,
				Transport: utils.AddHeaderTransport(config.recorder),
			},
			config: *config,
		}
	}
	return &Client{
		Client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 30 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout: 30 * time.Second,
			},
		},
		config: Config{userAgent: uarand.GetRandom()},
	}
}

func TestGetScreenerResultsLotOfPages(t *testing.T) {
	values := []struct {
		cassettePath     string
		url              string
		expectedColCount int
		expectedColNames []string
	}{
		{
			cassettePath:     "cassettes/custom_all",
			url:              "https://finviz.com/screener.ashx?v=150&c=0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70",
			expectedColCount: 71,
			expectedColNames: []string{"No.", "Ticker", "Company", "Sector", "Industry", "Country", "Market Cap", "P/E", "Fwd P/E", "PEG", "P/S", "P/B", "P/C", "P/FCF", "Dividend", "Payout Ratio", "EPS", "EPS this Y", "EPS next Y", "EPS past 5Y", "EPS next 5Y", "Sales past 5Y", "EPS Q/Q", "Sales Q/Q", "Outstanding", "Float", "Insider Own", "Insider Trans", "Inst Own", "Inst Trans", "Float Short", "Short Ratio", "ROA", "ROE", "ROI", "Curr R", "Quick R", "LTDebt/Eq", "Debt/Eq", "Gross M", "Oper M", "Profit M", "Perf Week", "Perf Month", "Perf Quart", "Perf Half", "Perf Year", "Perf YTD", "Beta", "ATR", "Volatility W", "Volatility M", "SMA20", "SMA50", "SMA200", "50D High", "50D Low", "52W High", "52W Low", "RSI", "from Open", "Gap", "Recom", "Avg Volume", "Rel Volume", "Price", "Change", "Volume", "Earnings", "Target Price", "IPO Date"},
		},
	}

	for _, v := range values {
		func() {
			r, err := recorder.New(v.cassettePath)
			require.Nil(t, err)
			defer func() {
				err = r.Stop()
				require.Nil(t, err)
			}()

			client := newTestClient(&Config{recorder: r, userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"})

			df, err := client.GetScreenerResults(v.url)
			require.Nil(t, err)
			utils.PrintFullDataFrame(df)
			require.Equal(t, v.expectedColCount, df.Ncol())
			require.Equal(t, v.expectedColNames, df.Names())
		}()
	}
}

func TestGetScreenerResults(t *testing.T) {
	values := []struct {
		cassettePath     string
		url              string
		expectedRowCount int
		expectedColCount int
		expectedColNames []string
	}{
		{
			cassettePath:     "cassettes/overview",
			url:              "https://finviz.com/screener.ashx?v=110&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 37,
			expectedColCount: 11,
			expectedColNames: []string{"No.", "Ticker", "Company", "Sector", "Industry", "Country", "Market Cap", "P/E", "Price", "Change", "Volume"},
		},
		{
			cassettePath:     "cassettes/valuation",
			url:              "https://finviz.com/screener.ashx?v=120&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 89,
			expectedColCount: 18,
			expectedColNames: []string{"No.", "Ticker", "Market Cap", "P/E", "Fwd P/E", "PEG", "P/S", "P/B", "P/C", "P/FCF", "EPS this Y", "EPS next Y", "EPS past 5Y", "EPS next 5Y", "Sales past 5Y", "Price", "Change", "Volume"},
		},
		{
			cassettePath:     "cassettes/ownership",
			url:              "https://finviz.com/screener.ashx?v=130&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 20,
			expectedColCount: 15,
			expectedColNames: []string{"No.", "Ticker", "Market Cap", "Outstanding", "Float", "Insider Own", "Insider Trans", "Inst Own", "Inst Trans", "Float Short", "Short Ratio", "Avg Volume", "Price", "Change", "Volume"},
		},
		{
			cassettePath:     "cassettes/performance",
			url:              "https://finviz.com/screener.ashx?v=140&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 20,
			expectedColCount: 16,
			expectedColNames: []string{"No.", "Ticker", "Perf Week", "Perf Month", "Perf Quart", "Perf Half", "Perf Year", "Perf YTD", "Volatility W", "Volatility M", "Recom", "Avg Volume", "Rel Volume", "Price", "Change", "Volume"},
		},
		{
			cassettePath:     "cassettes/custom",
			url:              "https://finviz.com/screener.ashx?v=150&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume&c=0,1,3,2,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70",
			expectedRowCount: 20,
			expectedColCount: 71,
			expectedColNames: []string{"No.", "Ticker", "Sector", "Company", "Industry", "Country", "Market Cap", "P/E", "Fwd P/E", "PEG", "P/S", "P/B", "P/C", "P/FCF", "Dividend", "Payout Ratio", "EPS", "EPS this Y", "EPS next Y", "EPS past 5Y", "EPS next 5Y", "Sales past 5Y", "EPS Q/Q", "Sales Q/Q", "Outstanding", "Float", "Insider Own", "Insider Trans", "Inst Own", "Inst Trans", "Float Short", "Short Ratio", "ROA", "ROE", "ROI", "Curr R", "Quick R", "LTDebt/Eq", "Debt/Eq", "Gross M", "Oper M", "Profit M", "Perf Week", "Perf Month", "Perf Quart", "Perf Half", "Perf Year", "Perf YTD", "Beta", "ATR", "Volatility W", "Volatility M", "SMA20", "SMA50", "SMA200", "50D High", "50D Low", "52W High", "52W Low", "RSI", "from Open", "Gap", "Recom", "Avg Volume", "Rel Volume", "Price", "Change", "Volume", "Earnings", "Target Price", "IPO Date"},
		},
		{
			cassettePath:     "cassettes/financial",
			url:              "https://finviz.com/screener.ashx?v=160&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 20,
			expectedColCount: 18,
			expectedColNames: []string{"No.", "Ticker", "Market Cap", "Dividend", "ROA", "ROE", "ROI", "Curr R", "Quick R", "LTDebt/Eq", "Debt/Eq", "Gross M", "Oper M", "Profit M", "Earnings", "Price", "Change", "Volume"},
		},
		{
			cassettePath:     "cassettes/technical",
			url:              "https://finviz.com/screener.ashx?v=170&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 20,
			expectedColCount: 15,
			expectedColNames: []string{"No.", "Ticker", "Beta", "ATR", "SMA20", "SMA50", "SMA200", "52W High", "52W Low", "RSI", "Price", "Change", "from Open", "Gap", "Volume"},
		},
		{
			cassettePath:     "cassettes/charts",
			url:              "https://finviz.com/screener.ashx?v=210&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 12,
			expectedColCount: 6,
			expectedColNames: []string{"Ticker", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
		{
			cassettePath:     "cassettes/charts_candle",
			url:              "https://finviz.com/screener.ashx?v=210&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume&ta=0&p=w",
			expectedRowCount: 24,
			expectedColCount: 6,
			expectedColNames: []string{"Ticker", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
		{
			cassettePath:     "cassettes/charts_line",
			url:              "https://finviz.com/screener.ashx?v=210&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume&ty=l&ta=0&p=m",
			expectedRowCount: 24,
			expectedColCount: 6,
			expectedColNames: []string{"Ticker", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
		{
			cassettePath:     "cassettes/basic",
			url:              "https://finviz.com/screener.ashx?v=310&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 10,
			expectedColCount: 29,
			expectedColNames: []string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range"},
		},
		{
			cassettePath:     "cassettes/news",
			url:              "https://finviz.com/screener.ashx?v=320&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 10,
			expectedColCount: 30,
			expectedColNames: []string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "News"},
		},
		{
			cassettePath:     "cassettes/description",
			url:              "https://finviz.com/screener.ashx?v=330&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 10,
			expectedColCount: 30,
			expectedColNames: []string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "Description"},
		},
		{
			cassettePath:     "cassettes/snapshot",
			url:              "https://finviz.com/screener.ashx?v=340&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 10,
			expectedColCount: 32,
			expectedColNames: []string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "News", "Description", "Insider Trading"},
		},
		{
			cassettePath:     "cassettes/ta",
			url:              "https://finviz.com/screener.ashx?v=350&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume",
			expectedRowCount: 10,
			expectedColCount: 29,
			expectedColNames: []string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "Perf Week", "Beta", "Perf Month", "ATR", "Perf Quarter", "Volatility W", "Perf Half Y", "Volatility M", "Perf Year", "SMA20", "Perf YTD", "SMA50", "RSI (14)", "SMA200", "Change Open", "52W High", "Gap", "52W Low", "Rel Volume", "Short Float", "Avg Volume", "Candlestick", "52W Range"},
		},
		{
			cassettePath:     "cassettes/tickers",
			url:              "https://finviz.com/screener.ashx?v=410&f=exch_nyse&o=-volume",
			expectedRowCount: 1000,
			expectedColCount: 7,
			expectedColNames: []string{"Ticker", "Change", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
		{
			cassettePath: "cassettes/bulk",
			url:          "https://finviz.com/screener.ashx?v=510&f=exch_nyse&o=-volume",

			expectedRowCount: 500,
			expectedColCount: 7,
			expectedColNames: []string{"Ticker", "Change", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
		{
			cassettePath: "cassettes/bulkfull",
			url:          "https://finviz.com/screener.ashx?v=520&f=exch_nyse&o=-volume",

			expectedRowCount: 500,
			expectedColCount: 10,
			expectedColNames: []string{"Ticker", "Change", "Price", "Volume", "Relative Volume", "Chart", "Company", "Industry", "Country", "Market Cap"},
		},
	}

	for _, v := range values {
		func() {
			r, err := recorder.New(v.cassettePath)
			require.Nil(t, err)
			defer func() {
				err = r.Stop()
				require.Nil(t, err)
			}()

			client := newTestClient(&Config{recorder: r, userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"})

			df, err := client.GetScreenerResults(v.url)
			require.Nil(t, err)
			utils.PrintFullDataFrame(df)
			require.Equal(t, v.expectedColCount, df.Ncol())
			require.Equal(t, v.expectedColNames, df.Names())
		}()
	}
}

func TestCleanDataFrame(t *testing.T) {
	values := []struct {
		cassettePath string
		url          string
	}{
		{
			cassettePath: "cassettes/custom",
			url:          "https://finviz.com/screener.ashx?v=150&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume&c=0,1,3,2,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,57,58,59,60,61,62,63,64,65,66,67,68,69,70",
		},
		{
			cassettePath: "cassettes/performance2",
			url:          "https://finviz.com/screener.ashx?v=140&s=ta_topgainers&f=exch_nasd|nyse,sh_avgvol_o50,sh_price_o1&o=-changeopen",
		},
	}

	for _, v := range values {
		func() {
			r, err := recorder.New(v.cassettePath)
			require.Nil(t, err)
			defer func() {
				err = r.Stop()
				require.Nil(t, err)
			}()
			client := newTestClient(&Config{recorder: r, userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"})

			df, err := client.GetScreenerResults(v.url)
			require.Nil(t, err)

			columnNames := df.Names()
			columnTypes := df.Types()
			columnCount := len(columnNames)

			for i := 0; i < columnCount; i++ {
				val, exists := utils.ColumnTypeLookup[strings.ToLower(columnNames[i])]
				require.True(t, exists)
				switch val {
				case "int", "bigint", "commaint":
					require.Equal(t, series.Int, columnTypes[i])
				case "string":
					require.Equal(t, series.String, columnTypes[i])
				case "float", "percent":
					require.Equal(t, series.Float, columnTypes[i])
				}
			}
		}()
	}
}
