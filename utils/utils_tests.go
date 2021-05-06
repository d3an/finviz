// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package utils

/*
func TestExportScreenCSV(t *testing.T) {
	r, err := recorder.New("fixtures/finviz_screener")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := r.Stop(); err != nil {
			t.Error(err)
		}
	}()

	client := newTestingClient(r)
	df, err := RunScreen(client, ScreenInput{
		Signal:        AllStocks,
		GeneralOrder:  Descending,
		SpecificOrder: ChangeFromOpen,
		Filters: []FilterInterface{
			IndustryFilter(WasteManagement, Airlines),
			AverageVolumeFilter(AvgVolOver50K),
			PriceFilter(PriceOver1),
		},
	})
	if err != nil {
		t.Error(err)
	}

	err = ExportScreenCSV(df, "test_data/finviz_screener_results.csv")
	if err != nil {
		t.Error(err)
	}
}

func TestExportScreenJSON(t *testing.T) {
	r, err := recorder.New("fixtures/finviz_screener")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := r.Stop(); err != nil {
			t.Error(err)
		}
	}()

	client := newTestingClient(r)
	df, err := RunScreen(client, ScreenInput{
		Signal:        AllStocks,
		GeneralOrder:  Descending,
		SpecificOrder: ChangeFromOpen,
		Filters: []FilterInterface{
			IndustryFilter(WasteManagement, Airlines),
			AverageVolumeFilter(AvgVolOver50K),
			PriceFilter(PriceOver1),
		},
	})
	if err != nil {
		t.Error(err)
	}

	err = ExportScreenJSON(df, "test_data/finviz_screener_results.json")
	if err != nil {
		t.Error(err)
	}
}
*/
