// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package quote

import (
	"github.com/spf13/cobra"

	"github.com/d3an/finviz/quote"
	"github.com/d3an/finviz/utils"
)

var (
	outputCSVArg  string
	outputJSONArg string
	tickerArgs    []string

	// Cmd is the CLI subcommand for FinViz news
	Cmd = &cobra.Command{
		Use:     "quote",
		Aliases: []string{"q"},
		Short:   "FinViz Quotes.",
		Long:    "FinViz Quotes returns the quotes for tickers provided.",
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			client := quote.New(nil)
			results, err := client.GetQuotes(tickerArgs)
			if err != nil {
				utils.Err(err)
			}

			if outputCSVArg != "" {
				if err := utils.ExportCSV(results.Data, outputCSVArg); err != nil {
					utils.Err(err)
				}
			} else if outputJSONArg != "" {
				if err := utils.ExportJSON(results.Data, outputJSONArg); err != nil {
					utils.Err(err)
				}
			} else {
				utils.PrintFullDataFrame(results.Data)
			}
		},
	}
)

func init() {
	// -v 1
	// --output-csv data.csv
	// --output-json data.json
	Cmd.Flags().StringSliceVarP(&tickerArgs, "tickers", "t", nil, "AAPL,GS,amzn")
	Cmd.Flags().StringVar(&outputCSVArg, "output-csv", "", "outputFileName.csv")
	Cmd.Flags().StringVar(&outputJSONArg, "output-json", "", "outputFileName.json")
}
