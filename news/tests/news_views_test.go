// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package tests

import (
	"fmt"
	"github.com/d3an/finviz"
	"github.com/d3an/finviz/news"
	"github.com/dnaeon/go-vcr/recorder"
	"testing"
)

func TestNewsView_GenerateURL(t *testing.T) {
	testInputs := []struct {
		viewInterfaceType finviz.ViewInterface
		expectedURL       string
	}{
		{
			&news.NewsView{},
			fmt.Sprintf("%v/news.ashx", finviz.APIURL),
		},
		{
			&news.TimeNewsView{},
			fmt.Sprintf("%v/news.ashx", finviz.APIURL),
		},
		{
			&news.SourceNewsView{},
			fmt.Sprintf("%v/news.ashx?v=2", finviz.APIURL),
		},
	}

	for _, testInput := range testInputs {
		if url, err := testInput.viewInterfaceType.GenerateURL(nil); err != nil {
			t.Errorf("GenerateURL failed. Error: %v", err)
		} else if url != testInput.expectedURL {
			t.Fail()
			t.Logf("URL Generation failed. Expected: \"%v\", Received: \"%v\"", testInput.expectedURL, url)
		}
	}
}

func TestNewsView_Scrape(t *testing.T) {
	testInputs := []struct {
		cassettePath        string
		viewInterface       finviz.ViewInterface
		expectedColumnNames []string
	}{
		{
			"cassettes/source_news_view",
			&news.SourceNewsView{},
			[]string{"Article Date", "Article Title", "Article URL", "Source Name", "Source URL", "News Type"},
		},
		{
			"cassettes/time_news_view",
			&news.TimeNewsView{},
			[]string{"Article Date", "Article Title", "Article URL", "Source Name", "News Type"},
		},
	}

	for _, testInput := range testInputs {
		r, err := recorder.New(testInput.cassettePath)
		if err != nil {
			t.Error(err)
		}

		// Scraping test
		df, err := news.GetNewsData(r, testInput.viewInterface)
		if err != nil {
			t.Error(err)
		}

		expectedColumnCount := len(testInput.expectedColumnNames)
		receivedColumnNames := df.Names()
		receivedColumnCount := len(receivedColumnNames)

		// Check column count
		if receivedColumnCount != expectedColumnCount {
			t.Fail()
			t.Logf("Expected column count: \"%v\", Received column count: \"%v\"", expectedColumnCount, receivedColumnCount)
		}

		// Check column names
		for i := 0; i < expectedColumnCount; i++ {
			if receivedColumnNames[i] != testInput.expectedColumnNames[i] {
				t.Fail()
				t.Logf("Expected column name: \"%v\", Received column name: \"%v\"", testInput.expectedColumnNames[i], receivedColumnNames[i])
			}
		}

		// Check minimum row count
		if df.Nrow() < 100 {
			t.Fail()
			t.Logf("Expected at least 100 rows, Received %v rows", df.Nrow())
		}

		if err := r.Stop(); err != nil {
			t.Error(err)
		}

		finviz.PrintFullDataFrame(df)
	}
}
