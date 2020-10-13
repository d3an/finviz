// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package screener

import (
	"fmt"
	"github.com/d3an/finviz"
	"os"
	"sort"
	"strings"
)

// GetName returns the name of a given filter
func (a *Filter) GetName() string {
	return a.Name
}

// GetProperties returns a dict of a filter's properties
func (a *Filter) GetProperties() map[string]bool {
	return a.Properties
}

// GetProperty returns the value of a specific filter property
func (a *Filter) GetProperty(key string) bool {
	return a.Properties[key]
}

// GetURLKey returns the filter's url key
func (a *Filter) GetURLKey() string {
	return fmt.Sprintf("%v_%v", a.URLPrefix, a.Value)
}

// GetValue returns the filter's value(s)
func (a *Filter) GetValue() string {
	return a.Value
}

// SetValues constructs the value(s) for a filter type.
func (a *Filter) SetValues(values []string) *Filter {
	numValues := len(values)

	if numValues == 0 {
		fmt.Println(finviz.NoValuesError(
			fmt.Sprintf("%v was initialized without a value.", a.GetName())))
		os.Exit(1)
	} else if numValues > 1 && !a.GetProperty("multipleValues") {
		fmt.Println(finviz.MultipleValuesError(
			fmt.Sprintf("%v filter does not support multiple values. Values passed: [%v]", a.GetName(), strings.Join(values, ","))))
		os.Exit(1)
	}

	sort.Strings(values)
	a.Value = strings.Join(values, "|")
	return a
}

// GetFilter attempts to create a custom filter dependent on the FilterLookup
func GetFilter(filterQuery string, filterValues ...string) (*Filter, error) {
	if filterConstructor, exists := FilterLookup[filterQuery]; exists {
		filter := filterConstructor(filterValues...)

		valueCount := len(filterValues)
		var parsedFilterValues []string
		for i := 0; i < valueCount; i++ {
			if value, exists := FilterValueLookup[strings.ToLower(filter.GetName())][strings.ToLower(filterValues[i])]; exists {
				parsedFilterValues = append(parsedFilterValues, value)
			} else {
				parsedFilterValues = append(parsedFilterValues, filterValues[i])
			}
		}
		filter = filterConstructor(parsedFilterValues...)

		return filter, nil
	}
	// Add logic that suggests similar matches
	return nil, finviz.FilterNotFoundError(fmt.Sprintf("Filter not found. Query: \"%v\"", filterQuery))
}

// filterArrayContains returns true if a filter exists in an array of filters and false otherwise
func filterArrayContains(fs []FilterInterface, f FilterInterface) bool {
	filterCount := len(fs)
	for i := 0; i < filterCount; i++ {
		if fs[i].GetName() == f.GetName() {
			return true
		}
	}
	return false
}

/***************************************************************************
*****	DESCRIPTIVE   ******************************************************
***************************************************************************/

// ExchangeFilter returns a reference to a Filter type
func ExchangeFilter(exchangeTypes ...ExchangeType) *Filter {
	filter := Filter{
		Name: "Exchange",
		Properties: map[string]bool{
			"multipleValues": true,
			"customRange":    false,
		},
		URLPrefix: "exch",
	}
	return filter.SetValues(exchangeTypes)
}

// IndexFilter returns a reference to a Filter type
func IndexFilter(indexTypes ...IndexType) *Filter {
	filter := Filter{
		Name: "Index",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "idx",
	}
	return filter.SetValues(indexTypes)
}

// SectorFilter returns a reference to a Filter type
func SectorFilter(sectorTypes ...SectorType) *Filter {
	filter := Filter{
		Name: "Sector",
		Properties: map[string]bool{
			"multipleValues": true,
			"customRange":    false,
		},
		URLPrefix: "sec",
	}
	return filter.SetValues(sectorTypes)
}

// IndustryFilter returns a reference to a Filter type
func IndustryFilter(industryTypes ...IndustryType) *Filter {
	filter := Filter{
		Name: "Industry",
		Properties: map[string]bool{
			"multipleValues": true,
			"customRange":    false,
		},
		URLPrefix: "ind",
	}
	return filter.SetValues(industryTypes)
}

