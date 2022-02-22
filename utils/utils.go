// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
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
		byteArray, err := io.ReadAll(html)
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
				if item == "-" {
					item = "NaN"
				}
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
	if df.Error() != nil || nrows == 0 || ncols == 0 {
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

// replaceCol modifies the typing of a DataFrame Series
func replaceCol(df *dataframe.DataFrame, columnName string, columnType series.Type, mapFunc func(series.Element)) {
	s := series.New(df.Col(columnName).Map(func(e series.Element) series.Element {
		result := e.Copy()

		mapFunc(result)

		return result
	}), columnType, columnName)

	*df = df.Mutate(s)
}

// CleanFinvizDataFrame attempts to improve the matching of a DataFrame's column type to the corresponding data
func CleanFinvizDataFrame(df *dataframe.DataFrame) *dataframe.DataFrame {
	columnNames := df.Names()
	columnCount := len(columnNames)

	for i := 0; i < columnCount; i++ {
		if val, exists := ColumnTypeLookup[strings.ToLower(columnNames[i])]; exists {
			switch val {
			default:
				continue
			case "percent":
				replaceCol(df, columnNames[i], series.Float, func(e series.Element) {
					num := strings.Split(e.String(), "%")[0]
					if percent, _ := strconv.ParseFloat(num, 64); percent != 0 {
						e.Set(percent / 100.0)
					} else {
						e.Set("NaN")
					}
				})

			case "float":
				replaceCol(df, columnNames[i], series.Float, func(e series.Element) {
					if num, _ := strconv.ParseFloat(e.String(), 64); num != 0 {
						e.Set(num)
					} else {
						e.Set("NaN")
					}
				})

			case "bigint":
				replaceCol(df, columnNames[i], series.Int, func(e series.Element) {
					stringResult := e.String()
					multiple := 1.0

					if strings.HasSuffix(stringResult, "B") {
						stringResult = stringResult[:len(stringResult)-1]
						multiple = 1000000000.0
					} else if strings.HasSuffix(stringResult, "M") {
						stringResult = stringResult[:len(stringResult)-1]
						multiple = 1000000.0
					} else if strings.HasSuffix(stringResult, "K") {
						stringResult = stringResult[:len(stringResult)-1]
						multiple = 1000.0
					}

					if num, _ := strconv.ParseFloat(stringResult, 64); num != 0 {
						e.Set(int(num * multiple))
					} else {
						e.Set("NaN")
					}
				})

			case "commaint":
				replaceCol(df, columnNames[i], series.Int, func(e series.Element) {
					stringResult := strings.Join(strings.Split(e.String(), ","), "")

					if num, _ := strconv.Atoi(stringResult); num != 0 {
						e.Set(num)
					} else {
						e.Set("NaN")
					}
				})

			case "int":
				replaceCol(df, columnNames[i], series.Int, func(e series.Element) {
					if num, _ := strconv.Atoi(e.String()); num != 0 {
						e.Set(num)
					} else {
						e.Set("NaN")
					}
				})

			case "string":
				replaceCol(df, columnNames[i], series.String, func(e series.Element) {
					if e.String() == "-" {
						e.Set("NaN")
					}
				})
			}
		}
	}

	return df
}

// ColumnTypeLookup specifies the type to parse results into
var ColumnTypeLookup = map[string]string{
	"no":                                "int",
	"no.":                               "int",
	"ticker":                            "string",
	"exchange":                          "string",
	"company":                           "string",
	"sector":                            "string",
	"industry":                          "string",
	"country":                           "string",
	"employees":                         "int",
	"index":                             "string",
	"income":                            "bigint",
	"sales":                             "bigint",
	"market cap":                        "bigint",
	"p/e":                               "float",
	"pe":                                "float",
	"forward p/e":                       "float",
	"fwd p/e":                           "float",
	"forward pe":                        "float",
	"peg":                               "float",
	"p/s":                               "float",
	"ps":                                "float",
	"p/b":                               "float",
	"pb":                                "float",
	"p/cash":                            "float",
	"p/c":                               "float",
	"pc":                                "float",
	"p/free cash flow":                  "float",
	"p/fcf":                             "float",
	"pfcf":                              "float",
	"book/sh":                           "float",
	"cash/sh":                           "float",
	"dividend yield":                    "percent",
	"dividend %":                        "percent",
	"dividend":                          "percent",
	"payout ratio":                      "percent",
	"payout":                            "percent",
	"eps":                               "float",
	"eps (ttm)":                         "float",
	"eps growth this year":              "percent",
	"eps this y":                        "percent",
	"eps growth next year":              "percent",
	"eps growth next y":                 "percent",
	"eps next y":                        "percent",
	"eps growth past 5 years":           "percent",
	"eps past 5y":                       "percent",
	"eps growth next 5 years":           "percent",
	"eps next 5y":                       "percent",
	"sales growth past 5 years":         "percent",
	"sales past 5y":                     "percent",
	"eps growth qtr over qtr":           "percent",
	"eps q/q":                           "percent",
	"eps growth quarter over quarter":   "percent",
	"sales growth qtr over qtr":         "percent",
	"sales q/q":                         "percent",
	"sales growth quarter over quarter": "percent",
	"shares outstanding":                "bigint",
	"shs outstand":                      "bigint",
	"outstanding":                       "bigint",
	"so":                                "bigint",
	"shares float":                      "bigint",
	"shs float":                         "bigint",
	"float":                             "bigint",
	"insider ownership":                 "percent",
	"insider own":                       "percent",
	"insider transactions":              "percent",
	"insider trans":                     "percent",
	"institutional ownership":           "percent",
	"inst own":                          "percent",
	"institutional transactions":        "percent",
	"inst trans":                        "percent",
	"float short":                       "percent",
	"short float":                       "percent",
	"short selling":                     "percent",
	"short ratio":                       "float",
	"return on assets":                  "percent",
	"roa":                               "percent",
	"return on equity":                  "percent",
	"roe":                               "percent",
	"return on investment":              "percent",
	"roi":                               "percent",
	"current ratio":                     "float",
	"curr r":                            "float",
	"quick ratio":                       "float",
	"quick r":                           "float",
	"long term debt/equity":             "float",
	"long-term debt/equity":             "float",
	"lt debt/eq":                        "float",
	"ltdebt/eq":                         "float",
	"lt d/e":                            "float",
	"lt de":                             "float",
	"total debt/equity":                 "float",
	"debt/equity":                       "float",
	"debt/eq":                           "float",
	"d/e":                               "float",
	"de":                                "float",
	"gross margin":                      "percent",
	"gross m":                           "percent",
	"gm":                                "percent",
	"operating margin":                  "percent",
	"oper. margin":                      "percent",
	"oper m":                            "percent",
	"om":                                "percent",
	"net profit margin":                 "percent",
	"profit margin":                     "percent",
	"profit m":                          "percent",
	"npm":                               "percent",
	"performance (week)":                "percent",
	"perf week":                         "percent",
	"performance (month)":               "percent",
	"perf month":                        "percent",
	"performance (quarter)":             "percent",
	"perf quart":                        "percent",
	"perf quarter":                      "percent",
	"performance (half year)":           "percent",
	"perf half":                         "percent",
	"perf half y":                       "percent",
	"performance (year)":                "percent",
	"perf year":                         "percent",
	"performance (yeartodate)":          "percent",
	"performance (ytd)":                 "percent",
	"perf ytd":                          "percent",
	"beta":                              "float",
	"average true range":                "float",
	"atr":                               "float",
	"volatility (week)":                 "percent",
	"volatility w":                      "percent",
	"volatility (month)":                "percent",
	"volatility m":                      "percent",
	"20-day simple moving average":      "percent",
	"sma20":                             "percent",
	"50-day simple moving average":      "percent",
	"sma50":                             "percent",
	"200-day simple moving average":     "percent",
	"sma200":                            "percent",
	"50-day high":                       "percent",
	"50d high":                          "percent",
	"50-day low":                        "percent",
	"50d low":                           "percent",
	"52w high":                          "percent",
	"52-week high":                      "percent",
	"52-week low":                       "percent",
	"52w low":                           "percent",
	"rsi":                               "float",
	"relative strength index":           "float",
	"change from open":                  "percent",
	"from open":                         "percent",
	"gap":                               "percent",
	"analyst recommendation":            "float",
	"analyst recom":                     "float",
	"recommendation":                    "float",
	"recom":                             "float",
	"average volume":                    "bigint",
	"avg volume":                        "bigint",
	"avgvol":                            "bigint",
	"relative volume":                   "float",
	"rel volume":                        "float",
	"relvol":                            "float",
	"prev close":                        "float",
	"price":                             "float",
	"change":                            "percent",
	"volume":                            "commaint",
	"earnings date":                     "string",
	"earnings":                          "string",
	"target price":                      "float",
	"ipo date":                          "string",
}
