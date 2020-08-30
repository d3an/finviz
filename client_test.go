// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

/*
func TestScreener(t *testing.T) {
	r, err := recorder.New("fixtures/finviz_screener_view_150_all")
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
		SpecificOrder: Change,
		View:          CustomView,
		CustomColumns: []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59", "60", "61", "62", "63", "64", "65", "66", "67", "68", "69", "70"},
	})
	if err != nil {
		t.Error(err)
	}

	PrintFullDataframe(df)
}
*/
