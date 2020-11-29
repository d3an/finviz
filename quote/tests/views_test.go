// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package tests

import (
	"fmt"
	"github.com/d3an/finviz"
	"github.com/d3an/finviz/quote"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-gota/gota/dataframe"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuoteView_GenerateURL(t *testing.T) {
	testInputs := []struct {
		viewArgs          map[string]interface{}
		viewInterfaceType finviz.ViewInterface
		expectedURL       string
	}{
		{
			map[string]interface{}{
				"ticker": "aapl",
			},
			&quote.QuoteView{},
			fmt.Sprintf("%v/quote.ashx?t=AAPL&ty=c&p=d&b=1", finviz.APIURL),
		},
		{
			map[string]interface{}{
				"ticker": "A",
			},
			&quote.QuoteView{},
			fmt.Sprintf("%v/quote.ashx?t=A&ty=c&p=d&b=1", finviz.APIURL),
		},
	}

	for i, ti := range testInputs {
		if url, err := ti.viewInterfaceType.GenerateURL(&ti.viewArgs); err != nil {
			if i != 0 {
				t.Errorf("GenerateURL failed. Error: %v", err)
			}
		} else if url != ti.expectedURL {
			t.Fail()
			t.Logf("URL Generation failed. Expected: \"%v\", Received: \"%v\"", ti.expectedURL, url)
		}
	}
}

func TestQuoteView_Scrape(t *testing.T) {
	testInputs := []struct {
		cassettePath      string
		viewInterfaceType finviz.ViewInterface
		viewArgs          *map[string]interface{}
		columnCount       int
		missingColumns    []string
	}{
		{
			"cassettes/full_quote",
			&quote.QuoteView{},
			&map[string]interface{}{
				"ticker": "AAPL", // Full column count
			},
			83,
			[]string{},
		},
		{
			"cassettes/missing_insdr_and_recom",
			&quote.QuoteView{},
			&map[string]interface{}{
				"ticker": "ATHE", // No Insider Trading or Analyst Recommendation table
			},
			81,
			[]string{"Insider Trading", "Analyst Recommendations"},
		},
		{
			"cassettes/missing_insdr",
			&quote.QuoteView{},
			&map[string]interface{}{
				"ticker": "AEZS", // No Insider Trading table
			},
			82,
			[]string{"Insider Trading"},
		},
	}

	for _, ti := range testInputs {
		r, err := recorder.New(ti.cassettePath)
		if err != nil {
			t.Error(err)
		}

		url, err := ti.viewInterfaceType.GenerateURL(ti.viewArgs)
		if err != nil {
			t.Error(err)
		}

		html, err := finviz.MakeGetRequest(r, url)
		if err != nil {
			t.Error(err)
		}

		doc, err := finviz.GenerateDocument(html)
		if err != nil {
			t.Error(err)
		}

		results, err := ti.viewInterfaceType.Scrape(doc)
		if err != nil {
			t.Fail()
			t.Log(err)
		}

		df := dataframe.LoadRecords(results)

		for name := range ti.missingColumns {
			assert.NotContains(t, df.Names(), name)
		}

		assert.Equal(t, ti.columnCount, df.Ncol())

		if err := r.Stop(); err != nil {
			t.Error(err)
		}
	}
}

func TestQuoteView_MapScrape(t *testing.T) {
	testInputs := []struct {
		cassettePath string
		viewArgs     *map[string]interface{}
	}{
		{
			"cassettes/full_quote",
			&map[string]interface{}{
				"ticker": "AAPL", // Full column count
			},
		},
		{
			"cassettes/missing_insdr_and_recom",
			&map[string]interface{}{
				"ticker": "ATHE", // No Insider Trading or Analyst Recommendation table
			},
		},
		{
			"cassettes/missing_insdr",
			&map[string]interface{}{
				"ticker": "AEZS", // No Insider Trading table
			},
		},
	}

	for _, ti := range testInputs {
		r, err := recorder.New(ti.cassettePath)
		if err != nil {
			t.Error(err)
		}

		view := quote.QuoteView{}

		url, err := view.GenerateURL(ti.viewArgs)
		if err != nil {
			t.Error(err)
		}

		html, err := finviz.MakeGetRequest(r, url)
		if err != nil {
			t.Error(err)
		}

		doc, err := finviz.GenerateDocument(html)
		if err != nil {
			t.Error(err)
		}

		resultMap, err := view.MapScrape(doc)
		if err != nil {
			t.Fail()
			t.Log(err)
		}

		assert.Equal(t, 83, len(*resultMap))
	}
}

// Tested retry system with 155 tickers; a couple 403s, but resolved after waiting 1ms. Total time: 1.26s
// Test not included due to VCR file being too big for GitHub.
func TestGetQuoteData(t *testing.T) {
	testInputs := []struct {
		cassettePath string
		viewArgs     *map[string]interface{}
		tickerCount  int
	}{
		{
			"cassettes/get_quote_data",
			&map[string]interface{}{
				"tickers": []string{"AAPL", "GOOG", "NFLX", "AMZN"},
			},
			4,
		},
	}

	for _, ti := range testInputs {
		r, err := recorder.New(ti.cassettePath)
		if err != nil {
			t.Error(err)
		}

		df, err := quote.GetQuoteData(r, ti.viewArgs)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 83, df.Ncol())
		assert.Equal(t, ti.tickerCount, df.Nrow())

		if err := r.Stop(); err != nil {
			t.Error(err)
		}
	}
}
