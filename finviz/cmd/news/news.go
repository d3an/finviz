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
	outFile string
	view    *utils.Enum

	// Cmd is the CLI subcommand for Finviz news
	Cmd = &cobra.Command{
		Use:     "news",
		Aliases: []string{"ns"},
		Short:   "Finviz News",
		Long:    "Finviz News returns the latest news.",
		Run: func(cmd *cobra.Command, args []string) {
			client := news.New(nil)
			df, err := client.GetNews(view.Value)
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
	// -v time|source
	// -o <filename>
	view = utils.NewEnum([]string{"time", "source"}, "time")
	Cmd.Flags().VarP(view, "view", "v", "time|source")
	Cmd.Flags().StringVarP(&outFile, "outfile", "o", "", "output.(csv|json)")
}
