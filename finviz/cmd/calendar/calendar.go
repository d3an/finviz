// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package calendar

import (
	"github.com/spf13/cobra"

	"github.com/d3an/finviz/calendar"
	"github.com/d3an/finviz/utils"
)

var (
	outFile string

	// Cmd is the CLI subcommand for Finviz news
	Cmd = &cobra.Command{
		Use:     "calendar",
		Aliases: []string{"cal", "ec"},
		Short:   "Finviz Economic Calendar",
		Long:    "Finviz Economic Calendar returns this week's Economic Calendar.",
		Run: func(cmd *cobra.Command, args []string) {
			client := calendar.New(nil)
			df, err := client.GetCalendar()
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
