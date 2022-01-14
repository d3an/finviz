// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package utils

import "net/http"

// headerTransport implements a Transport that can have its RoundTripper interface modified
type headerTransport struct {
	T http.RoundTripper
}

// RoundTrip implements the RoundTripper interface with a custom user-agent
func (adt *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return adt.T.RoundTrip(req)
}

// AddHeaderTransport adds the an HTTP recorder to the request for testing purposes
func AddHeaderTransport(t http.RoundTripper) *headerTransport {
	if t == nil {
		t = http.DefaultTransport
	}
	return &headerTransport{t}
}