// CountryFilter returns a reference to a Filter type
func CountryFilter(countryTypes ...CountryType) *Filter {
	filter := Filter{
		Name: "Country",
		Properties: map[string]bool{
			"multipleValues": true,
			"customRange":    false,
		},
		URLPrefix: "geo",
	}
	return filter.SetValues(countryTypes)
}

// MarketCapFilter returns a reference to a Filter type
func MarketCapFilter(marketCapTypes ...MarketCapType) *Filter {
	filter := Filter{
		Name: "Market Cap",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "cap",
	}
	return filter.SetValues(marketCapTypes)
}

// DividendYieldFilter returns a reference to a Filter type
func DividendYieldFilter(dividendYieldTypes ...DividendYieldType) *Filter {
	filter := Filter{
		Name: "Dividend Yield",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_div",
	}
	return filter.SetValues(dividendYieldTypes)
}

// ShortSellingFilter returns a reference to a Filter type
func ShortSellingFilter(shortSellingTypes ...ShortSellingType) *Filter {
	filter := Filter{
		Name: "Short Selling",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "sh_short",
	}
	return filter.SetValues(shortSellingTypes)
}

// RecommendationFilter returns a reference to a Filter type
func RecommendationFilter(recommendationTypes ...RecommendationType) *Filter {
	filter := Filter{
		Name: "Analyst Recommendation",
		Properties: map[string]bool{
			"multipleValues": true,
			"customRange":    false,
		},
		URLPrefix: "an_recom",
	}
	return filter.SetValues(recommendationTypes)
}

// OptionShortFilter returns a reference to a Filter type
func OptionShortFilter(optionShortTypes ...OptionShortType) *Filter {
	filter := Filter{
		Name: "Option Short",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "sh_opt",
	}
	return filter.SetValues(optionShortTypes)
}

// EarningsDateFilter returns a reference to a Filter type
func EarningsDateFilter(earningsDateTypes ...EarningsDateType) *Filter {
	filter := Filter{
		Name: "Earnings Date",
		Properties: map[string]bool{
			"multipleValues": true,
			"customRange":    false,
		},
		URLPrefix: "earningsdate",
	}
	return filter.SetValues(earningsDateTypes)
}

// AverageVolumeFilter returns a reference to a Filter type
func AverageVolumeFilter(averageVolumeTypes ...AverageVolumeType) *Filter {
	filter := Filter{
		Name: "Average Volume",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "sh_avgvol",
	}
	return filter.SetValues(averageVolumeTypes)
}

// RelativeVolumeFilter returns a reference to a Filter type
func RelativeVolumeFilter(relativeVolumeTypes ...RelativeVolumeType) *Filter {
	filter := Filter{
		Name: "Relative Volume",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "sh_relvol",
	}
	return filter.SetValues(relativeVolumeTypes)
}

// CurrentVolumeFilter returns a reference to a Filter type
func CurrentVolumeFilter(currentVolumeTypes ...CurrentVolumeType) *Filter {
	filter := Filter{
		Name: "Current Volume",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "sh_curvol",
	}
	return filter.SetValues(currentVolumeTypes)
}

// PriceFilter returns a reference to a Filter type
func PriceFilter(priceTypes ...PriceType) *Filter {
	filter := Filter{
		Name: "Price",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "sh_price",
	}
	return filter.SetValues(priceTypes)
}

// TargetPriceFilter returns a reference to a Filter type
func TargetPriceFilter(targetPriceTypes ...TargetPriceType) *Filter {
	filter := Filter{
		Name: "Target Price",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "targetprice",
	}
	return filter.SetValues(targetPriceTypes)
}

// IPODateFilter returns a reference to a Filter type
func IPODateFilter(ipoDateTypes ...IPODateType) *Filter {
	filter := Filter{
		Name: "IPO Date",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "ipodate",
	}
	return filter.SetValues(ipoDateTypes)
}

// SharesOutstandingFilter returns a reference to a Filter type
func SharesOutstandingFilter(sharesOutstandingTypes ...SharesOutstandingType) *Filter {
	filter := Filter{
		Name: "Shares Outstanding",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "sh_outstanding",
	}
	return filter.SetValues(sharesOutstandingTypes)
}

// FloatFilter returns a reference to a Filter type
func FloatFilter(floatTypes ...SharesOutstandingType) *Filter {
	filter := Filter{
		Name: "Float",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "sh_float",
	}
	return filter.SetValues(floatTypes)
}

