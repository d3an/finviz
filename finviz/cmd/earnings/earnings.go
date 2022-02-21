// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package earnings

import (
	"github.com/spf13/cobra"

	"github.com/d3an/finviz/earnings"
	"github.com/d3an/finviz/utils"
)

var (
	outFile string

	// Cmd is the CLI subcommand for Finviz news
	Cmd = &cobra.Command{
		Use:     "earnings",
		Aliases: []string{"e"},
		Short:   "Finviz Earnings",
		Long:    "Finviz Earnings returns the tickers with earnings releases left this week.",
		Run: func(cmd *cobra.Command, args []string) {
			client := earnings.New(nil)
			df, err := client.GetEarnings()
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
	// -o <filename>
	Cmd.Flags().StringVarP(&outFile, "outfile", "o", "", "output.(csv|json)")
}
