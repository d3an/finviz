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
	outFile string
	tickers []string

	// Cmd is the CLI subcommand for Finviz news
	Cmd = &cobra.Command{
		Use:     "quote",
		Aliases: []string{"q", "quotes"},
		Short:   "Finviz Quotes",
		Long:    "Finviz Quotes returns the quotes for tickers provided.",
		Run: func(cmd *cobra.Command, args []string) {
			client := quote.New(nil)
			results, err := client.GetQuotes(tickers)
			if err != nil {
				utils.Err(err)
			}

			if err = utils.ExportData(results.Data, outFile); err != nil {
				utils.Err(err)
			}
		},
	}
)

func init() {
	// -t aapl,amzn,tsla
	// -o <filename>
	Cmd.Flags().StringSliceVarP(&tickers, "tickers", "t", nil, "AAPL,GS,amzn")
	Cmd.Flags().StringVarP(&outFile, "outfile", "o", "", "output.(csv|json)")
}
