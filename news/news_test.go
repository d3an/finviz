// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package news

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/require"

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
		view     string
		expected string
	}{
		{
			view:     "by_time",
			expected: APIURL,
		},
		{
			view:     "by_source",
			expected: fmt.Sprintf("%s?v=2", APIURL),
		},
	}

	for _, v := range values {
		url, err := GenerateURL(v.view)
		require.Nil(t, err)
		require.Equal(t, v.expected, url)
	}
}

func TestGetData(t *testing.T) {
	values := []struct {
		cassettePath     string
		view             string
		expectedColCount int
		expectedColNames []string
	}{
		{
			cassettePath:     "cassettes/source_news_view",
			view:             "by_source",
			expectedColCount: 6,
			expectedColNames: []string{"Article Date", "Article Title", "Article URL", "Source Name", "Source URL", "News Type"},
		},
		{
			cassettePath:     "cassettes/time_news_view",
			view:             "by_time",
			expectedColCount: 5,
			expectedColNames: []string{"Article Date", "Article Title", "Article URL", "Source Name", "News Type"},
		},
	}

	for _, v := range values {
		r, err := recorder.New(v.cassettePath)
		require.Nil(t, err)
		client := newTestClient(&Config{recorder: r})

		df, err := client.GetNews(v.view)
		require.Nil(t, err)
		require.Equal(t, v.expectedColCount, df.Ncol())
		require.Equal(t, v.expectedColNames, df.Names())
		require.GreaterOrEqual(t, df.Nrow(), 100, "Expected at least 100 rows")

		err = r.Stop()
		require.Nil(t, err)

		utils.PrintFullDataFrame(df)
	}
}
