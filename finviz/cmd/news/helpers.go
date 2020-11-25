// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package news

import (
	"fmt"
	"github.com/d3an/finviz"
	"github.com/d3an/finviz/news"
	"os"
	"strings"
)

// ViewFactory consumes a view query string and returns the associated ViewInterface
func ViewFactory(viewQuery string) (finviz.ViewInterface, error) {
	switch strings.ToLower(viewQuery) {
	default:
		return nil, fmt.Errorf("view \"%v\" is not supported", viewQuery)
	case "time", "1":
		return &news.TimeNewsView{}, nil
	case "source", "2":
		return &news.SourceNewsView{}, nil
	}
}

func er(msg interface{}) {
	fmt.Println("Error: ", msg)
	os.Exit(1)
}
