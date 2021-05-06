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
