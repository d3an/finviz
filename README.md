# Unofficial Finviz API

## Installation

Run `$ go get github.com/d3an/finviz` to install the package.

## Documentation

`View Types` represent the different views the screener uses to display the data returned.
There are several codes following a 3-digit format:

The first digit is in the range `1-5`.
   1. Overview
   2. Charts
   3. Basic
   4. Tickers
   5. Max Ticker List

The second digit is in the range `1-7`.
   1. Overview Columns
   2. Valuation Columns
   3. Ownership Columns
   4. Performance Columns
   5. Custom Columns
   6. Financial Columns
   7. Technical Columns

The third digit is in the range `0-2`.
   - (0) No Extra Tab
   - (1) Filter Table Tab
   - (2) Custom Settings Tab

The default view is `111`, and only returns the top 20 stocks in the screen.


### Filters

`Filters` are the primary parameters used to modify a screen.
Each filter has an array of default options provided by Finviz.

Initializing a filter with a default value:
```go
exch := ExchangeFilter{NASDAQ}
```
If Finviz supports a filter option that is not currently available in the default list, it can still be used. (Note: Finviz doesn't currently support the `FTSE` value)

Initializing a filter with a custom value:
```go
exch := ExchangeFilter{"FTSE"}
```

#### Filter Implementation

Each filter, by definition, stores its own URL prefix. Examples include `exch_`, `fa_netmargin_`, and `ta_sma200_` respectively for Exchange, Net Profit Margin, and 200-Day Simple Moving Average.
Users only need to register the filter value they want to use, i.e. `NYSE`, `OMUnder80`, `Price10PercentBelowSMA200`.
The default values provided by Finviz have already been defined as constants in `filters.go`.

Descriptive Filters:
- TODO

Fundamental Filters:
- TODO

Technical Filters:
- TODO

### Screen Example

```go
package main

import (
  "fmt"

  . "github.com/d3an/finviz"
)

func main() {
    filters := []FilterInterface{
        ExchangeFilter{AMEX},
        PriceFilter{Price1to10},
    }

    input := ScreenInput{
        TopGainers,
        Ascending,
        Ticker,
        nil,
        filters,
    }

    client := NewClient()

    df, err := RunScreen(client, input)
    if err != nil {
        panic(err)
    }

    // Print screen results dataframe
    fmt.Println(df)
}
```

## ToDo

- [ ] Add CLI tools for running screens
- [ ] Update README.md with more documentation
- [ ] Add support for non-default views (increase the number of tickers a screen can return)
- [ ] Review dataframe package and consider migration for greater support

## Contributing

You can contribute to this project by reporting bugs, suggesting enhancements, or directly by extending and writing features.

## Disclaimer

Using this library to acquire data from Finviz is against their Terms of Service and `robots.txt`.
Use it responsively and at your own risk. This library was built purely for educational purposes.
