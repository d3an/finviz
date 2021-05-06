// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package screener

import (
	"strings"

	"github.com/spf13/cobra"

	. "github.com/d3an/finviz/screener"
	"github.com/d3an/finviz/utils"
)

var (
	signalArg     string
	orderArg      string
	tickerArgs    []string
	filterArgs    []string
	outputCSVArg  string
	outputJSONArg string
	viewArg       string

	// ScreenerCmd is the CLI subcommand for the Screener app
	ScreenerCmd = &cobra.Command{
		Use:     "screener",
		Aliases: []string{"screen", "scr"},
		Short:   "FinViz Stock Screener.",
		Long: "FinViz Stock Screener searches through large amounts of stock data and returns a list " +
			"of stocks that match one or more selected criteria.",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var signal SignalType
			var generalOrder GeneralOrderType
			var specificOrder SpecificOrderType
			var filters []FilterInterface

			// Handle signal
			if signalArg == "" {
				signal = ""
			} else {
				signal, err = GetSignal(strings.ToLower(signalArg))
				if err != nil {
					utils.Err(err)
				}
			}

			// Handle general order
			generalOrder = GetGeneralOrder(strings.ToLower(orderArg))
			if generalOrder == Descending {
				orderArg = strings.TrimPrefix(orderArg, "-")
			}

			// Handle specific order
			if orderArg == "" {
				specificOrder = ""
			} else {
				specificOrder, err = GetSpecificOrder(strings.ToLower(orderArg))
				if err != nil {
					utils.Err(err)
				}
			}

			// Handle filters
			if filterCount := len(filterArgs); filterCount == 0 {
				filters = nil
			} else {
				for i := 0; i < filterCount; i++ {
					var filterQuery string
					var filterValues []string
					var filter *Filter

					filterQuery, filterValues, err = extractFilterInput(filterArgs[i])
					if err != nil {
						utils.Err(err)
					}

					filter, err = GetFilter(strings.ToLower(filterQuery), filterValues...)
					if err != nil {
						utils.Err(err)
					}
					filters = append(filters, filter)
				}
			}

			client := New(nil)

			df, err := client.GetScreenerResults(viewArg, map[string]interface{}{
				"signal":         signal,
				"general_order":  generalOrder,
				"specific_order": specificOrder,
				"tickers":        tickerArgs,
				"filters":        filters,
			})
			if err != nil {
				utils.Err(err)
			}

			if outputCSVArg != "" {
				if err := utils.ExportCSV(df, outputCSVArg); err != nil {
					utils.Err(err)
				}
			} else if outputJSONArg != "" {
				if err := utils.ExportJSON(df, outputJSONArg); err != nil {
					utils.Err(err)
				}
			} else {
				utils.PrintFullDataFrame(df)
			}
		},
	}
)

func init() {
	// -s TopGainers
	// -ord -ChangeFromOpen
	// -t AAPL,GS,VIRT
	// -f Industry:gold,airlines -f Sector:Materials
	// -v 510
	// --output-csv data.csv
	// --output-json data.json
	ScreenerCmd.Flags().StringVarP(&signalArg, "signal", "s", "", "TopGainers")
	ScreenerCmd.Flags().StringVarP(&orderArg, "order", "o", "", "DividendYield")
	ScreenerCmd.Flags().StringVarP(&viewArg, "view", "v", "110", "510")
	ScreenerCmd.Flags().StringSliceVarP(&tickerArgs, "tickers", "t", nil, "AAPL,GS,VIRT")
	ScreenerCmd.Flags().StringArrayVarP(&filterArgs, "filter", "f", nil, "Industry:Gold,Airlines,\"Aerospace & Defense\",Airlines")
	ScreenerCmd.Flags().StringVar(&outputCSVArg, "output-csv", "", "outputFileName.csv")
	ScreenerCmd.Flags().StringVar(&outputJSONArg, "output-json", "", "outputFileName.json")
}
