# Unofficial Finviz API

## Installation

Run `go get github.com/d3an/finviz` to install the package.

## Documentation

View the [Wiki](https://github.com/d3an/finviz/wiki) for more extensive documentation on this package.

### Screen Example

```
package main

import (
  "fmt"

  . "github.com/d3an/finviz"
)

func main() {
    client := NewClient()

    df, err := RunScreen(client, ScreenInput{
        Signal: TopGainers,
        GeneralOrder: Descending,
        SpecificOrder: ChangeFromOpen,
    	Filters: []FilterInterface{
            IndustryFilter{}.SetMultipleValues(WasteManagement, Airlines),
    	    AverageVolumeFilter{Value: AvgVolOver50K},
            PriceFilter{Value: PriceOver1},
    	},
    })
    if err != nil {
        panic(err)
    }

    // Print screen results dataframe
    fmt.Println(df)
}
```

## ToDo

- [ ] Add CLI tools for running screens
- [x] Update README.md with more documentation
- [x] Add support for multiple of the same filters with `|` operator
- [ ] Add support for non-default views (increase the number of tickers a screen can return)
- [ ] Review dataframe package and consider migration for greater support

## Contributing

You can contribute to this project by reporting bugs, suggesting enhancements, or directly by extending and writing features.

## Disclaimer

Using this library to acquire data from Finviz is against their Terms of Service and `robots.txt`.
Use it responsively and at your own risk. This library was built purely for educational purposes.
