// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package calendar

import (
	"github.com/spf13/cobra"

	"github.com/d3an/finviz/earnings"
	"github.com/d3an/finviz/utils"
)

var (
	outputCSVArg  string
	outputJSONArg string

	// Cmd is the CLI subcommand for FinViz news
	Cmd = &cobra.Command{
		Use:     "earnings",
		Aliases: []string{"e"},
		Short:   "Finviz Earnings.",
		Long:    "Finviz Earnings returns the tickers with earnings releases left this week.",
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			client := earnings.New(nil)
			df, err := client.GetEarnings()
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
}
