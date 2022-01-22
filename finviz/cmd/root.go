// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/d3an/finviz/finviz/cmd/calendar"
	"github.com/d3an/finviz/finviz/cmd/earnings"
	"github.com/d3an/finviz/finviz/cmd/news"
	"github.com/d3an/finviz/finviz/cmd/quote"
	"github.com/d3an/finviz/finviz/cmd/screener"
)

var rootCmd = &cobra.Command{
	Use:   "finviz",
	Short: "This is an unofficial CLI for FinViz.com",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
	},
}

// Execute runs the rootCmd, i.e. finviz
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(screener.Cmd)
	rootCmd.AddCommand(news.Cmd)
	rootCmd.AddCommand(quote.Cmd)
	rootCmd.AddCommand(calendar.Cmd)
	rootCmd.AddCommand(earnings.Cmd)
}
