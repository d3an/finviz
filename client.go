// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-gota/gota/dataframe"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// HeaderTransport implements a Transport that can have its RoundTripper interface modified
type HeaderTransport struct {
	T http.RoundTripper
}

// RoundTrip implements the RoundTripper interface with a custom user-agent
func (adt *HeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", uarand.GetRandom())
	return adt.T.RoundTrip(req)
}

func addHeaderTransport(t http.RoundTripper) *HeaderTransport {
	if t == nil {
		t = http.DefaultTransport
	}
	return &HeaderTransport{t}
}

// NewClient generates a new client instance
func NewClient() *http.Client {
	return &http.Client{
		Timeout:   30 * time.Second,
		Transport: addHeaderTransport(nil),
	}
}

// newTestingClient generates a new testing client instance that uses go-vcr
func newTestingClient(rec *recorder.Recorder) *http.Client {
	return &http.Client{
		Timeout:   30 * time.Second,
		Transport: addHeaderTransport(rec),
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

func generateDocument(html interface{}) (doc *goquery.Document, err error) {
	switch html := html.(type) {
	default:
		return nil, fmt.Errorf("HTML object is not of type string or []byte or io.ReadCloser")
	case string:
		html = strings.ReplaceAll(html, "\\r", "")
		html = strings.ReplaceAll(html, "\\n", "")
		html = strings.ReplaceAll(html, "\\\"", "\"")

		html = strings.Map(func(r rune) rune {
			if r == '\n' || r == '\t' {
				return ' '
			}
			return r
		}, html)
		doc, err = goquery.NewDocumentFromReader(bytes.NewReader([]byte(html)))
		if err != nil {
			return nil, err
		}
	case []byte:
		doc, err = goquery.NewDocumentFromReader(bytes.NewReader(html))
		if err != nil {
			return nil, err
		}

	case io.ReadCloser:
		byteArray, err := ioutil.ReadAll(html)
		if err != nil {
			return nil, err
		}
		return generateDocument(byteArray)
	}
	return doc, nil
}

// GetStockDataframe consumes an html instance and returns a dataframe of stocks returned
func GetStockDataframe(html, view interface{}) (*dataframe.DataFrame, error) {
	doc, err := generateDocument(html)
	if err != nil {
		return nil, err
	}

	var results [][]string
	switch view := view.(type) {
	default:
		return nil, InvalidViewError("view was not initialized as a ViewInterface or ChartViewInterface")
	case ChartViewInterface:
		results, err = view.Scrape(doc)
		if err != nil {
			return nil, err
		}
	case ViewInterface:
		results, err = view.Scrape(doc)
		if err != nil {
			return nil, err
		}
	}

	df := dataframe.LoadRecords(results)
	return &df, nil
}

// GetViewFactory consumes a view query string and returns the associated ViewInterface
func GetViewFactory(viewQuery string) (interface{}, error) {
	switch strings.ToLower(viewQuery) {
	default:
		return &OverviewView{ViewType{"110"}}, InvalidViewError(fmt.Sprintf("view \"%v\" is not supported", viewQuery))
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
		return &ChartsView{ChartViewType{"210", Technical, Daily}}, nil
	case "basic":
		return &BasicView{ChartViewType{"310", Technical, Daily}}, nil
	case "news":
		return &NewsView{ChartViewType{"320", Technical, Daily}}, nil
	case "description":
		return &DescriptionView{ChartViewType{"330", Technical, Daily}}, nil
	case "snapshot":
		return &SnapshotView{ChartViewType{"340", Technical, Daily}}, nil
	case "ta":
		return &TAView{ChartViewType{"350", Technical, Daily}}, nil
	case "tickers":
		return &TickersView{ViewType{"410"}}, nil
	case "bulk":
		return &BulkView{ViewType{"510"}}, nil
	case "bulkfull":
		return &BulkFullView{ViewType{"520"}}, nil
	}
}
