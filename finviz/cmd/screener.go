// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"strings"

	"github.com/d3an/finviz"
	"github.com/spf13/cobra"
)

var (
	signalArg     string
	orderArg      string
	tickerArgs    []string
	filterArgs    []string
	outputCSVArg  string
	outputJSONArg string
	viewArg       string

	screenerCmd = &cobra.Command{
		Use:     "screener",
		Aliases: []string{"screen", "scr"},
		Short:   "Finviz Stock Screener.",
		Long: "Finviz Stock Screener searches through large amounts of stock data and returns a list " +
			"of stocks that match one or more selected criteria.",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var signal finviz.SignalType
			var generalOrder finviz.GeneralOrderType
			var specificOrder finviz.SpecificOrderType
			var filters []finviz.FilterInterface

			// Handle signal
			if signalArg == "" {
				signal = ""
			} else {
				signal, err = finviz.GetSignal(strings.ToLower(signalArg))
				if err != nil {
					er(err)
				}
			}

			// Handle general order
			generalOrder = finviz.GetGeneralOrder(strings.ToLower(orderArg))
			if generalOrder == finviz.Descending {
				orderArg = strings.TrimPrefix(orderArg, "-")
			}

			// Handle specific order
			if orderArg == "" {
				specificOrder = ""
			} else {
				specificOrder, err = finviz.GetSpecificOrder(strings.ToLower(orderArg))
				if err != nil {
					er(err)
				}
			}

			// Cobra parses tickers directly to slice

			// Handle filters
			if filterCount := len(filterArgs); filterCount == 0 {
				filters = nil
			} else {
				for i := 0; i < filterCount; i++ {
					var filterQuery string
					var filterValues []string
					var filter *finviz.Filter

					filterQuery, filterValues, err = extractFilterInput(filterArgs[i])
					if err != nil {
						er(err)
					}

					filter, err = finviz.GetFilter(strings.ToLower(filterQuery), filterValues...)
					if err != nil {
						er(err)
					}
					filters = append(filters, filter)
				}
			}

			client := finviz.NewClient()
			df, err := finviz.RunScreen(client, finviz.ScreenInput{
				Signal:        signal,
				GeneralOrder:  generalOrder,
				SpecificOrder: specificOrder,
				Tickers:       tickerArgs,
				Filters:       filters,
				View:          viewArg,
			})
			if err != nil {
				er(err)
			}

			if outputCSVArg != "" {
				if err := finviz.ExportScreenCSV(df, outputCSVArg); err != nil {
					er(err)
				}
			} else if outputJSONArg != "" {
				if err := finviz.ExportScreenJSON(df, outputJSONArg); err != nil {
					er(err)
				}
			} else {
				finviz.PrintFullDataframe(df)
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
	screenerCmd.Flags().StringVarP(&signalArg, "signal", "s", "", "TopGainers")
	screenerCmd.Flags().StringVarP(&orderArg, "order", "o", "", "DividendYield")
	screenerCmd.Flags().StringVarP(&viewArg, "view", "v", "110", "510")
	screenerCmd.Flags().StringSliceVarP(&tickerArgs, "tickers", "t", nil, "AAPL,GS,VIRT")
	screenerCmd.Flags().StringArrayVarP(&filterArgs, "filter", "f", nil, "Industry:Gold,Airlines,\"Aerospace & Defense\",Airlines")
	screenerCmd.Flags().StringVar(&outputCSVArg, "output-csv", "", "outputFileName.csv")
	screenerCmd.Flags().StringVar(&outputJSONArg, "output-json", "", "outputFileName.json")
}
