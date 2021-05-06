package screener

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func DefaultScrape(doc *goquery.Document) *scrapeResult {
	scr := &scrapeResult{
		PageCount: doc.Find("#pageSelect").Children().Size(),
	}

	doc.Find("#screener-content").Find("tbody").Eq(3).Each(func(i int, childNode *goquery.Selection) {
		childNode.Children().Each(func(j int, childNode *goquery.Selection) {
			if j == 0 {
				childNode.Children().Each(func(k int, childNode *goquery.Selection) {
					scr.Keys = append(scr.Keys, childNode.Text())
				})
				return
			}
			result := make(map[string]interface{})
			childNode.Children().Each(func(k int, childNode *goquery.Selection) {
				result[scr.Keys[k]] = childNode.Text()
			})
			scr.Results = append(scr.Results, result)
		})
	})

	return scr
}

func ChartScrape(doc *goquery.Document) *scrapeResult {
	scr := &scrapeResult{
		Keys:      []string{"Ticker", "Chart", "Company", "Industry", "Country", "Market Cap"},
		PageCount: doc.Find("#pageSelect").Children().Size(),
	}

	doc.Find("#screener-content").Find("tbody").Children().Eq(4).Find("tbody").Find("td").Children().Each(func(i int, spanNode *goquery.Selection) {
		var rawTickerData = make(map[string]interface{})
		if titleAttr := spanNode.AttrOr("title", ""); titleAttr != "" {
			body := strings.Split(strings.Split(strings.Split(titleAttr, "body=[")[2], "]")[0], "<b>")[1]
			rawTickerData["Company"] = strings.Split(body, "</b>")[0]
			body = strings.Split(body, "<br>")[1]
			rawTickerData["Industry"] = strings.Split(body, " | ")[0]
			rawTickerData["Country"] = strings.Split(body, " | ")[1]
			rawTickerData["Market Cap"] = strings.Split(body, " | ")[2]
		}
		rawTickerData["Ticker"] = strings.Split(strings.Split(spanNode.Find("a").AttrOr("href", ""), "quote.ashx?t=")[1], "&")[0]
		rawTickerData["Chart"] = spanNode.Find("img").AttrOr("src", "")
		scr.Results = append(scr.Results, rawTickerData)
	})

	return scr
}

func BasicScrape(doc *goquery.Document) *scrapeResult {
	scr := &scrapeResult{
		Keys:      []string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range"},
		PageCount: doc.Find("#pageSelect").Children().Size(),
	}

	rootNode := doc.Find("div > table > tbody").Children()
	firstDataNodeIndex := 4
	lastDataNodeIndex := rootNode.Size() - 3

	rootNode.Slice(firstDataNodeIndex, lastDataNodeIndex).Each(func(i int, childNode *goquery.Selection) {
		var rawTickerData = make(map[string]interface{})
		_, rawTickerData = basicDefaultViewHelper(childNode, []string{}, rawTickerData)
		scr.Results = append(scr.Results, rawTickerData)
	})

	return scr
}

func NewsScrape(doc *goquery.Document) *scrapeResult {
	scr := &scrapeResult{
		Keys:      []string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "News"},
		PageCount: doc.Find("#pageSelect").Children().Size(),
	}

	var headers []string
	var rawTickerData map[string]interface{}

	rootNode := doc.Find("div > table > tbody").Children()
	firstDataNodeIndex := 4
	lastDataNodeIndex := rootNode.Size() - 3

	rootNode.Slice(firstDataNodeIndex, lastDataNodeIndex).Each(func(i int, childNode *goquery.Selection) {
		if i%3 == 0 {
			rawTickerData = make(map[string]interface{})
			headers, rawTickerData = basicDefaultViewHelper(childNode, headers, rawTickerData)
		} else if i%3 == 1 {
			headers, rawTickerData = basicNewsViewHelper(childNode, headers, rawTickerData)
			scr.Results = append(scr.Results, rawTickerData)
		}
	})

	return scr
}

func DescriptionScrape(doc *goquery.Document) *scrapeResult {
	scr := &scrapeResult{
		Keys:      []string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "Description"},
		PageCount: doc.Find("#pageSelect").Children().Size(),
	}
	var headers []string
	var rawTickerData map[string]interface{}

	rootNode := doc.Find("div > table > tbody").Children()
	firstDataNodeIndex := 4
	lastDataNodeIndex := rootNode.Size() - 3

	rootNode.Slice(firstDataNodeIndex, lastDataNodeIndex).Each(func(i int, childNode *goquery.Selection) {
		if i%3 == 0 {
			rawTickerData = make(map[string]interface{})
			headers, rawTickerData = basicDefaultViewHelper(childNode, headers, rawTickerData)
		} else if i%3 == 1 {
			headers, rawTickerData = basicDescriptionViewHelper(childNode, headers, rawTickerData)
			scr.Results = append(scr.Results, rawTickerData)
		}
	})

	return scr
}

