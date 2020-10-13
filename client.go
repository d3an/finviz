// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
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

// NewTestingClient generates a new testing client instance that uses go-vcr
func NewTestingClient(rec *recorder.Recorder) *http.Client {
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
		return b, StatusCodeError(fmt.Sprintf("received HTTP response status %v: %v", resp.StatusCode, resp.Status))
	}

	// Convert the response body to a string
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return html, nil
}

// GenerateDocument is a helper function for Scraping
func GenerateDocument(html interface{}) (doc *goquery.Document, err error) {
	switch html := html.(type) {
	default:
		return nil, fmt.Errorf("HTML object type is not one of string, []byte, or io.ReadCloser")
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
		return GenerateDocument(byteArray)
	}
	return doc, nil
}

// GenerateRows is a helper function for DataFrame construction
func GenerateRows(headers []string, dataSlice []map[string]interface{}) (rows [][]string, err error) {
	headerCount := len(headers)
	resultCount := len(dataSlice)

	rows = append(rows, headers)
	for i := 0; i < resultCount; i++ {
		var row []string

		for j := 0; j < headerCount; j++ {
			item := dataSlice[i][headers[j]]
			switch item := item.(type) {
			default:
				return nil, fmt.Errorf("unexpected type for #%v: %v -> %v", i, headers[j], dataSlice[i][headers[j]])
			case nil:
				row = append(row, "-")
			case string:
				row = append(row, item)
			case []map[string]string:
				news, err := json.Marshal(item)
				if err != nil {
					return nil, err
				}
				row = append(row, string(news))
			}
		}
		rows = append(rows, row)
	}

	return rows, nil
}
