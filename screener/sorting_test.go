// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package screener

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetGeneralOrder(t *testing.T) {
	values := []struct {
		query    string
		expected GeneralOrderType
	}{
		{
			"",
			Ascending,
		},
		{
			"-ticker",
			Descending,
		},
		{
			"-",
			Descending,
		},
		{
			"ticker",
			Ascending,
		},
	}

	for _, v := range values {
		genOrder := GetGeneralOrder(v.query)
		require.Equal(t, v.expected, genOrder)
	}
}

func TestGetSpecificOrder(t *testing.T) {
	values := []struct {
		query    string
		expected SpecificOrderType
	}{
		{
			"",
			Ticker,
		},
		{
			"Fwd P/E",
			ForwardPE,
		},
		{
			"rsi",
			RSI14Day,
		},
		{
			"Ticker",
			Ticker,
		},
	}

	for _, v := range values {
		specOrder, err := GetSpecificOrder(v.query)
		require.Nil(t, err)
		require.Equal(t, v.expected, specOrder)
	}
}
