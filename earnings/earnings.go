// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package earnings

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-gota/gota/dataframe"

	"github.com/d3an/finviz/utils"
)

const (
	APIURL = "https://finviz.com"
)

var (
	once     sync.Once
	instance *Client
)

type Config struct {
	userAgent string
	recorder  *recorder.Recorder
}

type Client struct {
	*http.Client
	config Config
}

func New(config *Config) *Client {
	once.Do(func() {
		transport := &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 30 * time.Second,
		}
		client := &http.Client{
			Timeout:   30 * time.Second,
			Transport: transport,
		}
		if config != nil {
			instance = &Client{Client: client, config: *config}
		}
		instance = &Client{
			Client: client,
			config: Config{userAgent: uarand.GetRandom()},
		}
	})

	return instance
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", c.config.userAgent)
	return c.Client.Do(req)
}

func (c *Client) GetEarnings() (*dataframe.DataFrame, error) {
	req, err := http.NewRequest(http.MethodGet, APIURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting earnings, status code: '%d', body: '%s'", resp.StatusCode, string(body))
	}

	doc, err := utils.GenerateDocument(body)
	if err != nil {
		return nil, err
	}

	results, err := Scrape(doc)
	if err != nil {
		return nil, err
	}

	df := dataframe.LoadRecords(results)
	return &df, nil
}

func Scrape(doc *goquery.Document) ([][]string, error) {
	data := doc.Find("#homepage_bottom").Find("table[class=\"t-home-table\"] > tbody").Eq(5)

	var earningsDataSlice []map[string]interface{}
	data.Children().Each(func(i int, row *goquery.Selection) {
		if i == 0 {
			return
		}
		var date string
		row.Children().Each(func(j int, rowItem *goquery.Selection) {
			switch j {
			case 0:
				date = rowItem.Text()
			case 1:
				return
			default:
				if rowItem.AttrOr("title", "") == "" {
					return
				}
				earningsDataSlice = append(earningsDataSlice, map[string]interface{}{"Date": date, "Ticker": rowItem.Text()})
			}
		})
	})

	headers := []string{"Date", "Ticker"}
	return utils.GenerateRows(headers, earningsDataSlice)
}
