package finviz

import (
	"fmt"

	"github.com/go-gota/gota/dataframe"

	"net/http"
	"strings"
)

// APIURL is the base URL for the screener. Elite not supported yet.
const APIURL = "https://finviz.com/screener.ashx"

// ScreenInput represents the data passed to the screen
type ScreenInput struct {
	signal        SignalType
	generalOrder  GeneralOrderType
	specificOrder SpecificOrderType
	tickers       []string
	filters       []FilterInterface
	// view			ViewType
	// column		ColumnType
	// tab			TabType
	// settings		[]SettingsType
}

func getScreenerView(viewType ViewType, columnType ColumnType, tabType TabType) string {
	if viewType == "" {
		viewType = "1"
	}
	if columnType == "" {
		columnType = "1"
	}
	if tabType == "" {
		tabType = "1"
	}

	return fmt.Sprintf("?v=%v%v%v", viewType, columnType, tabType)
}

func getFilterList(filters []FilterInterface) string {
	filterSize := len(filters)
	if filterSize == 0 {
		return ""
	}

	var filterKeys []string
	for i := 0; i < filterSize; i++ {
		filterKeys = append(filterKeys, filters[i].GetURLKey())
	}

	filterList := strings.Join(filterKeys, ",")
	return fmt.Sprintf("&f=%v", filterList)
}

func getSignal(signal SignalType) string {
	if signal == "" {
		return ""
	}
	return fmt.Sprintf("&s=%v", signal)
}

func getSortOrder(generalOrder GeneralOrderType, signal SignalType, specificOrder SpecificOrderType) string {
	if specificOrder == Signal && signal == "" && generalOrder == Ascending {
		return ""
	} else if specificOrder == Signal && signal == "" {
		return fmt.Sprintf("&o=%v", generalOrder)
	}

	if specificOrder == "" && generalOrder == "" {
		return ""
	}

	return fmt.Sprintf("&o=%v%v", generalOrder, specificOrder)
}

func getTickerList(tickers []string) string {
	tickersSize := len(tickers)
	if tickersSize == 0 {
		return ""
	}

	for i := 0; i < tickersSize; i++ {
		tickers[i] = strings.ToUpper(tickers[i])
	}

	tickerList := strings.Join(tickers, ",")
	return fmt.Sprintf("&t=%v", tickerList)
}

// GenerateURL consumes valid inputs to the screen and generates a corresponding valid URL
func GenerateURL(input ScreenInput) string {
	screenerView := getScreenerView("", "", "")
	signal := getSignal(input.signal)
	filterList := getFilterList(input.filters)
	sortOrder := getSortOrder(input.generalOrder, input.signal, input.specificOrder)
	tickerList := getTickerList(input.tickers)

	return fmt.Sprintf("%v?%v%v%v%v%v", APIURL, screenerView, signal, filterList, tickerList, sortOrder)
}

// RunScreen consumes a client and screen input to produce a dataframe of results
func RunScreen(c *http.Client, input ScreenInput) (*dataframe.DataFrame, error) {
	url := GenerateURL(input)

	html, err := MakeGetRequest(c, url)
	if err != nil {
		return nil, err
	}

	df, err := GetStockDataframe(html)
	if err != nil {
		return nil, err
	}

	return df, nil
}
