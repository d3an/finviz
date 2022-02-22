// Copyright (c) 2022 James Bury. All rights reserved.
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
	outFile string

	// Cmd is the CLI subcommand for the Screener app
	Cmd = &cobra.Command{
		Use:     "screener <url>",
		Aliases: []string{"screen", "scr"},
		Short:   "Finviz Stock Screener",
		Long: "Finviz Stock Screener searches through large amounts of stock data and returns a list " +
			"of stocks that match one or more selected criteria.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				utils.Err("URL not provided")
			}

			client := New(nil)
			df, err := client.GetScreenerResults(args[0])
			if err != nil {
				utils.Err(err)
			}

			if err = utils.ExportData(df, outFile); err != nil {
				utils.Err(err)
			}
		},
	}
)

func init() {
	Cmd.Flags().StringVarP(&outFile, "outfile", "o", "", "output.(csv|json)")
}
