// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package quote

/*
func getTickersSlice(viewArgs *map[string]interface{}) ([]string, error) {
	if value, exists := (*viewArgs)["tickers"]; exists {
		if value, ok := value.([]string); ok {
			return value, nil
		}
		return nil, fmt.Errorf("invalid type for \"tickers\" argument")
	}
	return nil, fmt.Errorf("\"tickers\" argument not found")
}

func buildQuoteData(wg *sync.WaitGroup, rec *recorder.Recorder, url string, c *chan interface{}) {
	defer wg.Done()
	defer close(*c)

	var html []byte
	var err error
	waitTime := 1 * time.Millisecond

	for {
		html, err = finviz.MakeGetRequest(rec, url)
		if err == nil {
			break
		}
		time.Sleep(waitTime)
		waitTime *= 2
		if waitTime > 10000*time.Millisecond {
			*c <- err
			return
		}
	}

	doc, err := finviz.GenerateDocument(html)
	if err != nil {
		*c <- err
		return
	}

	var view QuoteView
	results, err := view.MapScrape(doc)
	if err != nil {
		*c <- err
		return
	}

	*c <- results
}

// GetQuoteData returns a DataFrame with screener data results
func GetQuoteData(rec *recorder.Recorder, viewArgs *map[string]interface{}) (*dataframe.DataFrame, error) {
	tickers, err := getTickersSlice(viewArgs)
	if err != nil {
		return nil, err
	}

	tickerCount := len(tickers)
	c := make([]chan interface{}, tickerCount)
	for i := range c {
		c[i] = make(chan interface{}, 1)
	}

	var wg sync.WaitGroup
	var view QuoteView

	for i, t := range tickers {
		(*viewArgs)["ticker"] = t
		url, err := view.GenerateURL(viewArgs)
		if err != nil {
			return nil, err
		}

		wg.Add(1)
		go buildQuoteData(&wg, rec, url, &c[i])
	}

	wg.Wait()

	var results []map[string]interface{}
	var errors []error

	for i := 0; i < tickerCount; i++ {
		switch j := (<-c[i]).(type) {
		default:
			return nil, fmt.Errorf("invalid type in channel, value: \"%v\"", i)
		case error:
			errors = append(errors, j)
		case *map[string]interface{}:
			results = append(results, *j)
		}
	}

	// Put this in a global variable
	quoteHeaders := []string{"Ticker", "Company", "Industry", "Sector", "Country", "Index", "Market Cap", "Price", "Change", "Volume", "Income", "Sales", "Book/sh", "Cash/sh", "Dividend", "Dividend %", "Employees", "Optionable", "Shortable", "Recom", "P/E", "Forward P/E", "PEG", "P/S", "P/B", "P/C", "P/FCF", "Quick Ratio", "Current Ratio", "Debt/Eq", "LT Debt/Eq", "EPS (ttm)", "EPS next Y", "EPS next Q", "EPS this Y", "EPS growth next Y", "EPS next 5Y", "EPS past 5Y", "Sales past 5Y", "Sales Q/Q", "EPS Q/Q", "Earnings", "Insider Own", "Insider Trans", "Inst Own", "Inst Trans", "ROA", "ROE", "ROI", "Gross Margin", "Oper. Margin", "Profit Margin", "Payout", "Shs Outstand", "Shs Float", "Short Float", "Short Ratio", "Target Price", "52W Range", "52W High", "52W Low", "RSI (14)", "SMA20", "SMA50", "SMA200", "Rel Volume", "Avg Volume", "Perf Week", "Perf Month", "Perf Quarter", "Perf Half Y", "Perf Year", "Perf YTD", "Beta", "ATR", "Volatility (Week)", "Volatility (Month)", "Prev Close", "Chart", "Analyst Recommendations", "News", "Description", "Insider Trading"}

	orderedRows, err := utils.GenerateRows(quoteHeaders, results)
	if err != nil {
		return nil, fmt.Errorf("row generation failed")
	}

	df := dataframe.LoadRecords(orderedRows)

	if len(errors) > 0 {
		errMsg := fmt.Errorf("(1) %v", errors[0])
		for j, err := range errors {
			if j != 0 {
				errMsg = fmt.Errorf("%v\n(%v) %v", errMsg, j+1, err)
			}
		}
		return &df, errMsg
	}

	return &df, nil
}
*/
