// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package news

import (
	"github.com/d3an/finviz"
	"github.com/d3an/finviz/news"
	"github.com/spf13/cobra"
)

var (
	outputCSVArg  string
	outputJSONArg string
	viewArg       string

	// NewsCmd is the CLI subcommand for FinViz news
	NewsCmd = &cobra.Command{
		Use:     "news",
		Aliases: []string{"ns"},
		Short:   "FinViz News.",
		Long:    "FinViz News returns the latest news.",
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			viewInterface, err := ViewFactory(viewArg)
			if err != nil {
				er(err)
			}

			df, err := news.GetNewsData(nil, viewInterface)
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
	NewsCmd.Flags().StringVarP(&viewArg, "view", "v", "1", "2")
	NewsCmd.Flags().StringVar(&outputCSVArg, "output-csv", "", "outputFileName.csv")
	NewsCmd.Flags().StringVar(&outputJSONArg, "output-json", "", "outputFileName.json")
}
