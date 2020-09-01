// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "finviz",
	Short: "FinViz is an unofficial CLI for FinViz.com",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			er(err)
		}
	},
}

// Execute runs the rootCmd, i.e. finviz
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(screenerCmd)
}
