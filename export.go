// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

// ExportScreenCSV generates a csv file of the Finviz screen results
func ExportScreenCSV(df *dataframe.DataFrame, outFileName string) error {
	f, err := os.Create(outFileName)
	if err != nil {
		return err
	}

	err = df.WriteCSV(f)
	if err != nil {
		return err
	}
	return nil
}

// ExportScreenJSON generates a json file of the Finviz screen results
func ExportScreenJSON(df *dataframe.DataFrame, outFileName string) error {
	f, err := os.Create(outFileName)
	if err != nil {
		return err
	}

	err = df.WriteJSON(f)
	if err != nil {
		return err
	}
	return nil
}

// PrintFullDataframe prints an entire dataframe to console
// Derived from https://github.com/go-gota/gota/blob/master/dataframe/dataframe.go until print method is made public
func PrintFullDataframe(df *dataframe.DataFrame) {
	class := "DataFrame"
	str := ""

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

	nrows, ncols := df.Dims()
	if df.Err != nil || nrows == 0 || ncols == 0 {
		return
	}
	var records [][]string
	records = df.Records()

	str += fmt.Sprintf("[%dx%d] %s\n\n", nrows, ncols, class)

	// Add the row numbers
	for i := 0; i < df.Nrow()+1; i++ {
		add := ""
		if i != 0 {
			add = strconv.Itoa(i-1) + ":"
		}
		records[i] = append([]string{add}, records[i]...)
	}
	types := df.Types()
	typesrow := make([]string, ncols)
	for i := 0; i < ncols; i++ {
		typesrow[i] = fmt.Sprintf("<%v>", types[i])
	}
	typesrow = append([]string{""}, typesrow...)
	records = append(records, typesrow)

	maxChars := make([]int, df.Ncol()+1)
	for i := 0; i < len(records); i++ {
		for j := 0; j < df.Ncol()+1; j++ {
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
	for i := 0; i < len(records); i++ {
		// Add right padding to all elements
		records[i][0] = addLeftPadding(records[i][0], maxChars[0]+1)
		for j := 1; j < df.Ncol()+1; j++ {
			records[i][j] = addRightPadding(records[i][j], maxChars[j])
		}
		records[i] = records[i][0:maxCols]
		// Create the final string
		str += strings.Join(records[i], " ")
		str += "\n"
	}
	fmt.Println(str)
}
