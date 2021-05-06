// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"

	"github.com/d3an/finviz/utils"
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

// MakeGetRequest is used to get a byte array of the screen given a valid URL
func MakeGetRequest(rec *recorder.Recorder, url string) ([]byte, error) {
	c := &http.Client{
		Timeout:   30 * time.Second,
		Transport: addHeaderTransport(rec),
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Handle unsuccessful GET requests
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, utils.StatusCodeError(fmt.Sprintf("received HTTP response status %v: %v", resp.StatusCode, resp.Status))
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
		return nil, fmt.Errorf("HTML object type is not 'string', '[]byte', or 'io.ReadCloser'")
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
