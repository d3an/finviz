// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
)

// Contains is a semi-generic (string/int) function for checking if an item exists within a list
func Contains(slice, value interface{}) bool {
	switch slice := slice.(type) {
	case []string:
		switch value := value.(type) {
		case string:
			for _, a := range slice {
				if a == value {
					return true
				}
			}
		case int:
			return false
		}
	case []int:
		switch value := value.(type) {
		case int:
			for _, a := range slice {
				if a == value {
					return true
				}
			}
		case string:
			return false
		}
	}
	return false
}

// ExportCSV generates a csv file of the Finviz screen results
func ExportCSV(df *dataframe.DataFrame, outFileName string) error {
	f, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	return df.WriteCSV(f)
}

// ExportJSON generates a json file of the Finviz screen results
func ExportJSON(df *dataframe.DataFrame, outFileName string) error {
	f, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	return df.WriteJSON(f)
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

// GenerateRows is a helper function for DataFrame construction
func GenerateRows(headers []string, tickerDataSlice []map[string]interface{}) (rows [][]string, err error) {
	headerCount := len(headers)
	resultCount := len(tickerDataSlice)

	rows = append(rows, headers)
	for i := 0; i < resultCount; i++ {
		var row []string

		for j := 0; j < headerCount; j++ {
			switch item := tickerDataSlice[i][headers[j]].(type) {
			default:
				return nil, fmt.Errorf("unexpected type for %v: %v -> %v", tickerDataSlice[i]["Ticker"], headers[j], tickerDataSlice[i][headers[j]])
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

// PrintFullDataFrame prints an entire dataframe to console
// Derived from https://github.com/go-gota/gota/blob/master/dataframe/dataframe.go until print method is made public
func PrintFullDataFrame(df *dataframe.DataFrame) {
	nrows, ncols := df.Dims()
	if df.Err != nil || nrows == 0 || ncols == 0 {
		return
	}

	var dfStr string
	records := df.Records()
	typesrow := make([]string, ncols)
	maxChars := make([]int, ncols+1)

	addRightPadding := func(s string, nchar int) string {
		if utf8.RuneCountInString(s) < nchar {
			return s + strings.Repeat(" ", nchar-utf8.RuneCountInString(s))
		}
		return s
	}

	addLeftPadding := func(s string, nchar int) string {
		if utf8.RuneCountInString(s) < nchar {
			return strings.Repeat(" ", nchar-utf8.RuneCountInString(s)) + s
		}
		return s
	}

	dfStr += fmt.Sprintf("[%dx%d] %s\n\n", nrows, ncols, "DataFrame")

	// Add the row numbers
	for i := 0; i < nrows+1; i++ {
		if i != 0 {
			records[i] = append([]string{fmt.Sprintf("%v:", strconv.Itoa(i-1))}, records[i]...)
			continue
		}
		records[i] = append([]string{""}, records[i]...)
	}

	types := df.Types()
	for i := 0; i < ncols; i++ {
		typesrow[i] = fmt.Sprintf("<%v>", types[i])
	}
	typesrow = append([]string{""}, typesrow...)
	records = append(records, typesrow)
	recordsLen := len(records)
	for i := 0; i < recordsLen; i++ {
		for j := 0; j < ncols+1; j++ {
			// Escape special characters
			records[i][j] = strconv.Quote(records[i][j])
			records[i][j] = records[i][j][1 : len(records[i][j])-1]

			// Detect maximum number of characters per column
			if len(records[i][j]) > maxChars[j] {
				maxChars[j] = utf8.RuneCountInString(records[i][j])
			}
		}
	}
	maxCols := len(records[0])
	for i := 0; i < recordsLen; i++ {
		// Add right padding to all elements
		records[i][0] = addLeftPadding(records[i][0], maxChars[0]+1)
		for j := 1; j < ncols+1; j++ {
			records[i][j] = addRightPadding(records[i][j], maxChars[j])
		}
		records[i] = records[i][0:maxCols]
		// Create the final string

		dfStr += fmt.Sprintf("%v\n", strings.Join(records[i], " "))
	}

	fmt.Println(dfStr)
}
