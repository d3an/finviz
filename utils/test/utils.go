package test

import (
	"net"
	"net/http"
	"time"

	"github.com/dnaeon/go-vcr/recorder"
)

// headerTransport implements a Transport that can have its RoundTripper interface modified
type headerTransport struct {
	T http.RoundTripper
}

// RoundTrip implements the RoundTripper interface with a custom user-agent
func (adt *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return adt.T.RoundTrip(req)
}

// AddHeaderTransport adds an HTTP recorder to the request for testing purposes
func AddHeaderTransport(t http.RoundTripper) *headerTransport {
	if t == nil {
		t = http.DefaultTransport
	}
	return &headerTransport{t}
}

type Config struct {
	Recorder  *recorder.Recorder
	UserAgent string
}

type Client struct {
	*http.Client
	config Config
}

func NewClient(config *Config) *Client {
	if config != nil {
		return &Client{
			Client: &http.Client{
				Timeout:   30 * time.Second,
				Transport: AddHeaderTransport(config.Recorder),
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
