// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"fmt"
	"strings"
)

// GeneralOrderType represents ascending or descending order
type GeneralOrderType string

// SpecificOrderType represents a more specific order, typically based on a filter
type SpecificOrderType string

// Order of Data
const (
	Ascending  GeneralOrderType = ""
	Descending GeneralOrderType = "-"
)

// Specific Order of Data
const (
	Signal                    SpecificOrderType = "sigrow"
	Ticker                    SpecificOrderType = "ticker"
	Company                   SpecificOrderType = "company"
	Sector                    SpecificOrderType = "sector"
	Industry                  SpecificOrderType = "industry"
	Country                   SpecificOrderType = "country"
	MarketCap                 SpecificOrderType = "marketcap"
	PE                        SpecificOrderType = "pe"
	ForwardPE                 SpecificOrderType = "forwardpe"
	PriceEarningsGrowth       SpecificOrderType = "peg"
	PriceSales                SpecificOrderType = "ps"
	PriceBook                 SpecificOrderType = "pb"
	PriceCash                 SpecificOrderType = "pc"
	PriceFCF                  SpecificOrderType = "pfcf"
	DividendYield             SpecificOrderType = "dividendyield"
	PayoutRatio               SpecificOrderType = "payoutratio"
	EPS                       SpecificOrderType = "eps"
	EPSGrowthThisYear         SpecificOrderType = "epsyoy"
	EPSGrowthNextYear         SpecificOrderType = "epsyoy1"
	EPSGrowthPast5Years       SpecificOrderType = "eps5years"
	EPSGrowthNext5Years       SpecificOrderType = "estltgrowth"
	SalesGrowthPast5Years     SpecificOrderType = "sales5years"
	EPSGrowthQtrOverQtr       SpecificOrderType = "epsqoq"
	SalesGrowthQtrOverQtr     SpecificOrderType = "salesqoq"
	SharesOutstanding         SpecificOrderType = "sharesoutstanding2"
	SharesFloat               SpecificOrderType = "sharesfloat"
	InsiderOwnership          SpecificOrderType = "insiderown"
	InsiderTransactions       SpecificOrderType = "insidertrans"
	InstitutionalOwnership    SpecificOrderType = "instown"
	InstitutionalTransactions SpecificOrderType = "insttrans"
	ShortInterestShare        SpecificOrderType = "shortinterestshare"
	ShortInterestRatio        SpecificOrderType = "shortinterestratio"
	EarningsDate              SpecificOrderType = "earningsdate"
	ROA                       SpecificOrderType = "roa"
	ROE                       SpecificOrderType = "roe"
	ROI                       SpecificOrderType = "roi"
	CurrentRatio              SpecificOrderType = "curratio"
	QuickRatio                SpecificOrderType = "quickratio"
	LTDebtEquity              SpecificOrderType = "ltdebteq"
	TotalDebtEquity           SpecificOrderType = "debteq"
	GrossMargin               SpecificOrderType = "grossmargin"
	OperatingMargin           SpecificOrderType = "opermargin"
	NetProfitMargin           SpecificOrderType = "netmargin"
	AnalystRecommendation     SpecificOrderType = "recom"
	PerformanceWeek           SpecificOrderType = "perf1w"
	PerformanceMonth          SpecificOrderType = "perf4w"
	PerformanceQuarter        SpecificOrderType = "perf13w"
	PerformanceHalfYear       SpecificOrderType = "perf26w"
	PerformanceYear           SpecificOrderType = "perf52w"
	PerformanceYTD            SpecificOrderType = "perfytd"
	Beta                      SpecificOrderType = "beta"
	AverageTrueRange          SpecificOrderType = "averagetruerange"
	VolatilityWeek            SpecificOrderType = "volatility1w"
	VolatilityMonth           SpecificOrderType = "volatility4w"
	RelativeSMA20Day          SpecificOrderType = "sma20"
	RelativeSMA50Day          SpecificOrderType = "sma50"
	RelativeSMA200Day         SpecificOrderType = "sma200"
	Relative50DayHigh         SpecificOrderType = "high50d"
	Relative50DayLow          SpecificOrderType = "low50d"
	Relative52WeekHigh        SpecificOrderType = "high52w"
	Relative52WeekLow         SpecificOrderType = "low52w"
	RSI14Day                  SpecificOrderType = "rsi"
	AverageVolume3Months      SpecificOrderType = "averagevolume"
	RelativeVolume            SpecificOrderType = "relativevolume"
	Change                    SpecificOrderType = "change"
	ChangeFromOpen            SpecificOrderType = "changeopen"
	Gap                       SpecificOrderType = "gap"
	Volume                    SpecificOrderType = "volume"
	Price                     SpecificOrderType = "price"
	TargetPrice               SpecificOrderType = "targetprice"
	IPODate                   SpecificOrderType = "ipodate"
)

