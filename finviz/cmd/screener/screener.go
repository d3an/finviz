// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package screener

import (
	"github.com/spf13/cobra"

	. "github.com/d3an/finviz/screener"
	"github.com/d3an/finviz/utils"
)

var (
	url           string
	outputCSVArg  string
	outputJSONArg string

	// Cmd is the CLI subcommand for the Screener app
	Cmd = &cobra.Command{
		Use:     "screener <url>",
		Aliases: []string{"screen", "scr"},
		Short:   "FinViz Stock Screener.",
		Long: "FinViz Stock Screener searches through large amounts of stock data and returns a list " +
			"of stocks that match one or more selected criteria.",
		Run: func(cmd *cobra.Command, args []string) {
			if url == "" {
				utils.Err("URL not provided")
			}

			client := New(nil)
			df, err := client.GetScreenerResults(url)
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
	// --output-csv data.csv
	// --output-json data.json
	Cmd.Flags().StringVar(&outputCSVArg, "output-csv", "", "outputFileName.csv")
	Cmd.Flags().StringVar(&outputJSONArg, "output-json", "", "outputFileName.json")
	Cmd.Flags().StringVar(&url, "url", "", "https://finviz.com/screener.ashx?v=110&f=exch_nyse")
}
