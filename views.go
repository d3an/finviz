// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"github.com/PuerkitoBio/goquery"
)

// ViewInterface is the base interface for FinViz views
type ViewInterface interface {
	GenerateURL(viewArgs *map[string]interface{}) (string, error)
	Scrape(doc *goquery.Document) ([][]string, error)
}

// View is the base class for FinViz views
type View struct{}

// GenerateURL is the default method for FinViz views
func (v *View) GenerateURL(_ *map[string]interface{}) (string, error) {
	return APIURL, nil
}

// Scrape is the default method for FinViz views
func (v *View) Scrape(_ *goquery.Document) ([][]string, error) {
	return make([][]string, 0), MethodNotImplementedError("Scrape method not implemented")
}
