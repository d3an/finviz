// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package tests

import (
	"fmt"
	"github.com/d3an/finviz"
	"testing"
)

func TestView_GenerateURL(t *testing.T) {
	v := finviz.View{}
	url, err := v.GenerateURL(nil)
	if err != nil {
		t.Fail()
		t.Log(err)
	}

	if url != finviz.APIURL {
		t.Fail()
		t.Logf(fmt.Sprintf("Expected: \"%v\", Received: \"%v\"", finviz.APIURL, url))
	}
}

func TestView_Scrape(t *testing.T) {
	v := finviz.View{}
	if _, err := v.Scrape(nil); err == nil {
		t.Fail()
		t.Log("Expected function to throw \"MethodNotImplementedError\"")
	}
}