/***************************************************************************
*****	FUNDAMENTALS   *****************************************************
***************************************************************************/

// PEFilter returns a reference to a Filter type
func PEFilter(peTypes ...PEType) *Filter {
	filter := Filter{
		Name: "P/E",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_pe",
	}
	return filter.SetValues(peTypes)
}

// ForwardPEFilter returns a reference to a Filter type
func ForwardPEFilter(forwardPETypes ...PEType) *Filter {
	filter := Filter{
		Name: "Forward P/E",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_fpe",
	}
	return filter.SetValues(forwardPETypes)
}

// PEGFilter returns a reference to a Filter type
func PEGFilter(pegTypes ...PEGType) *Filter {
	filter := Filter{
		Name: "PEG",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_peg",
	}
	return filter.SetValues(pegTypes)
}

// PriceSalesFilter returns a reference to a Filter type
func PriceSalesFilter(priceSalesTypes ...PriceSalesType) *Filter {
	filter := Filter{
		Name: "P/S",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_ps",
	}
	return filter.SetValues(priceSalesTypes)
}

// PriceBookFilter returns a reference to a Filter type
func PriceBookFilter(priceBookTypes ...PriceBookType) *Filter {
	filter := Filter{
		Name: "P/B",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_pb",
	}
	return filter.SetValues(priceBookTypes)
}

// PriceCashFilter returns a reference to a Filter type
func PriceCashFilter(priceCashTypes ...PriceCashType) *Filter {
	filter := Filter{
		Name: "Price/Cash",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_pc",
	}
	return filter.SetValues(priceCashTypes)
}

// PriceFCFFilter returns a reference to a Filter type
func PriceFCFFilter(priceFCFTypes ...PriceFCFType) *Filter {
	filter := Filter{
		Name: "Price/Free Cash Flow",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_pfcf",
	}
	return filter.SetValues(priceFCFTypes)
}

// EPSGrowthThisYearFilter returns a reference to a Filter type
func EPSGrowthThisYearFilter(epsGrowthThisYearTypes ...GrowthType) *Filter {
	filter := Filter{
		Name: "EPS Growth This Year",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_epsyoy",
	}
	return filter.SetValues(epsGrowthThisYearTypes)
}

// EPSGrowthNextYearFilter returns a reference to a Filter type
func EPSGrowthNextYearFilter(epsGrowthNextYearTypes ...GrowthType) *Filter {
	filter := Filter{
		Name: "EPS Growth Next Year",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_epsyoy1",
	}
	return filter.SetValues(epsGrowthNextYearTypes)
}

// EPSGrowthPast5YearsFilter returns a reference to a Filter type
func EPSGrowthPast5YearsFilter(epsGrowthPast5YearsTypes ...GrowthType) *Filter {
	filter := Filter{
		Name: "EPS Growth Past 5 Years",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_eps5years",
	}
	return filter.SetValues(epsGrowthPast5YearsTypes)
}

// EPSGrowthNext5YearsFilter returns a reference to a Filter type
func EPSGrowthNext5YearsFilter(epsGrowthNext5YearsTypes ...GrowthType) *Filter {
	filter := Filter{
		Name: "EPS Growth Next 5 Years",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_estltgrowth",
	}
	return filter.SetValues(epsGrowthNext5YearsTypes)
}

// SalesGrowthPast5YearsFilter returns a reference to a Filter type
func SalesGrowthPast5YearsFilter(salesGrowthPast5YearsTypes ...GrowthType) *Filter {
	filter := Filter{
		Name: "Sales Growth Past 5 Years",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_sales5years",
	}
	return filter.SetValues(salesGrowthPast5YearsTypes)
}

// EPSGrowthQtrOverQtrFilter returns a reference to a Filter type
func EPSGrowthQtrOverQtrFilter(epsGrowthQtrOverQtrTypes ...GrowthType) *Filter {
	filter := Filter{
		Name: "EPS Growth Quarter Over Quarter",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_epsqoq",
	}
	return filter.SetValues(epsGrowthQtrOverQtrTypes)
}

