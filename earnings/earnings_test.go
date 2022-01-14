// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package earnings

import (
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/corpix/uarand"
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

func TestGetEarnings(t *testing.T) {
	func() {
		r, err := recorder.New("cassettes/earnings")
		require.Nil(t, err)
		defer func() {
			err = r.Stop()
			require.Nil(t, err)
		}()
		client := newTestClient(&Config{recorder: r, userAgent: uarand.GetRandom()})

		earnings, err := client.GetEarnings()
		require.Nil(t, err)
		utils.PrintFullDataFrame(earnings)
	}()
}
