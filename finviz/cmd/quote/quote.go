// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package quote

import (
	"github.com/d3an/finviz"
	"github.com/d3an/finviz/quote"
	"github.com/spf13/cobra"
)

var (
	outputCSVArg  string
	outputJSONArg string
	tickerArgs    []string

	// QuoteCmd is the CLI subcommand for FinViz news
	QuoteCmd = &cobra.Command{
		Use:     "quote",
		Aliases: []string{"q"},
		Short:   "FinViz Quotes.",
		Long:    "FinViz Quotes returns the quotes for tickers provided.",
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			df, err := quote.GetQuoteData(nil, &map[string]interface{}{
				"tickers": tickerArgs,
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
				finviz.PrintFullDataFrame(df)
			}
		},
	}
)

func init() {
	// -v 1
	// --output-csv data.csv
	// --output-json data.json
	QuoteCmd.Flags().StringSliceVarP(&tickerArgs, "tickers", "t", nil, "AAPL,GS,amzn")
	QuoteCmd.Flags().StringVar(&outputCSVArg, "output-csv", "", "outputFileName.csv")
	QuoteCmd.Flags().StringVar(&outputJSONArg, "output-json", "", "outputFileName.json")
}
