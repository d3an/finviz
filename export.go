// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import (
	"github.com/go-gota/gota/dataframe"
	"os"
)

// ExportScreenCSV generates a csv file of the Finviz screen results
func ExportScreenCSV(df *dataframe.DataFrame, outFileName string) error {
	f, err := os.Create(outFileName)
	if err != nil {
		return err
	}

	err = df.WriteCSV(f)
	if err != nil {
		return err
	}
	return nil
}

// ExportScreenJSON generates a json file of the Finviz screen results
func ExportScreenJSON(df *dataframe.DataFrame, outFileName string) error {
	f, err := os.Create(outFileName)
	if err != nil {
		return err
	}

	err = df.WriteJSON(f)
	if err != nil {
		return err
	}
	return nil
}
