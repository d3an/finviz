// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package quote

import (
	"fmt"
	"github.com/d3an/finviz"
	"github.com/go-gota/gota/dataframe"
	"net/http"
	"sync"
	"time"
)

func getTickerString(viewArgs *map[string]interface{}) (string, error) {
	if value, exists := (*viewArgs)["ticker"]; exists {
		if value, ok := value.(string); ok {
			return value, nil
		}
		return "", fmt.Errorf("\"ticker\" argument not of type \"string\"")
	}
	return "", fmt.Errorf("\"ticker\" argument not found")
}

func getTickersSlice(viewArgs *map[string]interface{}) ([]string, error) {
	if value, exists := (*viewArgs)["tickers"]; exists {
		if value, ok := value.([]string); ok {
			return value, nil
		}
		return nil, fmt.Errorf("invalid type for \"tickers\" argument")
	}
	return nil, fmt.Errorf("\"tickers\" argument not found")
}

func buildQuoteData(wg *sync.WaitGroup, client *http.Client, url string, c *chan interface{}) {
	defer wg.Done()
	defer close(*c)

	html, err := finviz.MakeGetRequest(client, url)
	if err != nil {
		*c <- err
		return
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
func GetQuoteData(client *http.Client, viewArgs *map[string]interface{}) (*dataframe.DataFrame, error) {
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

		fmt.Println(url)

		wg.Add(1)
		go buildQuoteData(&wg, client, url, &c[i])
		time.Sleep(500 * time.Millisecond)
	}

	wg.Wait()

	var results []map[string]interface{}
	var errors []error
	var headers []string

	for i := 0; i < tickerCount; i++ {
		switch j := (<-c[i]).(type) {
		default:
			return nil, fmt.Errorf("invalid type in channel, value: \"%v\"", i)
		case error:
			errors = append(errors, j)
		case *map[string]interface{}:
			if i == 0 {
				for key := range *j {
					headers = append(headers, key)
				}
			}
			results = append(results, *j)
		}
	}

	orderedRows, err := finviz.GenerateRows(headers, results)
	if err != nil {
		return nil, fmt.Errorf("row generation failed")
	}

	df := dataframe.LoadRecords(orderedRows)
	fmt.Println(df)

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
