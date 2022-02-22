// Copyright (c) 2022 James Bury. All rights reserved.
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

	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stretchr/testify/require"

	"github.com/d3an/finviz/utils/test"
)

func newTestClient(config *Config) *Client {
	if config != nil {
		return &Client{
			Client: &http.Client{
				Timeout:   30 * time.Second,
				Transport: test.AddHeaderTransport(config.recorder),
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
			view:     "time",
			expected: APIURL,
		},
		{
			view:     "source",
			expected: fmt.Sprintf("%s?v=2", APIURL),
		},
	}

	for _, v := range values {
		url, err := GenerateURL(v.view)
		require.Nil(t, err)
		require.Equal(t, v.expected, url)
	}
}

func TestGetNews(t *testing.T) {
	values := []struct {
		cassettePath     string
		view             string
		expectedColCount int
		expectedColNames []string
	}{
		{
			cassettePath:     "cassettes/source_news_view",
			view:             "source",
			expectedColCount: 6,
			expectedColNames: []string{"Article Date", "Article Title", "Article URL", "Source Name", "Source URL", "News Type"},
		},
		{
			cassettePath:     "cassettes/time_news_view",
			view:             "time",
			expectedColCount: 5,
			expectedColNames: []string{"Article Date", "Article Title", "Article URL", "Source Name", "News Type"},
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

			df, err := client.GetNews(v.view)
			require.Nil(t, err)
			require.Equal(t, v.expectedColCount, df.Ncol())
			require.Equal(t, v.expectedColNames, df.Names())
			require.GreaterOrEqual(t, df.Nrow(), 100, "Expected at least 100 rows")
		}()
	}
}
