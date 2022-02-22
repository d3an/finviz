// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package calendar

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-gota/gota/dataframe"

	"github.com/d3an/finviz/utils"
)

const (
	APIURL = "https://finviz.com/calendar.ashx"
	YEAR   = 2022
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

func (c *Client) GetCalendar() (*dataframe.DataFrame, error) {
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
		return nil, fmt.Errorf("error getting calendar, status code: '%d', body: '%s'", resp.StatusCode, string(body))
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
	var year int
	var err error

	year, err = strconv.Atoi(doc.Find(".copyright").Text()[89:93])
	if err != nil {
		year = YEAR
	}

	data := doc.Find("tr .calendar-header")

	var dow string
	var calendarDataSlice []map[string]interface{}
	data.Each(func(i int, day *goquery.Selection) {
		dow = day.Children().Eq(0).Text()

		day.Parent().Children().Each(func(j int, event *goquery.Selection) {
			if j == 0 || len(event.Children().Nodes) < 9 {
				return
			}
			var calendarEvent = make(map[string]interface{})
			event.Children().Each(func(k int, eventDetails *goquery.Selection) {
				switch k {
				case 0:
					date, err := time.Parse("Mon Jan 02 2006 3:04 PM MST", fmt.Sprintf("%s %d %s EST", dow, year, eventDetails.Text()))
					if err != nil {
						calendarEvent["Date"] = eventDetails.Text()
						return
					}
					calendarEvent["Date"] = date.Format(time.RFC3339)
				case 1:
					return
				case 2:
					calendarEvent["Release"] = eventDetails.Text()
				case 3:
					switch eventDetails.Find("img").Eq(0).AttrOr("src", "") {
					default:
						calendarEvent["Impact"] = "-"
					case "gfx/calendar/impact_3.gif":
						calendarEvent["Impact"] = "critical"
					case "gfx/calendar/impact_2.gif":
						calendarEvent["Impact"] = "moderate"
					case "gfx/calendar/impact_1.gif":
						calendarEvent["Impact"] = "low"
					}
				case 4:
					calendarEvent["For"] = eventDetails.Text()
				case 5:
					calendarEvent["Actual"] = eventDetails.Text()
				case 6:
					calendarEvent["Expected"] = eventDetails.Text()
				case 7:
					calendarEvent["Prior"] = eventDetails.Text()
				}
			})
			calendarDataSlice = append(calendarDataSlice, calendarEvent)
		})
	})

	headers := []string{"Date", "Release", "Impact", "For", "Actual", "Expected", "Prior"}
	return utils.GenerateRows(headers, calendarDataSlice)
}
