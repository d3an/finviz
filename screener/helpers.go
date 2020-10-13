// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package screener

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/d3an/finviz"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// ChartStyle defines the types of charts available in the FinViz screener
type ChartStyle = string

// ChartStyle constant definitions
const (
	Candle    ChartStyle = "candle"
	Line      ChartStyle = "line"
	Technical ChartStyle = "technical"
)

// TimeFrame defines the period length of charts available in the FinViz screener
type TimeFrame = string

// TimeFrame constant definitions
const (
	Min1    TimeFrame = "1min"
	Min5    TimeFrame = "5min"
	Min15   TimeFrame = "15min"
	Min30   TimeFrame = "30min"
	Daily   TimeFrame = "daily"
	Weekly  TimeFrame = "weekly"
	Monthly TimeFrame = "monthly"
)

func contains(slice, value interface{}) bool {
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

func basicDefaultViewHelper(rootNode *goquery.Selection, headers []string, rawTickerData map[string]interface{}) (extraHeaders []string, extraRawTickerData map[string]interface{}) {
	rootNode.Find(".snapshot-table > tbody").Children().Each(func(j int, childNode *goquery.Selection) {
		if !contains(headers, childNode.Children().First().Text()) {
			headers = append(headers, childNode.Children().First().Text())
		}
		rawTickerData[childNode.Children().First().Text()] = childNode.Children().Last().Find("a").Text()
	})

	if !contains(headers, "Chart") {
		headers = append(headers, "Chart")
	}
	rawTickerData["Chart"] = rootNode.Find("img").AttrOr("src", "")

	rootNode.Find(".snapshot-table2 > tbody").Children().Each(func(j int, childNode *goquery.Selection) {
		if !contains(headers, childNode.Children().Eq(0).Text()) && !contains(headers, childNode.Children().Eq(2).Text()) {
			headers = append(append(headers, childNode.Children().Eq(0).Text()), childNode.Children().Eq(2).Text())
		}
		rawTickerData[childNode.Children().Eq(0).Text()] = childNode.Children().Eq(1).Text()
		rawTickerData[childNode.Children().Eq(2).Text()] = childNode.Children().Eq(3).Text()
	})

	return headers, rawTickerData
}

func basicNewsViewHelper(rootNode *goquery.Selection, headers []string, rawTickerData map[string]interface{}) (extraHeaders []string, extraRawTickerData map[string]interface{}) {
	if !contains(headers, "News") {
		headers = append(headers, "News")
	}
	var news []map[string]string
	rootNode.Find(".body-table-news > tbody").Children().Each(func(j int, childNode *goquery.Selection) {
		news = append(news, map[string]string{
			"Datetime": childNode.Children().Eq(0).Text(),
			"Link":     childNode.Children().Eq(1).Find("a").AttrOr("href", ""),
			"Source":   childNode.Children().Eq(1).Find("span").Text(),
			"Title":    childNode.Children().Eq(1).Find("a").Text(),
		})
	})
	rawTickerData["News"] = news
	return headers, rawTickerData
}

func basicDescriptionViewHelper(rootNode *goquery.Selection, headers []string, rawTickerData map[string]interface{}) (extraHeaders []string, extraRawTickerData map[string]interface{}) {
	if !contains(headers, "Description") {
		headers = append(headers, "Description")
	}
	description := rootNode.Find(".body-table-profile").Text()
	description = strings.TrimPrefix(description, "\"")
	description = strings.TrimSuffix(description, "\"")
	rawTickerData["Description"] = description
	return headers, rawTickerData
}

func basicInsiderTradingViewHelper(rootNode *goquery.Selection, headers []string, rawTickerData map[string]interface{}) (extraHeaders []string, extraRawTickerData map[string]interface{}) {
	if !contains(headers, "Insider Trading") {
		headers = append(headers, "Insider Trading")
	}
	var insiderTrading []map[string]string
	if insiderNode := rootNode.Find(".body-table > tbody"); insiderNode != nil {
		insiderNode.Children().Each(func(k int, childNode *goquery.Selection) {
			if k > 0 {
				insiderTrading = append(insiderTrading, map[string]string{
					"Owner":             childNode.Children().Eq(0).Find("a").Text(),
					"Owner Link":        fmt.Sprintf("/%v", childNode.Children().Eq(0).Find("a").AttrOr("href", "")),
					"Relationship":      childNode.Children().Eq(1).Text(),
					"Date":              childNode.Children().Eq(2).Text(),
					"Transaction":       childNode.Children().Eq(3).Text(),
					"Cost":              childNode.Children().Eq(4).Text(),
					"# of Shares":       childNode.Children().Eq(5).Text(),
					"Value ($)":         childNode.Children().Eq(6).Text(),
					"# of Shares Total": childNode.Children().Eq(7).Text(),
					"SEC Form 4 Date":   childNode.Children().Eq(8).Find("a").Text(),
					"SEC Form 4 Link":   childNode.Children().Eq(8).Find("a").AttrOr("href", ""),
				})
			}
		})
		rawTickerData["Insider Trading"] = insiderTrading
	}
	return headers, rawTickerData
}

func generateRows(headers []string, tickerDataSlice []map[string]interface{}) (rows [][]string, err error) {
	headerCount := len(headers)
	resultCount := len(tickerDataSlice)

	rows = append(rows, headers)
	for i := 0; i < resultCount; i++ {
		var row []string

		for j := 0; j < headerCount; j++ {
			item := tickerDataSlice[i][headers[j]]
			switch item := item.(type) {
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

func getFilterURLComponent(filters interface{}) (string, error) {
	switch filters := filters.(type) {
	default:
		return "", fmt.Errorf("invalid type for list of filters")
	case []FilterInterface:
		filterSize := len(filters)
		if filterSize == 0 {
			return "", nil
		}

		var validFilters []FilterInterface
		var filterKeys []string

		for i := 0; i < filterSize; i++ {
			if filters[i].GetValue() == "" {
				return "", finviz.NoValuesError(fmt.Sprintf("%v filter was initialized without a value.", filters[i].GetName()))
			}
			if filterArrayContains(validFilters, filters[i]) {
				return "", finviz.DuplicateFilterError(fmt.Sprintf("%v filter was declared more than once.", filters[i].GetName()))
			}
			validFilters = append(validFilters, filters[i])
			filterKeys = append(filterKeys, filters[i].GetURLKey())
		}

		filterList := strings.Join(filterKeys, ",")
		return fmt.Sprintf("&f=%v", filterList), nil
	}
}

func getSignalURLComponent(signal interface{}) string {
	switch signal := signal.(type) {
	default:
		return ""
	case string:
		if signal == "" {
			return ""
		}
		return fmt.Sprintf("&s=%v", signal)
	}
}

func getSortOrderURLComponent(generalOrder, signal, specificOrder interface{}) string {
	// To sort by Signal, the Signal field must be non-empty
	if specificOrder == Signal && signal == "" && generalOrder == "" {
		return ""
	} else if specificOrder == Signal && signal == "" {
		return fmt.Sprintf("&o=%v", generalOrder)
	}

	if specificOrder == "" && generalOrder == "" {
		return ""
	}

	return fmt.Sprintf("&o=%v%v", generalOrder, specificOrder)
}

func getTickerURLComponent(tickers interface{}) string {
	switch tickers := tickers.(type) {
	default:
		return ""
	case []string:
		tickersSize := len(tickers)
		if tickersSize == 0 {
			return ""
		}

		for i := 0; i < tickersSize; i++ {
			tickers[i] = strings.ToUpper(tickers[i])
		}

		sort.Strings(tickers)
		tickerList := strings.Join(tickers, ",")
		return fmt.Sprintf("&t=%v", tickerList)
	}
}

func getSignalValue(viewArgs *map[string]interface{}) string {
	if value, exists := (*viewArgs)["signal"]; exists {
		return getSignalURLComponent(value)
	}
	return ""
}

func getCustomColumnsURLComponent(columns interface{}) (urlComponent string, err error) {
	switch columns := columns.(type) {
	default:
		return "", fmt.Errorf("invalid type for list of custom columns")
	case []string:
		columnsLen := len(columns)
		if columnsLen == 0 {
			return "", nil
		}

		var orderedColumns []string
		var bookKeeper [71]int // Update this to 74? for Elite
		for i := 0; i < columnsLen; i++ {
			if entry, err := strconv.Atoi(columns[i]); err == nil {
				if bookKeeper[entry] == 0 {
					bookKeeper[entry]++
					orderedColumns = append(orderedColumns, columns[i])
				}
			} else if val, ok := CustomColumnLookup[strings.ToLower(columns[i])]; ok {
				entry, _ := strconv.Atoi(val)
				if bookKeeper[entry] == 0 {
					bookKeeper[entry]++
					orderedColumns = append(orderedColumns, val)
				}
			} else {
				return "", fmt.Errorf("%v is not a valid custom column", columns[i])
			}
		}

		return fmt.Sprintf("&c=%v", strings.Join(orderedColumns, ",")), nil
	}
}

func getChartStylingURLComponent(cs, tf interface{}) (string, error) {
	var url string
	switch cs {
	default:
		return "", fmt.Errorf("invalid chart type")
	case Technical:
		url = ""
	case Line:
		url = "&ty=l&ta=0"
	case Candle:
		url = "&ta=0"
	}

	switch tf {
	default:
		return "", fmt.Errorf("invalid timeframe")
	case Daily:
		return url, nil
	case Weekly:
		if cs == Technical {
			return "", fmt.Errorf("incompatible chart type and time frame")
		}
		return fmt.Sprintf("%v&p=w", url), nil
	case Monthly:
		if cs == Technical {
			return "", fmt.Errorf("incompatible chart type and time frame")
		}
		return fmt.Sprintf("%v&p=m", url), nil
	case Min30:
		if cs == Line || cs == Technical {
			return "", fmt.Errorf("incompatible chart type and time frame")
		}
		return fmt.Sprintf("%v&p=i30", url), nil
	case Min15:
		if cs == Line || cs == Technical {
			return "", fmt.Errorf("incompatible chart type and time frame")
		}
		return fmt.Sprintf("%v&p=i15", url), nil
	case Min5:
		if cs == Technical {
			return "", fmt.Errorf("incompatible chart type and time frame")
		}
		return fmt.Sprintf("%v&p=i5", url), nil
	case Min1:
		if cs == Candle || cs == Technical {
			return "", fmt.Errorf("incompatible chart type and time frame")
		}
		return fmt.Sprintf("%v&p=i1", url), nil
	}
}

func getChartStylingValue(viewArgs *map[string]interface{}) (string, error) {
	csValue, csExists := (*viewArgs)["chart_type"]
	tfValue, tfExists := (*viewArgs)["timeframe"]

	var ct string
	if csExists {
		switch csValue := csValue.(type) {
		default:
			return "", fmt.Errorf("invalid chart type")
		case string:
			ct = csValue
		}
	}

	var tf string
	if tfExists {
		switch tfValue := tfValue.(type) {
		default:
			return "", fmt.Errorf("invalid chart type")
		case string:
			tf = tfValue
		}
	}

	if csExists && tfExists {
		result, err := getChartStylingURLComponent(ct, tf)
		if err != nil {
			return "", err
		}
		return result, nil
	} else if csExists {
		result, err := getChartStylingURLComponent(ct, Daily)
		if err != nil {
			return "", err
		}
		return result, nil
	} else if tfExists {
		result, err := getChartStylingURLComponent(Technical, tf)
		if err != nil {
			return "", err
		}
		return result, nil
	}
	return "", nil
}

func getCustomColumnsValue(viewArgs *map[string]interface{}) (string, error) {
	if value, exists := (*viewArgs)["custom_columns"]; exists {
		result, err := getCustomColumnsURLComponent(value)
		if err != nil {
			return "", err
		}
		return result, nil
	}
	return "", nil
}

func getFiltersValue(viewArgs *map[string]interface{}) (string, error) {
	if value, exists := (*viewArgs)["filters"]; exists {
		result, err := getFilterURLComponent(value)
		if err != nil {
			return "", err
		}
		return result, nil
	}
	return "", nil
}

func getTickersValue(viewArgs *map[string]interface{}) string {
	if value, exists := (*viewArgs)["tickers"]; exists {
		return getTickerURLComponent(value)
	}
	return ""
}

func getOrderValue(viewArgs *map[string]interface{}) string {
	generalOrderValue, generalOrderExists := (*viewArgs)["general_order"]
	if !generalOrderExists {
		generalOrderValue = ""
	}

	specificOrderValue, specificOrderExists := (*viewArgs)["specific_order"]
	if !specificOrderExists {
		specificOrderValue = ""
	}

	if generalOrderValue != "" || specificOrderValue != "" {
		signalValue, signalExists := (*viewArgs)["signal"]
		if !signalExists {
			signalValue = ""
		}
		return getSortOrderURLComponent(generalOrderValue, signalValue, specificOrderValue)
	}
	return ""
}

// GetScreenerData returns a DataFrame with screener data results
func GetScreenerData(c *http.Client, v finviz.ViewInterface, viewArgs *map[string]interface{}) (*dataframe.DataFrame, error) {
	url, err := v.GenerateURL(viewArgs)
	if err != nil {
		return nil, err
	}

	html, err := finviz.MakeGetRequest(c, url)
	if err != nil {
		return nil, err
	}

	doc, err := finviz.GenerateDocument(html)
	if err != nil {
		return nil, err
	}

	results, err := v.Scrape(doc)
	if err != nil {
		return nil, err
	}

	df := dataframe.LoadRecords(results)
	processedDf := CleanScreenerDataFrame(df)
	return processedDf, nil
}

func CleanScreenerDataFrame(df dataframe.DataFrame) *dataframe.DataFrame {
	columnNames := df.Names()
	columnCount := len(columnNames)

	for i := 0; i < columnCount; i++ {
		if val, exists := ColumnTypeLookup[strings.ToLower(columnNames[i])]; exists {
			switch val {
			default:
				continue
			case "percent":
				replaceCol(&df, columnNames[i], series.Float, func(e series.Element) {
					num := strings.Split(e.String(), "%")[0]
					if percent, _ := strconv.ParseFloat(num, 64); percent != 0 {
						e.Set(percent / 100.0)
					} else {
						e.Set("NaN")
					}
				})

			case "float":
				replaceCol(&df, columnNames[i], series.Float, func(e series.Element) {
					if num, _ := strconv.ParseFloat(e.String(), 64); num != 0 {
						e.Set(num)
					} else {
						e.Set("NaN")
					}
				})

			case "bigint":
				replaceCol(&df, columnNames[i], series.Int, func(e series.Element) {
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
				replaceCol(&df, columnNames[i], series.Int, func(e series.Element) {
					stringResult := strings.Join(strings.Split(e.String(), ","), "")

					if num, _ := strconv.Atoi(stringResult); num != 0 {
						e.Set(num)
					} else {
						e.Set("NaN")
					}
				})

			case "int":
				replaceCol(&df, columnNames[i], series.Int, func(e series.Element) {
					if num, _ := strconv.Atoi(e.String()); num != 0 {
						e.Set(num)
					} else {
						e.Set("NaN")
					}
				})

			case "string":
				replaceCol(&df, columnNames[i], series.String, func(e series.Element) {
					if e.String() == "-" {
						e.Set("NaN")
					}
				})
			}
		}
	}

	return &df
}

func replaceCol(df *dataframe.DataFrame, columnName string, columnType series.Type, mapFunc func(series.Element)) {
	s := series.New(df.Col(columnName).Map(func(e series.Element) series.Element {
		result := e.Copy()

		mapFunc(result)

		return result
	}), columnType, columnName)

	*df = df.Mutate(s)
}

// CustomColumnLookup provides an interface for selecting different columns for the CustomView
var CustomColumnLookup = map[string]string{
	"no":                                "0",
	"no.":                               "0",
	"ticker":                            "1",
	"company":                           "2",
	"sector":                            "3",
	"industry":                          "4",
	"country":                           "5",
	"market cap":                        "6",
	"p/e":                               "7",
	"pe":                                "7",
	"forward p/e":                       "8",
	"forward pe":                        "8",
	"fwd p/e":                           "8",
	"peg":                               "9",
	"p/s":                               "10",
	"ps":                                "10",
	"p/b":                               "11",
	"pb":                                "11",
	"p/cash":                            "12",
	"p/c":                               "12",
	"pc":                                "12",
	"p/free cash flow":                  "13",
	"p/fcf":                             "13",
	"pfcf":                              "13",
	"dividend yield":                    "14",
	"payout ratio":                      "15",
	"eps":                               "16",
	"eps growth this year":              "17",
	"eps growth next year":              "18",
	"eps growth past 5 years":           "19",
	"eps growth next 5 years":           "20",
	"sales growth past 5 years":         "21",
	"eps growth qtr over qtr":           "22",
	"eps growth quarter over quarter":   "22",
	"sales growth qtr over qtr":         "23",
	"sales growth quarter over quarter": "23",
	"shares outstanding":                "24",
	"so":                                "24",
	"shares float":                      "25",
	"float":                             "25",
	"insider ownership":                 "26",
	"insider transactions":              "27",
	"institutional ownership":           "28",
	"institutional transactions":        "29",
	"float short":                       "30",
	"short selling":                     "30",
	"short ratio":                       "31",
	"return on assets":                  "32",
	"roa":                               "32",
	"return on equity":                  "33",
	"roe":                               "33",
	"return on investment":              "34",
	"roi":                               "34",
	"current ratio":                     "35",
	"quick ratio":                       "36",
	"long term debt/equity":             "37",
	"long-term debt/equity":             "37",
	"lt d/e":                            "37",
	"lt de":                             "37",
	"total debt/equity":                 "38",
	"debt/equity":                       "38",
	"d/e":                               "38",
	"de":                                "38",
	"gross margin":                      "39",
	"gm":                                "39",
	"operating margin":                  "40",
	"om":                                "40",
	"net profit margin":                 "41",
	"npm":                               "41",
	"performance (week)":                "42",
	"performance (month)":               "43",
	"performance (quarter)":             "44",
	"performance (half year)":           "45",
	"performance (year)":                "46",
	"performance (yeartodate)":          "47",
	"performance (ytd)":                 "47",
	"beta":                              "48",
	"average true range":                "49",
	"atr":                               "49",
	"volatility (week)":                 "50",
	"volatility w":                      "50",
	"volatility (month)":                "51",
	"volatility m":                      "51",
	"20-day simple moving average":      "52",
	"sma20":                             "52",
	"50-day simple moving average":      "53",
	"sma50":                             "53",
	"200-day simple moving average":     "54",
	"sma200":                            "54",
	"50-day high":                       "55",
	"50-day low":                        "56",
	"52-week high":                      "57",
	"52-week low":                       "58",
	"rsi":                               "59",
	"relative strength index":           "59",
	"change from open":                  "60",
	"gap":                               "61",
	"analyst recommendation":            "62",
	"analyst recom":                     "62",
	"recommendation":                    "62",
	"recom":                             "62",
	"average volume":                    "63",
	"avgvol":                            "63",
	"relative volume":                   "64",
	"relvol":                            "64",
	"price":                             "65",
	"change":                            "66",
	"volume":                            "67",
	"earnings date":                     "68",
	"target price":                      "69",
	"ipo date":                          "70",
}

// ColumnTypeLookup specifies the type to parse results into
var ColumnTypeLookup = map[string]string{
	"no":                                "int",
	"no.":                               "int",
	"ticker":                            "string",
	"company":                           "string",
	"sector":                            "string",
	"industry":                          "string",
	"country":                           "string",
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
	"dividend yield":                    "percent",
	"dividend":                          "percent",
	"payout ratio":                      "percent",
	"eps":                               "float",
	"eps growth this year":              "percent",
	"eps this y":                        "percent",
	"eps growth next year":              "percent",
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
	"outstanding":                       "bigint",
	"so":                                "bigint",
	"shares float":                      "bigint",
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
	"oper m":                            "percent",
	"om":                                "percent",
	"net profit margin":                 "percent",
	"profit m":                          "percent",
	"npm":                               "percent",
	"performance (week)":                "percent",
	"perf week":                         "percent",
	"performance (month)":               "percent",
	"perf month":                        "percent",
	"performance (quarter)":             "percent",
	"perf quart":                        "percent",
	"performance (half year)":           "percent",
	"perf half":                         "percent",
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
	"price":                             "float",
	"change":                            "percent",
	"volume":                            "commaint",
	"earnings date":                     "string",
	"earnings":                          "string",
	"target price":                      "float",
	"ipo date":                          "string",
}