// SalesGrowthQtrOverQtrFilter returns a reference to a Filter type
func SalesGrowthQtrOverQtrFilter(salesGrowthQtrOverQtrTypes ...GrowthType) *Filter {
	filter := Filter{
		Name: "Sales Growth Quarter over Quarter",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_salesqoq",
	}
	return filter.SetValues(salesGrowthQtrOverQtrTypes)
}

// ROAFilter returns a reference to a Filter type
func ROAFilter(roaTypes ...ReturnType) *Filter {
	filter := Filter{
		Name: "ROA",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_roa",
	}
	return filter.SetValues(roaTypes)
}

// ROEFilter returns a reference to a Filter type
func ROEFilter(roeTypes ...ReturnType) *Filter {
	filter := Filter{
		Name: "ROE",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_roe",
	}
	return filter.SetValues(roeTypes)
}

// ROIFilter returns a reference to a Filter type
func ROIFilter(roiTypes ...ReturnType) *Filter {
	filter := Filter{
		Name: "ROI",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_roi",
	}
	return filter.SetValues(roiTypes)
}

// CurrentRatioFilter returns a reference to a Filter type
func CurrentRatioFilter(currentRatioTypes ...AssetRatioType) *Filter {
	filter := Filter{
		Name: "Current Ratio",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_curratio",
	}
	return filter.SetValues(currentRatioTypes)
}

// QuickRatioFilter returns a reference to a Filter type
func QuickRatioFilter(quickRatioTypes ...AssetRatioType) *Filter {
	filter := Filter{
		Name: "Quick Ratio",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_quickratio",
	}
	return filter.SetValues(quickRatioTypes)
}

// LTDebtEquityFilter returns a reference to a Filter type
func LTDebtEquityFilter(ltDebtEquityTypes ...DebtEquityType) *Filter {
	filter := Filter{
		Name: "Long-Term Debt/Equity",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_ltdebteq",
	}
	return filter.SetValues(ltDebtEquityTypes)
}

// DebtEquityFilter returns a reference to a Filter type
func DebtEquityFilter(debtEquityTypes ...DebtEquityType) *Filter {
	filter := Filter{
		Name: "Debt/Equity",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_debteq",
	}
	return filter.SetValues(debtEquityTypes)
}

// GrossMarginFilter returns a reference to a Filter type
func GrossMarginFilter(grossMarginTypes ...GrossMarginType) *Filter {
	filter := Filter{
		Name: "Gross Margin",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_grossmargin",
	}
	return filter.SetValues(grossMarginTypes)
}

// OperatingMarginFilter returns a reference to a Filter type
func OperatingMarginFilter(operatingMarginTypes ...OperatingMarginType) *Filter {
	filter := Filter{
		Name: "Operating Margin",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_opermargin",
	}
	return filter.SetValues(operatingMarginTypes)
}

// NetProfitMarginFilter returns a reference to a Filter type
func NetProfitMarginFilter(netProfitMarginTypes ...NetProfitMarginType) *Filter {
	filter := Filter{
		Name: "Net Profit Margin",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_netmargin",
	}
	return filter.SetValues(netProfitMarginTypes)
}

// PayoutRatioFilter returns a reference to a Filter type
func PayoutRatioFilter(payoutRatioTypes ...PayoutRatioType) *Filter {
	filter := Filter{
		Name: "Payout Ratio",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "fa_payoutratio",
	}
	return filter.SetValues(payoutRatioTypes)
}

// InsiderOwnershipFilter returns a reference to a Filter type
func InsiderOwnershipFilter(insiderOwnershipTypes ...InsiderOwnershipType) *Filter {
	filter := Filter{
		Name: "Insider Ownership",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "sh_insiderown",
	}
	return filter.SetValues(insiderOwnershipTypes)
}

// InsiderTransactionsFilter returns a reference to a Filter type
func InsiderTransactionsFilter(insiderTransactionsTypes ...InsiderTransactionsType) *Filter {
	filter := Filter{
		Name: "Insider Transactions",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "sh_insidertrans",
	}
	return filter.SetValues(insiderTransactionsTypes)
}

// InstitutionalOwnershipFilter returns a reference to a Filter type
func InstitutionalOwnershipFilter(institutionalOwnershipTypes ...InstitutionalOwnershipType) *Filter {
	filter := Filter{
		Name: "Institutional Ownership",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "sh_instown",
	}
	return filter.SetValues(institutionalOwnershipTypes)
}

