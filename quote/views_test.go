// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package quote

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
	// "github.com/go-gota/gota/dataframe"
	"github.com/stretchr/testify/require"

	// "github.com/d3an/finviz"
	"github.com/d3an/finviz/utils"
)

func newTestClient(config *Config) *Client {
	if config != nil {
		return &Client{
			Client: &http.Client{
				Timeout:   30 * time.Second,
				Transport: utils.AddHeaderTransport(config.recorder),
			},
		}
	}
	return &Client{
		Client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 30 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout: 30 * time.Second,
			},
		},
	}
}

func TestGenerateURL(t *testing.T) {
	values := []struct {
		ticker   string
		expected string
	}{
		{
			ticker:   "aapl",
			expected: fmt.Sprintf("%s?t=AAPL&ty=c&p=d&b=1", APIURL),
		},
		{
			ticker:   "A",
			expected: fmt.Sprintf("%s?t=A&ty=c&p=d&b=1", APIURL),
		},
	}

	for _, v := range values {
		url, err := GenerateURL(v.ticker)
		require.Nil(t, err)
		require.Equal(t, v.expected, url)
	}
}

func TestGetData(t *testing.T) {
	values := []struct {
		cassettePath        string
		ticker              string
		expectedColCount    int
		expectedMissingCols []string
	}{
		{ // Full column count
			cassettePath:        "cassettes/full_quote",
			ticker:              "AAPL",
			expectedColCount:    83,
			expectedMissingCols: []string{},
		},
		{ // No Insider Trading or Analyst Recommendation table
			cassettePath:        "cassettes/missing_insdr_and_recom",
			ticker:              "ATHE",
			expectedColCount:    83,
			expectedMissingCols: []string{"Insider Trading", "Analyst Recommendations"},
		},
		{ // No Insider Trading table
			cassettePath:        "cassettes/missing_insdr",
			ticker:              "AEZS",
			expectedColCount:    83,
			expectedMissingCols: []string{"Insider Trading"},
		},
	}

	for _, v := range values {
		func() {
			r, err := recorder.New(v.cassettePath)
			require.Nil(t, err)
			defer func() {
				err = r.Stop()
				require.Nil(t, err)
			}()
			client := newTestClient(&Config{recorder: r, userAgent: uarand.GetRandom()})

			df, err := client.GetQuotes([]string{v.ticker})
			require.Nil(t, err)
			require.Equal(t, v.expectedColCount, df.Ncol())
			for name := range v.expectedMissingCols {
				require.NotContains(t, df.Names(), name)
			}
		}()
	}
}