func SnapshotScrape(doc *goquery.Document) *scrapeResult {
	scr := &scrapeResult{
		Keys:      []string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "EPS (ttm)", "P/E", "EPS this Y", "Forward P/E", "EPS next Y", "PEG", "EPS past 5Y", "P/S", "EPS next 5Y", "P/B", "EPS Q/Q", "Dividend", "Sales Q/Q", "Insider Own", "Inst Own", "Insider Trans", "Inst Trans", "Short Float", "Earnings", "Analyst Recom", "Target Price", "Avg Volume", "52W Range", "News", "Description", "Insider Trading"},
		PageCount: doc.Find("#pageSelect").Children().Size(),
	}

	var headers []string
	var rawTickerData map[string]interface{}

	rootNode := doc.Find("div > table > tbody").Children()
	firstDataNodeIndex := 4
	lastDataNodeIndex := rootNode.Size() - 3

	rootNode.Slice(firstDataNodeIndex, lastDataNodeIndex).Each(func(i int, childNode *goquery.Selection) {
		if i%3 == 0 {
			rawTickerData = make(map[string]interface{})
			headers, rawTickerData = basicDefaultViewHelper(childNode, headers, rawTickerData)
		} else if i%3 == 1 {
			snapshotNodeLen := childNode.Find("td > table > tbody").Size()
			for j := 0; j < snapshotNodeLen; j++ {
				switch j {
				case 0:
					headers, rawTickerData = basicNewsViewHelper(childNode, headers, rawTickerData)
				case 1:
					headers, rawTickerData = basicDescriptionViewHelper(childNode, headers, rawTickerData)
				case 2:
					headers, rawTickerData = basicInsiderTradingViewHelper(childNode, headers, rawTickerData)
				}
			}
			scr.Results = append(scr.Results, rawTickerData)
		}
	})

	return scr
}

func TAScrape(doc *goquery.Document) *scrapeResult {
	scr := &scrapeResult{
		Keys:      []string{"Ticker", "Company", "Country", "Industry", "Chart", "Market Cap", "Perf Week", "Beta", "Perf Month", "ATR", "Perf Quarter", "Volatility W", "Perf Half Y", "Volatility M", "Perf Year", "SMA20", "Perf YTD", "SMA50", "RSI (14)", "SMA200", "Change Open", "52W High", "Gap", "52W Low", "Rel Volume", "Short Float", "Avg Volume", "Candlestick", "52W Range"},
		PageCount: doc.Find("#pageSelect").Children().Size(),
	}

	var headers []string

	rootNode := doc.Find("div > table > tbody").Children()
	firstDataNodeIndex := 4
	lastDataNodeIndex := rootNode.Size() - 3

	rootNode.Slice(firstDataNodeIndex, lastDataNodeIndex).Each(func(i int, childNode *goquery.Selection) {
		var rawTickerData = make(map[string]interface{})
		headers, rawTickerData = basicDefaultViewHelper(childNode, headers, rawTickerData)
		scr.Results = append(scr.Results, rawTickerData)
	})

	return scr
}

func TickerScrape(doc *goquery.Document) *scrapeResult {
	scr := &scrapeResult{
		Keys:      []string{"Ticker", "Change", "Chart", "Company", "Industry", "Country", "Market Cap"},
		PageCount: doc.Find("#pageSelect").Children().Size(),
	}
	var rawTickerData map[string]interface{}

	doc.Find("#screener-content").Find("tbody").Eq(3).Find("tr").Children().Children().Each(func(i int, tickerNode *goquery.Selection) {
		rawTickerData = make(map[string]interface{})
		if titleAttr := tickerNode.AttrOr("title", ""); titleAttr != "" {
			body := strings.Split(strings.Split(titleAttr, "body=[")[2], "]")[0]
			rawTickerData["Chart"] = strings.Split(strings.Split(body, "<img src='")[1], "'>")[0]
			body = strings.Split(strings.Split(body, "<img src='")[1], "'>")[1]
			rawTickerData["Company"] = strings.Split(strings.Split(body, "<b>")[1], "</b>")[0]
			body = strings.Split(strings.Split(strings.Split(body, "<b>")[1], "</b>")[1], "<br>\u00a0")[1]
			rawTickerData["Industry"] = strings.Split(body, " | ")[0]
			rawTickerData["Country"] = strings.Split(body, " | ")[1]
			rawTickerData["Market Cap"] = strings.Split(body, " | ")[2]
			rawTickerData["Change"] = strings.Split(strings.Split(body, " | ")[3], "Change: ")[1]
		}
		rawTickerData["Ticker"] = strings.TrimSpace(tickerNode.Text())
		scr.Results = append(scr.Results, rawTickerData)
	})

	return scr
}