// SpecificOrderLookup provides an interface for finding signals given string representations
var SpecificOrderLookup = map[string]SpecificOrderType{
	"signal":                    Signal,
	"ticker":                    Ticker,
	"company":                   Company,
	"sector":                    Sector,
	"industry":                  Industry,
	"country":                   Country,
	"marketcap":                 MarketCap,
	"pe":                        PE,
	"forwardpe":                 ForwardPE,
	"priceearningsgrowth":       PriceEarningsGrowth,
	"pricesales":                PriceSales,
	"ps":                        PriceSales,
	"pricebook":                 PriceBook,
	"pb":                        PriceBook,
	"pricecash":                 PriceCash,
	"pc":                        PriceCash,
	"pricefcf":                  PriceFCF,
	"pricefreecashflow":         PriceFCF,
	"dividendyield":             DividendYield,
	"payoutratio":               PayoutRatio,
	"eps":                       EPS,
	"epsgrowththisyear":         EPSGrowthThisYear,
	"epsgrowthnextyear":         EPSGrowthNextYear,
	"epsgrowthpast5years":       EPSGrowthPast5Years,
	"epsgrowthnext5years":       EPSGrowthNext5Years,
	"salesgrowthpast5years":     SalesGrowthPast5Years,
	"epsgrowthqtroverqtr":       EPSGrowthQtrOverQtr,
	"salesgrowthqtroverqtr":     SalesGrowthQtrOverQtr,
	"sharesoutstanding":         SharesOutstanding,
	"sharesfloat":               SharesFloat,
	"insiderownership":          InsiderOwnership,
	"insidertransactions":       InsiderTransactions,
	"institutionalownership":    InstitutionalOwnership,
	"institutionaltransactions": InstitutionalTransactions,
	"shortinterestshare":        ShortInterestShare,
	"shortinterestratio":        ShortInterestRatio,
	"earningsdate":              EarningsDate,
	"roa":                       ROA,
	"returnonassets":            ROA,
	"roe":                       ROE,
	"returnonequity":            ROE,
	"roi":                       ROI,
	"returnoninvestment":        ROI,
	"currentratio":              CurrentRatio,
	"quickratio":                QuickRatio,
	"ltdebtequity":              LTDebtEquity,
	"totaldebtequity":           TotalDebtEquity,
	"grossmargin":               GrossMargin,
	"operatingmargin":           OperatingMargin,
	"netprofitmargin":           NetProfitMargin,
	"analystrecommendation":     AnalystRecommendation,
	"performanceweek":           PerformanceWeek,
	"performancemonth":          PerformanceMonth,
	"performancequarter":        PerformanceQuarter,
	"performancehalfyear":       PerformanceHalfYear,
	"performanceyear":           PerformanceYear,
	"performanceytd":            PerformanceYTD,
	"beta":                      Beta,
	"averagetruerange":          AverageTrueRange,
	"volatilityweek":            VolatilityWeek,
	"volatilitymonth":           VolatilityMonth,
	"relativesma20day":          RelativeSMA20Day,
	"relativesma50day":          RelativeSMA50Day,
	"relativesma200day":         RelativeSMA200Day,
	"relative50dayhigh":         Relative50DayHigh,
	"relative50daylow":          Relative50DayLow,
	"relative52weekhigh":        Relative52WeekHigh,
	"relative52weeklow":         Relative52WeekLow,
	"rsi14day":                  RSI14Day,
	"rsi":                       RSI14Day,
	"averagevolume3months":      AverageVolume3Months,
	"relativevolume":            RelativeVolume,
	"change":                    Change,
	"changefromopen":            ChangeFromOpen,
	"gap":                       Gap,
	"volume":                    Volume,
	"price":                     Price,
	"targetprice":               TargetPrice,
	"ipodate":                   IPODate,
}

// GetSpecificOrder returns the constant of the queried signal
func GetSpecificOrder(query string) (SpecificOrderType, error) {
	if specificOrder, exists := SpecificOrderLookup[query]; exists {
		return specificOrder, nil
	}
	// Add logic that suggests similar matches
	return "", SpecificOrderNotFoundError(fmt.Sprintf("Specific order not found. Query: \"%v\"", query))
}

// GetGeneralOrder returns the constant of the queried signal
func GetGeneralOrder(query string) GeneralOrderType {
	if strings.HasPrefix(query, "-") {
		return Descending
	}
	return Ascending
}