// InstitutionalTransactionsFilter returns a reference to a Filter type
func InstitutionalTransactionsFilter(institutionalTransactionsTypes ...InstitutionalTransactionsType) *Filter {
	filter := Filter{
		Name: "Institutional Transactions",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "sh_insttrans",
	}
	return filter.SetValues(institutionalTransactionsTypes)
}

/***************************************************************************
*****	TECHNICAL   ********************************************************
***************************************************************************/

// PerformanceFilter returns a reference to a Filter type
func PerformanceFilter(performanceTypes ...PerformanceType) *Filter {
	filter := Filter{
		Name: "Performance",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "ta_perf",
	}
	return filter.SetValues(performanceTypes)
}

// Performance2Filter returns a reference to a Filter type
func Performance2Filter(performance2Types ...PerformanceType) *Filter {
	filter := Filter{
		Name: "Performance2",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "ta_perf2",
	}
	return filter.SetValues(performance2Types)
}

// VolatilityFilter returns a reference to a Filter type
func VolatilityFilter(volatilityTypes ...VolatilityType) *Filter {
	filter := Filter{
		Name: "Volatility",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "ta_volatility",
	}
	return filter.SetValues(volatilityTypes)
}

// RSIFilter returns a reference to a Filter type
func RSIFilter(rsiTypes ...RSIType) *Filter {
	filter := Filter{
		Name: "RSI",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "ta_rsi",
	}
	return filter.SetValues(rsiTypes)
}

// GapFilter returns a reference to a Filter type
func GapFilter(gapTypes ...GapType) *Filter {
	filter := Filter{
		Name: "Gap",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "ta_gap",
	}
	return filter.SetValues(gapTypes)
}

// SMA20Filter returns a reference to a Filter type
func SMA20Filter(sma20Types ...SMA20Type) *Filter {
	filter := Filter{
		Name: "SMA 20",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "ta_sma20",
	}
	return filter.SetValues(sma20Types)
}

// SMA50Filter returns a reference to a Filter type
func SMA50Filter(sma50Types ...SMA50Type) *Filter {
	filter := Filter{
		Name: "SMA 50",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "ta_sma50",
	}
	return filter.SetValues(sma50Types)
}

// SMA200Filter returns a reference to a Filter type
func SMA200Filter(sma200Types ...SMA200Type) *Filter {
	filter := Filter{
		Name: "SMA 200",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "ta_sma200",
	}
	return filter.SetValues(sma200Types)
}

// ChangeFilter returns a reference to a Filter type
func ChangeFilter(changeTypes ...ChangeType) *Filter {
	filter := Filter{
		Name: "Change",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "ta_change",
	}
	return filter.SetValues(changeTypes)
}

// ChangeFromOpenFilter returns a reference to a Filter type
func ChangeFromOpenFilter(changeFromOpenTypes ...ChangeType) *Filter {
	filter := Filter{
		Name: "Change From Open",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "ta_changeopen",
	}
	return filter.SetValues(changeFromOpenTypes)
}

// HighLow20DayFilter returns a reference to a Filter type
func HighLow20DayFilter(highLow20DayTypes ...HighLowDayType) *Filter {
	filter := Filter{
		Name: "20 Day High/Low",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "ta_highlow20d",
	}
	return filter.SetValues(highLow20DayTypes)
}

// HighLow50DayFilter returns a reference to a Filter type
func HighLow50DayFilter(highLow50DayTypes ...HighLowDayType) *Filter {
	filter := Filter{
		Name: "50 Day High/Low",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "ta_highlow50d",
	}
	return filter.SetValues(highLow50DayTypes)
}

// HighLow52WeekFilter returns a reference to a Filter type
func HighLow52WeekFilter(highLow52WeekTypes ...HighLow52WeekType) *Filter {
	filter := Filter{
		Name: "52 Week High/Low",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    false,
		},
		URLPrefix: "ta_highlow52w",
	}
	return filter.SetValues(highLow52WeekTypes)
}

// PatternFilter returns a reference to a Filter type
func PatternFilter(patternTypes ...PatternType) *Filter {
	filter := Filter{
		Name: "Pattern",
		Properties: map[string]bool{
			"multipleValues": true,
			"customRange":    false,
		},
		URLPrefix: "ta_pattern",
	}
	return filter.SetValues(patternTypes)
}

