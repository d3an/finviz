// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"fmt"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-gota/gota/dataframe"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// NewClient generates a new client instance
func NewClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}

// newTestingClient generates a new testing client instance that uses go-vcr
func newTestingClient(r *recorder.Recorder) *http.Client {
	return &http.Client{
		Timeout:   30 * time.Second,
		Transport: r,
	}
}

// MakeGetRequest is used to get a byte array of the screen given a valid URL
func MakeGetRequest(c *http.Client, url string) ([]byte, error) {
	// Set up GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Make GET request
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle unsuccessful GET requests
	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return b, StatusCodeError(fmt.Sprintf("Received HTTP response status %v: %v", resp.StatusCode, resp.Status))
	}

	// Convert the response body to a string
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return html, nil
}

// GetStockDataframe consumes an html instance and returns a dataframe of stocks returned
func GetStockDataframe(html interface{}, view ViewInterface) (*dataframe.DataFrame, error) {
	doc, err := generateDocument(html)
	if err != nil {
		return nil, err
	}

	results, err := view.Scrape(doc)
	if err != nil {
		return nil, err
	}

	df := dataframe.LoadRecords(results)
	return &df, nil
}

// GetViewFactory consumes a view query string and returns the associated ViewInterface
func GetViewFactory(viewQuery string) (ViewInterface, error) {
	switch strings.ToLower(viewQuery) {
	default:
		return &OverviewView{ViewType{"110"}}, ViewNotFoundError(fmt.Sprintf("View \"%v\" not found. ", viewQuery))
	case "overview":
		return &OverviewView{ViewType{"110"}}, nil
	case "valuation":
		return &ValuationView{ViewType{"120"}}, nil
	case "ownership":
		return &OwnershipView{ViewType{"130"}}, nil
	case "performance":
		return &PerformanceView{ViewType{"140"}}, nil
	case "custom":
		return &CustomView{ViewType{"150"}}, nil
	case "financial":
		return &FinancialView{ViewType{"160"}}, nil
	case "technical":
		return &TechnicalView{ViewType{"170"}}, nil
	case "charts":
		return &ChartsView{ViewType{"210"}}, nil
	case "basic":
		return &BasicView{ViewType{"310"}}, nil
	case "news":
		return &NewsView{ViewType{"320"}}, nil
	case "description":
		return &DescriptionView{ViewType{"330"}}, nil
	case "snapshot":
		return &SnapshotView{ViewType{"340"}}, nil
	case "ta":
		return &TAView{ViewType{"350"}}, nil
	case "tickers":
		return &TickersView{ViewType{"410"}}, nil
	case "bulk":
		return &BulkView{ViewType{"510"}}, nil
	case "fullbulk":
		return &BulkFullView{ViewType{"520"}}, nil
	}
}