func BulkScrape(doc *goquery.Document) *scrapeResult {
	scr := &scrapeResult{
		Keys:      []string{"Ticker", "Change", "Chart", "Company", "Industry", "Country", "Market Cap"},
		PageCount: doc.Find("#pageSelect").Children().Size(),
	}
	var tickerDataSlice [][]map[string]interface{}
	var rawTickerData map[string]interface{}
	rootNode := doc.Find("#screener-content").Find("tbody").Eq(3).Find("tr")

	rootNode.Children().Each(func(i int, columnNode *goquery.Selection) {
		var columnDataSlice []map[string]interface{}
		columnNode.Children().Find("table > tbody").Children().Each(func(j int, rowNode *goquery.Selection) {
			rawTickerData = make(map[string]interface{})
			rowNode.Children().Each(func(k int, itemNode *goquery.Selection) {
				if k == 0 {
					if titleAttr := itemNode.AttrOr("title", ""); titleAttr != "" {
						body := strings.Split(strings.Split(titleAttr, "body=[")[2], "]")[0]
						rawTickerData["Chart"] = strings.Split(strings.Split(body, "<img src='")[1], "'>")[0]
						body = strings.Split(strings.Split(body, "<img src='")[1], "'>")[1]
						rawTickerData["Company"] = strings.Split(strings.Split(body, "<b>")[1], "</b>")[0]
						body = strings.Split(strings.Split(strings.Split(body, "<b>")[1], "</b>")[1], "<br>\u00a0")[1]
						rawTickerData["Industry"] = strings.Split(body, " | ")[0]
						rawTickerData["Country"] = strings.Split(body, " | ")[1]
						rawTickerData["Market Cap"] = strings.Split(body, " | ")[2]
					}
					rawTickerData["Ticker"] = itemNode.Text()
				} else if k == 1 {
					rawTickerData["Change"] = strings.TrimSpace(itemNode.Text())
				}
			})
			columnDataSlice = append(columnDataSlice, rawTickerData)
		})
		tickerDataSlice = append(tickerDataSlice, columnDataSlice)
	})

	columnCount := len(tickerDataSlice)
	rowCount := len(tickerDataSlice[0])
	for j := 0; j < rowCount; j++ {
		for i := 0; i < columnCount; i++ {
			if j <= len(tickerDataSlice[i])-1 {
				scr.Results = append(scr.Results, tickerDataSlice[i][j])
			}
		}
	}

	return scr
}

func BulkFullScrape(doc *goquery.Document) *scrapeResult {
	scr := &scrapeResult{
		Keys:      []string{"Ticker", "Change", "Price", "Volume", "Relative Volume", "Chart", "Company", "Industry", "Country", "Market Cap"},
		PageCount: doc.Find("#pageSelect").Children().Size(),
	}
	var tickerDataSlice [][]map[string]interface{}
	var rawTickerData map[string]interface{}
	rootNode := doc.Find("#screener-content").Find("tbody").Eq(3).Find("tr")

	rootNode.Children().Each(func(i int, columnNode *goquery.Selection) {
		var columnDataSlice []map[string]interface{}
		columnNode.Children().Find("table > tbody").Children().Each(func(j int, rowNode *goquery.Selection) {
			rawTickerData = make(map[string]interface{})
			rowNode.Children().Each(func(k int, itemNode *goquery.Selection) {
				switch k {
				case 0:
					if titleAttr := itemNode.AttrOr("title", ""); titleAttr != "" {
						body := strings.Split(strings.Split(titleAttr, "body=[")[2], "]")[0]
						rawTickerData["Chart"] = strings.Split(strings.Split(body, "<img src='")[1], "'>")[0]
						body = strings.Split(strings.Split(body, "<img src='")[1], "'>")[1]
						rawTickerData["Company"] = strings.Split(strings.Split(body, "<b>")[1], "</b>")[0]
						body = strings.Split(strings.Split(strings.Split(body, "<b>")[1], "</b>")[1], "<br>\u00a0")[1]
						rawTickerData["Industry"] = strings.Split(body, " | ")[0]
						rawTickerData["Country"] = strings.Split(body, " | ")[1]
						rawTickerData["Market Cap"] = strings.Split(body, " | ")[2]
					}
					rawTickerData["Ticker"] = itemNode.Text()
				case 1:
					rawTickerData["Price"] = strings.TrimSpace(itemNode.Text())
				case 2:
					rawTickerData["Change"] = itemNode.Children().Eq(0).Text()
				case 3:
					rawTickerData["Volume"] = strings.TrimSpace(itemNode.Text())
				case 4:
					rawTickerData["Relative Volume"] = itemNode.Find("small").Text()
				}
			})
			columnDataSlice = append(columnDataSlice, rawTickerData)
		})
		tickerDataSlice = append(tickerDataSlice, columnDataSlice)
	})

	columnCount := len(tickerDataSlice)
	rowCount := len(tickerDataSlice[0])
	for j := 0; j < rowCount; j++ {
		for i := 0; i < columnCount; i++ {
			if j <= len(tickerDataSlice[i])-1 {
				scr.Results = append(scr.Results, tickerDataSlice[i][j])
			}
		}
	}

	return scr
}
