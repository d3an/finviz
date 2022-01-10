package screener

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/cenkalti/backoff/v4"
	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-gota/gota/dataframe"
	"github.com/pkg/errors"

	"github.com/d3an/finviz/utils"
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

type scrapeResult struct {
	Results   []map[string]interface{}
	Keys      []string
	PageCount int
	Error     error
}

func scrape(view int, doc *goquery.Document) *scrapeResult {
	switch view / 10 {
	case 10, 11, 12, 13, 14, 15, 16, 17:
		return DefaultScrape(doc)
	case 20, 21:
		return ChartScrape(doc)
	case 30, 31:
		return BasicScrape(doc)
	case 32:
		return NewsScrape(doc)
	case 33:
		return DescriptionScrape(doc)
	case 34:
		return SnapshotScrape(doc)
	case 35:
		return TAScrape(doc)
	case 40, 41:
		return TickerScrape(doc)
	case 50, 51:
		return BulkScrape(doc)
	case 52:
		return BulkFullScrape(doc)
	default:
		return &scrapeResult{Error: fmt.Errorf("error view '%d' not found", view)}
	}
}

func (c *Client) GetScreenerResults(url string) (*dataframe.DataFrame, error) {
	var wg sync.WaitGroup
	var results []map[string]interface{}

	firstLook := make(chan scrapeResult, 1)
	wg.Add(1)
	go c.getData(url, &wg, &firstLook)
	wg.Wait()

	firstPage := <-firstLook
	if firstPage.Error != nil {
		return nil, errors.Wrapf(firstPage.Error, "error received while scraping screener")
	}
	results = append(results, firstPage.Results...)
	pagesLeft := firstPage.PageCount - 1
	if pagesLeft <= 0 {
		return processScrapeResults(firstPage.Keys, results)
	}
	maxRows := len(results)

	scrapeResults := make([]chan scrapeResult, pagesLeft)
	for i := range scrapeResults {
		scrapeResults[i] = make(chan scrapeResult, 1)
	}

	for i := 0; i < pagesLeft; i++ {
		wg.Add(1)
		go c.getData(fmt.Sprintf("%s&r=%d", url, (maxRows*(i+1))+1), &wg, &scrapeResults[i])
		wg.Wait()
	}

	for i := 0; i < pagesLeft; i++ {
		page := <-scrapeResults[i]
		if page.Error != nil {
			return nil, errors.Wrapf(page.Error, "error received while scraping screener: page '%d'", i+2)
		}
		results = append(results, page.Results...)
	}

	return processScrapeResults(firstPage.Keys, results)
}

func processScrapeResults(keys []string, results []map[string]interface{}) (*dataframe.DataFrame, error) {
	rows, err := utils.GenerateRows(keys, results)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate rows from KVP map")
	}
	df := dataframe.LoadRecords(rows)
	return CleanScreenerDataFrame(&df), nil
}

func (c *Client) getData(url string, wg *sync.WaitGroup, scr *chan scrapeResult) {
	defer wg.Done()
	defer close(*scr)

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		*scr <- scrapeResult{Error: err}
		return
	}

	view, err := strconv.ParseInt(req.URL.Query().Get("v"), 10, 64)
	if err != nil {
		*scr <- scrapeResult{Error: err}
		return
	}

	var body []byte

	if err := backoff.RetryNotify(func() error {
		resp, err := c.Do(req)
		if err != nil {
			return backoff.Permanent(err)
		}
		defer resp.Body.Close()

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return backoff.Permanent(err)
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to get url: '%s', status code: '%d', body: '%s'", url, resp.StatusCode, string(body))
		}

		if string(body) == "Too many requests." {
			return fmt.Errorf("request rate limit reached")
		}
		return nil
	}, backoff.NewExponentialBackOff(), func(err error, td time.Duration) {
		fmt.Printf("[ERROR]: %v\n", err)
		fmt.Printf("[WAIT_IN_SECONDS]: %v\n", 2 * td.Seconds())
		time.Sleep(td)
	}); err != nil {
		*scr <- scrapeResult{Error: err}
		return
	}

	doc, err := utils.GenerateDocument(body)
	if err != nil {
		*scr <- scrapeResult{Error: err}
		return
	}

	res := scrape(int(view), doc)
	if res.Error != nil {
		*scr <- scrapeResult{Error: res.Error}
		return
	}

	fmt.Printf("[SUCCESS]: %v\n", url)
	*scr <- *res
}
