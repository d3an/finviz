// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package screener

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"

	"github.com/d3an/finviz/utils"
)

func basicDefaultViewHelper(rootNode *goquery.Selection, headers []string, rawTickerData map[string]interface{}) (extraHeaders []string, extraRawTickerData map[string]interface{}) {
	rootNode.Find(".snapshot-table > tbody").Children().Each(func(j int, childNode *goquery.Selection) {
		if !utils.Contains(headers, childNode.Children().First().Text()) {
			headers = append(headers, childNode.Children().First().Text())
		}
		rawTickerData[childNode.Children().First().Text()] = childNode.Children().Last().Find("a").Text()
	})

	if !utils.Contains(headers, "Chart") {
		headers = append(headers, "Chart")
	}
	rawTickerData["Chart"] = rootNode.Find("img").AttrOr("src", "")

	rootNode.Find(".snapshot-table2 > tbody").Children().Each(func(j int, childNode *goquery.Selection) {
		if !utils.Contains(headers, childNode.Children().Eq(0).Text()) && !utils.Contains(headers, childNode.Children().Eq(2).Text()) {
			headers = append(append(headers, childNode.Children().Eq(0).Text()), childNode.Children().Eq(2).Text())
		}
		rawTickerData[childNode.Children().Eq(0).Text()] = childNode.Children().Eq(1).Text()
		rawTickerData[childNode.Children().Eq(2).Text()] = childNode.Children().Eq(3).Text()
	})

	return headers, rawTickerData
}

func basicNewsViewHelper(rootNode *goquery.Selection, headers []string, rawTickerData map[string]interface{}) (extraHeaders []string, extraRawTickerData map[string]interface{}) {
	if !utils.Contains(headers, "News") {
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
	if !utils.Contains(headers, "Description") {
		headers = append(headers, "Description")
	}
	description := rootNode.Find(".body-table-profile").Text()
	description = strings.TrimPrefix(description, "\"")
	description = strings.TrimSuffix(description, "\"")
	rawTickerData["Description"] = description
	return headers, rawTickerData
}

func basicInsiderTradingViewHelper(rootNode *goquery.Selection, headers []string, rawTickerData map[string]interface{}) (extraHeaders []string, extraRawTickerData map[string]interface{}) {
	if !utils.Contains(headers, "Insider Trading") {
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

func CleanScreenerDataFrame(df *dataframe.DataFrame) *dataframe.DataFrame {
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

func replaceCol(df *dataframe.DataFrame, columnName string, columnType series.Type, mapFunc func(series.Element)) {
	s := series.New(df.Col(columnName).Map(func(e series.Element) series.Element {
		result := e.Copy()

		mapFunc(result)

		return result
	}), columnType, columnName)

	*df = df.Mutate(s)
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
