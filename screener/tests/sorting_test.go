// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package tests

import (
	"github.com/d3an/finviz/screener"
	"testing"
)

func TestGetGeneralOrder(t *testing.T) {
	testInputs := []struct {
		query    string
		expected screener.GeneralOrderType
	}{
		{
			"",
			screener.Ascending,
		},
		{
			"-ticker",
			screener.Descending,
		},
		{
			"-",
			screener.Descending,
		},
		{
			"ticker",
			screener.Ascending,
		},
	}

	for _, testInput := range testInputs {
		genOrder := screener.GetGeneralOrder(testInput.query)
		if genOrder != testInput.expected {
			t.Fail()
			t.Logf("Expected GeneralOrderType: \"%v\", Received GeneralOrderType: \"%v\"", testInput.expected, genOrder)
		}
	}
}

func TestGetSpecificOrder(t *testing.T) {
	testInputs := []struct {
		query    string
		expected screener.SpecificOrderType
	}{
		{
			"",
			screener.Ticker,
		},
		{
			"Fwd P/E",
			screener.ForwardPE,
		},
		{
			"rsi",
			screener.RSI14Day,
		},
		{
			"Ticker",
			screener.Ticker,
		},
	}

	for _, testInput := range testInputs {
		if specOrder, err := screener.GetSpecificOrder(testInput.query); err != nil {
			t.Fail()
			t.Log(err)
		} else if specOrder != testInput.expected {
			t.Fail()
			t.Logf("Expected SpecificOrderType: \"%v\", Received SpecificOrderType: \"%v\"", testInput.expected, specOrder)
		}
	}
}
