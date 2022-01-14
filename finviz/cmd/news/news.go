// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package news

import (
	"github.com/spf13/cobra"

	"github.com/d3an/finviz/news"
	"github.com/d3an/finviz/utils"
)

var (
	outputCSVArg  string
	outputJSONArg string
	viewArg       string

	// Cmd is the CLI subcommand for FinViz news
	Cmd = &cobra.Command{
		Use:     "news",
		Aliases: []string{"ns"},
		Short:   "FinViz News.",
		Long:    "FinViz News returns the latest news.",
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			client := news.New(nil)
			df, err := client.GetNews(viewArg)
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
	// -v 1
	// --output-csv data.csv
	// --output-json data.json
	Cmd.Flags().StringVarP(&viewArg, "view", "v", "1", "2")
	Cmd.Flags().StringVar(&outputCSVArg, "output-csv", "", "outputFileName.csv")
	Cmd.Flags().StringVar(&outputJSONArg, "output-json", "", "outputFileName.json")
}