// CandlestickFilter returns a reference to a Filter type
func CandlestickFilter(candlestickTypes ...CandlestickType) *Filter {
	filter := Filter{
		Name: "Candlestick",
		Properties: map[string]bool{
			"multipleValues": true,
			"customRange":    false,
		},
		URLPrefix: "ta_candlestick",
	}
	return filter.SetValues(candlestickTypes)
}

// BetaFilter returns a reference to a Filter type
func BetaFilter(betaTypes ...BetaType) *Filter {
	filter := Filter{
		Name: "Beta",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "ta_beta",
	}
	return filter.SetValues(betaTypes)
}

// AverageTrueRangeFilter returns a reference to a Filter type
func AverageTrueRangeFilter(averageTrueRangeTypes ...AverageTrueRangeType) *Filter {
	filter := Filter{
		Name: "Average True Range",
		Properties: map[string]bool{
			"multipleValues": false,
			"customRange":    true,
		},
		URLPrefix: "ta_averagetruerange",
	}
	return filter.SetValues(averageTrueRangeTypes)
}

// CustomFilter returns a reference to a filter based on custom set values
func CustomFilter(name, urlPrefix string, supportsMultipleValues, supportsCustomRange bool, values ...string) FilterInterface {
	filter := Filter{
		Name: name,
		Properties: map[string]bool{
			"multipleValues": supportsMultipleValues,
			"customRange":    supportsCustomRange,
		},
		URLPrefix: urlPrefix,
	}
	return filter.SetValues(values)
}

