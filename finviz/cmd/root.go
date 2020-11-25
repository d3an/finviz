// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"
	"github.com/d3an/finviz/finviz/cmd/news"
	"github.com/d3an/finviz/finviz/cmd/quote"
	"github.com/d3an/finviz/finviz/cmd/screener"
	"github.com/spf13/cobra"
	"os"
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
	rootCmd.AddCommand(screener.ScreenerCmd)
	rootCmd.AddCommand(news.NewsCmd)
	rootCmd.AddCommand(quote.QuoteCmd)
}
