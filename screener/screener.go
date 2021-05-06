package screener

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/cenkalti/backoff/v4"
	"github.com/corpix/uarand"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/go-gota/gota/dataframe"
	"github.com/pkg/errors"

	"github.com/d3an/finviz"
	"github.com/d3an/finviz/utils"
)

const (
	APIURL = "https://finviz.com/screener.ashx"
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

func generateURL(fargs *finvizArgs) (string, error) {
	args := fargs.args
	filters, err := getFiltersValue(args)
	if err != nil {
		return "", err
	}

	var rowParam string
	if row, ok := args["row"]; ok {
		rowParam = fmt.Sprintf("&r=%d", row.(int))
	}

	switch fargs.view {
	case "overview":
		return fmt.Sprintf("%s?v=110%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), rowParam), nil
	case "valuation":
		return fmt.Sprintf("%v?v=120%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), rowParam), nil
	case "ownership":
		return fmt.Sprintf("%v?v=130%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), rowParam), nil
	case "performance":
		return fmt.Sprintf("%v?v=140%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), rowParam), nil
	case "custom":
		customColumns, err := getCustomColumnsValue(args)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v?v=150%v%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), customColumns, rowParam), nil
	case "financial":
		return fmt.Sprintf("%v?v=160%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), rowParam), nil
	case "technical":
		return fmt.Sprintf("%v?v=170%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), rowParam), nil
	case "charts":
		chartStyle, err := getChartStylingValue(args)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v?v=210%v%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), chartStyle, rowParam), nil
	case "basic":
		chartStyle, err := getChartStylingValue(args)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v?v=310%v%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), chartStyle, rowParam), nil
	case "news":
		chartStyle, err := getChartStylingValue(args)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v?v=320%v%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), chartStyle, rowParam), nil
	case "description":
		chartStyle, err := getChartStylingValue(args)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v?v=330%v%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), chartStyle, rowParam), nil
	case "snapshot":
		chartStyle, err := getChartStylingValue(args)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v?v=340%v%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), chartStyle, rowParam), nil
	case "ta":
		chartStyle, err := getChartStylingValue(args)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v?v=350%v%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), chartStyle, rowParam), nil
	case "tickers":
		return fmt.Sprintf("%v?v=410%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), rowParam), nil
	case "bulk":
		return fmt.Sprintf("%v?v=510%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), rowParam), nil
	case "bulkfull":
		return fmt.Sprintf("%v?v=520%v%v%v%v%v", APIURL, getSignalValue(args),
			filters, getTickersValue(args), getOrderValue(args), rowParam), nil
	default:
		return "", fmt.Errorf("error view '%s' not found", fargs.view)
	}
}

func scrape(view string, doc *goquery.Document) *scrapeResult {
	switch view {
	case "overview", "valuation", "ownership", "performance", "custom", "financial", "technical":
		return DefaultScrape(doc)
	case "charts":
		return ChartScrape(doc)
	case "basic":
		return BasicScrape(doc)
	case "news":
		return NewsScrape(doc)
	case "description":
		return DescriptionScrape(doc)
	case "snapshot":
		return SnapshotScrape(doc)
	case "ta":
		return TAScrape(doc)
	case "tickers":
		return TickerScrape(doc)
	case "bulk":
		return BulkScrape(doc)
	case "bulkfull":
		return BulkFullScrape(doc)
	default:
		return &scrapeResult{Error: fmt.Errorf("error view '%s' not found", view)}
	}
}

type finvizArgs struct {
	mu   sync.Mutex
	view string
	args map[string]interface{}
}

func (f *finvizArgs) SetRow(maxRows, index int) {
	f.mu.Lock()
	f.args["row"] = (maxRows * (index + 1)) + 1
}

func (c *Client) GetScreenerResults(view string, args map[string]interface{}) (*dataframe.DataFrame, error) {
	var wg sync.WaitGroup
	var results []map[string]interface{}

	fargs := &finvizArgs{view: view, args: args}

	firstLook := make(chan scrapeResult, 1)
	wg.Add(1)

	fargs.mu.Lock()
	go c.getData(fargs, &wg, &firstLook)
	wg.Wait()

	firstPage := <-firstLook
	if firstPage.Error != nil {
		return nil, errors.Wrapf(firstPage.Error, "error received while scraping screener '%s'", view)
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
		fargs.SetRow(maxRows, i)
		go c.getData(fargs, &wg, &scrapeResults[i])
	}
	wg.Wait()

	for i := 0; i < pagesLeft; i++ {
		page := <-scrapeResults[i]
		if page.Error != nil {
			return nil, errors.Wrapf(page.Error, "error received while scraping screener '%s', page '%d'", view, i+2)
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

func (c *Client) getData(fargs *finvizArgs, wg *sync.WaitGroup, scr *chan scrapeResult) {
	defer wg.Done()
	defer close(*scr)

	url, err := generateURL(fargs)
	fargs.mu.Unlock()
	if err != nil {
		*scr <- scrapeResult{Error: err}
		return
	}
	fmt.Println(url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
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
		fmt.Printf("[WAIT_IN_SECONDS]: %v\n", td.Seconds())
	}); err != nil {
		*scr <- scrapeResult{Error: err}
		return
	}

	doc, err := finviz.GenerateDocument(body)
	if err != nil {
		*scr <- scrapeResult{Error: err}
		return
	}

	res := scrape(fargs.view, doc)
	if res.Error != nil {
		*scr <- scrapeResult{Error: res.Error}
		return
	}

	*scr <- *res
}

/*
func (c *Client) GetData(view string, args *map[string]interface{}) (*dataframe.DataFrame, error) {
	url, err := GenerateURL(view, args)
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

	doc, err := finviz.GenerateDocument(body)
	if err != nil {
		return nil, err
	}

	results, err := Scrape(view, doc)
	if err != nil {
		return nil, err
	}

	df := dataframe.LoadRecords(results)
	return CleanScreenerDataFrame(&df), nil
}

func main() {
	client := New(&Config{userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"})

	df, err := client.getData("overview", &map[string]interface{}{
		"signal":         TopGainers,
		"general_order":  Descending,
		"specific_order": ChangeFromOpen,
		"filters": []FilterInterface{
			ExchangeFilter(NYSE, NASDAQ),
			AverageVolumeFilter(AvgVolOver50K),
			PriceFilter(PriceOver1),
		}})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(df)
}
*/
