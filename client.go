// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"io/ioutil"
	"net/http"
	"time"
)

// StatusCodeError is the error given if a request's status code is not 200
type StatusCodeError string

func (err StatusCodeError) Error() string {
	return string(err)
}

// NoStocksMatchedQueryError is the error given if a screen returns no results
type NoStocksMatchedQueryError string

func (err NoStocksMatchedQueryError) Error() string {
	return string(err)
}

// NewClient generates a new client instance
func NewClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
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

// ScrapeScreenResults scrapes an HTML document for the screen's ticker results
func ScrapeScreenResults(doc *goquery.Document) ([][]string, error) {
	var row []string
	var rows [][]string

	// Only collects data from the v=111 view (Note: maximum 20 stocks are returned)
	doc.Find("[bgcolor=\"#d3d3d3\"]").Each(func(i int, tableHTML *goquery.Selection) {
		tableHTML.Find("tbody").Each(func(j int, tbodyHTML *goquery.Selection) {
			tbodyHTML.Find("[align=\"center\"]").Each(func(k int, rowHTML *goquery.Selection) {
				rowHTML.Find("td").Each(func(l int, tableCell *goquery.Selection) {
					row = append(row, tableCell.Text())
				})
				rows = append(rows, row)
				row = nil
			})
			tbodyHTML.Find("[valign=\"top\"]").Each(func(k int, rowHTML *goquery.Selection) {
				rowHTML.Find("td").Each(func(l int, tableCell *goquery.Selection) {
					row = append(row, tableCell.Text())
				})
				rows = append(rows, row)
				row = nil
			})
		})
	})

	// Check if only column titles exist
	if len(rows) == 1 {
		return nil, NoStocksMatchedQueryError("The screen returned no tickers. Try diluting your search.")
	}

	return rows, nil
}

// GetStockDataframe consumes a byte array of the html request and returns a dataframe of stocks returned
func GetStockDataframe(html []byte) (*dataframe.DataFrame, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, err
	}

	rows, err := ScrapeScreenResults(doc)
	if err != nil {
		return nil, err
	}

	// Declare dataframe for further analysis
	df := dataframe.LoadRecords(rows)

	return &df, nil
}

// GetOrderedTickerList is a helper function for viewing tickers from a screen dataframe
func GetOrderedTickerList(df *dataframe.DataFrame) []string {
	var tickers []string
	column := df.Col("Ticker")
	colLen := column.Len()

	for i := 0; i < colLen; i++ {
		tickers = append(tickers, column.Val(i).(string))
	}

	return tickers
}