// FilterLookup is a map of common query strings to default filter constructors
var FilterLookup = map[string]func(...string) *Filter{
	"exchange":               ExchangeFilter,
	"index":                  IndexFilter,
	"sector":                 SectorFilter,
	"industry":               IndustryFilter,
	"country":                CountryFilter,
	"marketcap":              MarketCapFilter,
	"market cap":             MarketCapFilter,
	"dividendyield":          DividendYieldFilter,
	"dividend yield":         DividendYieldFilter,
	"divyield":               DividendYieldFilter,
	"floatshort":             ShortSellingFilter,
	"float short":            ShortSellingFilter,
	"shortselling":           ShortSellingFilter,
	"short selling":          ShortSellingFilter,
	"analystrecommendation":  RecommendationFilter,
	"analyst recommendation": RecommendationFilter,
	"analystrecom":           RecommendationFilter,
	"analyst recom":          RecommendationFilter,
	"recommendation":         RecommendationFilter,
	"recom":                  RecommendationFilter,
	"optionshort":            OptionShortFilter,
	"option short":           OptionShortFilter,
	"earningsdate":           EarningsDateFilter,
	"earnings date":          EarningsDateFilter,
	"averagevolume":          AverageVolumeFilter,
	"average volume":         AverageVolumeFilter,
	"avgvolume":              AverageVolumeFilter,
	"avg volume":             AverageVolumeFilter,
	"avgvol":                 AverageVolumeFilter,
	"avg vol":                AverageVolumeFilter,
	"relativevolume":         RelativeVolumeFilter,
	"relative volume":        RelativeVolumeFilter,
	"relvolume":              RelativeVolumeFilter,
	"rel volume":             RelativeVolumeFilter,
	"relvol":                 RelativeVolumeFilter,
	"rel vol":                RelativeVolumeFilter,
	"currentvolume":          CurrentVolumeFilter,
	"curvolume":              CurrentVolumeFilter,
	"curvol":                 CurrentVolumeFilter,
	"price":                  PriceFilter,
	"targetprice":            TargetPriceFilter,
	"ipodate":                IPODateFilter,
	"ipo":                    IPODateFilter,
	"sharesoutstanding":      SharesOutstandingFilter,
	"so":                     SharesOutstandingFilter,
	"floatshares":            FloatFilter,
	"float":                  FloatFilter,

	"pe":                            PEFilter,
	"p/e":                           PEFilter,
	"forwardpe":                     ForwardPEFilter,
	"forward pe":                    ForwardPEFilter,
	"forward p/e":                   ForwardPEFilter,
	"peg":                           PEGFilter,
	"pricesales":                    PriceSalesFilter,
	"price/sales":                   PriceSalesFilter,
	"p/s":                           PriceSalesFilter,
	"ps":                            PriceSalesFilter,
	"pricebook":                     PriceBookFilter,
	"pb":                            PriceBookFilter,
	"pricecash":                     PriceCashFilter,
	"pc":                            PriceCashFilter,
	"pricefreecashflow":             PriceFCFFilter,
	"pricefcf":                      PriceFCFFilter,
	"pfcf":                          PriceFCFFilter,
	"epsgrowththisyear":             EPSGrowthThisYearFilter,
	"epsgrowthnextyear":             EPSGrowthNextYearFilter,
	"epsgrowthpast5years":           EPSGrowthPast5YearsFilter,
	"epsgrowthnext5years":           EPSGrowthNext5YearsFilter,
	"salesgrowthpast5years":         SalesGrowthPast5YearsFilter,
	"epsgrowthqtroverqtr":           EPSGrowthQtrOverQtrFilter,
	"epsgrowthquarteroverquarter":   EPSGrowthQtrOverQtrFilter,
	"epsgrowthqoq":                  EPSGrowthQtrOverQtrFilter,
	"salesgrowthqtroverqtr":         SalesGrowthQtrOverQtrFilter,
	"salesgrowthquarteroverquarter": SalesGrowthQtrOverQtrFilter,
	"salesgrowthqoq":                SalesGrowthQtrOverQtrFilter,
	"roa":                           ROAFilter,
	"returnonassets":                ROAFilter,
	"roe":                           ROEFilter,
	"returnonequity":                ROEFilter,
	"roi":                           ROIFilter,
	"returnoninvestment":            ROIFilter,
	"currentratio":                  CurrentRatioFilter,
	"curratio":                      CurrentRatioFilter,
	"quickratio":                    QuickRatioFilter,
	"ltdebtequity":                  LTDebtEquityFilter,
	"longtermdebtequity":            LTDebtEquityFilter,
	"debtequity":                    DebtEquityFilter,
	"grossmargin":                   GrossMarginFilter,
	"gm":                            GrossMarginFilter,
	"operatingmargin":               OperatingMarginFilter,
	"om":                            OperatingMarginFilter,
	"netprofitmargin":               NetProfitMarginFilter,
	"npm":                           NetProfitMarginFilter,
	"payoutratio":                   PayoutRatioFilter,
	"insiderownership":              InsiderOwnershipFilter,
	"insider ownership":             InsiderOwnershipFilter,
	"insdrown":                      InsiderOwnershipFilter,
	"insidertransactions":           InsiderTransactionsFilter,
	"insider transactions":          InsiderTransactionsFilter,
	"insdrtrans":                    InsiderTransactionsFilter,
	"institutionalownership":        InstitutionalOwnershipFilter,
	"institutional ownership":       InstitutionalOwnershipFilter,
	"instown":                       InstitutionalOwnershipFilter,
	"institutionaltransactions":     InsiderTransactionsFilter,
	"institutional transactions":    InsiderTransactionsFilter,
	"insttrans":                     InstitutionalTransactionsFilter,

	"performance":      PerformanceFilter,
	"perf":             PerformanceFilter,
	"performance2":     Performance2Filter,
	"perf2":            Performance2Filter,
	"volatility":       VolatilityFilter,
	"rsi":              RSIFilter,
	"rsi14day":         RSIFilter,
	"gap":              GapFilter,
	"sma20":            SMA20Filter,
	"sma50":            SMA50Filter,
	"sma200":           SMA200Filter,
	"change":           ChangeFilter,
	"changefromopen":   ChangeFromOpenFilter,
	"20dayhighlow":     HighLow20DayFilter,
	"highlow20day":     HighLow20DayFilter,
	"20dayhl":          HighLow20DayFilter,
	"50dayhighlow":     HighLow50DayFilter,
	"highlow50day":     HighLow50DayFilter,
	"50dayhl":          HighLow50DayFilter,
	"52weekhighlow":    HighLow52WeekFilter,
	"highlow52week":    HighLow52WeekFilter,
	"52weekhl":         HighLow52WeekFilter,
	"pattern":          PatternFilter,
	"candlestick":      CandlestickFilter,
	"beta":             BetaFilter,
	"averagetruerange": AverageTrueRangeFilter,
	"atr":              AverageTrueRangeFilter,
}
