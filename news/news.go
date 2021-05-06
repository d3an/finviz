// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package news

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-gota/gota/dataframe"

	"github.com/d3an/finviz/utils"
)

const (
	APIURL = "https://finviz.com/news.ashx"
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

// GetNews returns a DataFrame containing recent news data
func (c *Client) GetNews(view string) (*dataframe.DataFrame, error) {
	url, err := GenerateURL(view)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error getting url: '%s', status code: '%d', body: '%s'", url, resp.StatusCode, string(body))
	}

	doc, err := utils.GenerateDocument(body)
	if err != nil {
		return nil, err
	}

	results, err := Scrape(view, doc)
	if err != nil {
		return nil, err
	}

	df := dataframe.LoadRecords(results)
	return &df, nil
}

func GenerateURL(view string) (string, error) {
	switch view {
	case "by_time":
		return APIURL, nil
	case "by_source":
		return fmt.Sprintf("%s?v=2", APIURL), nil
	default:
		return "", fmt.Errorf("error view '%s' not found", view)
	}
}

func Scrape(view string, doc *goquery.Document) ([][]string, error) {
	switch view {
	case "by_time":
		return ByTimeScrape(doc)
	case "by_source":
		return BySourceScrape(doc)
	default:
		return nil, fmt.Errorf("error view '%s' not found", view)
	}
}

func ByTimeScrape(doc *goquery.Document) ([][]string, error) {
	var newsDataSlice []map[string]interface{}

	doc.Find("#news > div").Children().Eq(1).Find("tbody").Eq(0).Children().Eq(1).Children().Each(func(i int, newsColumn *goquery.Selection) {
		var newsType string

		switch i {
		case 0:
			newsType = "news"
		case 2:
			newsType = "blog"
		}

		if i != 1 {
			newsColumn.Find("tbody").Eq(0).Children().Each(func(j int, newsItem *goquery.Selection) {
				var rawNewsData = make(map[string]interface{})
				rawNewsData["Article Date"] = newsItem.Children().Eq(1).Text()
				rawNewsData["Article Title"] = newsItem.Children().Eq(2).Children().Eq(0).Text()
				rawNewsData["Article URL"] = newsItem.Children().Eq(2).Children().Eq(0).AttrOr("href", "")
				if classes := newsItem.Children().Eq(0).AttrOr("class", ""); classes != "" {
					if val, exists := NewsSourceAttributeLookup[strings.Split(classes, " ")[1]]; exists {
						rawNewsData["Source Name"] = val
					} else {
						rawNewsData["Source Name"] = "Unknown"
					}
				}
				rawNewsData["News Type"] = newsType
				newsDataSlice = append(newsDataSlice, rawNewsData)
			})
		}
	})

	headers := []string{"Article Date", "Article Title", "Article URL", "Source Name", "News Type"}
	return utils.GenerateRows(headers, newsDataSlice)
}

func BySourceScrape(doc *goquery.Document) ([][]string, error) {
	var newsDataSlice []map[string]interface{}

	for tableIndex := 2; tableIndex <= 4; tableIndex++ {
		var newsType string

		switch tableIndex {
		case 2:
			newsType = "news"
		case 4:
			newsType = "blog"
		}

		if tableIndex != 3 {
			doc.Find("#news > div").Children().Eq(tableIndex).Find("tr").Children().Each(func(i int, newsColumn *goquery.Selection) {
				if align := newsColumn.AttrOr("align", ""); align != "" {
					newsColumn.Children().Each(func(j int, newsSource *goquery.Selection) {
						var sourceName string
						var sourceURL string
						newsSource.Find("tbody").Eq(0).Children().Each(func(k int, item *goquery.Selection) {
							if k == 1 {
								sourceItem := item.Find("tr").Children().Eq(1).Find("a")
								sourceName = sourceItem.Text()
								sourceURL = sourceItem.AttrOr("href", "")
							} else if k > 1 {
								var rawNewsData = make(map[string]interface{})
								rawNewsData["Article Date"] = item.Children().Eq(0).Text()
								rawNewsData["Article Title"] = item.Children().Eq(1).Find("a").Text()
								rawNewsData["Article URL"] = item.Children().Eq(1).Children().Eq(0).AttrOr("href", "")
								rawNewsData["Source Name"] = sourceName
								rawNewsData["Source URL"] = sourceURL
								rawNewsData["News Type"] = newsType
								newsDataSlice = append(newsDataSlice, rawNewsData)
							}
						})
					})
				}
			})
		}
	}

	headers := []string{"Article Date", "Article Title", "Article URL", "Source Name", "Source URL", "News Type"}
	return utils.GenerateRows(headers, newsDataSlice)
}

// NewsSourceAttributeLookup provides an interface for identifying news sources based on their CSS attributes
var NewsSourceAttributeLookup = map[string]string{
	"is-1":   "MarketWatch",
	"is-2":   "WSJ",
	"is-3":   "Reuters",
	"is-4":   "Yahoo Finance",
	"is-5":   "CNN",
	"is-6":   "The New York Times",
	"is-7":   "Bloomberg",
	"is-9":   "BBC",
	"is-10":  "CNBC",
	"is-11":  "Fox Business",
	"is-102": "Mish's Global Economic Trend Analysis",
	"is-105": "Trader Feed",
	"is-113": "Howard Lindzon",
	"is-114": "Seeking Alpha",
	"is-121": "The Disciplined Investor",
	"is-123": "Fallond Stock Picks",
	"is-132": "Zero Hedge",
	"is-133": "market folly",
	"is-136": "Daily Reckoning",
	"is-137": "Vantage Point Trading",
	"is-141": "Abnormal Returns",
	"is-142": "Calculated Risk",
}
