# Unofficial Finviz API

## Installation

Run `go get github.com/d3an/finviz` to install the package.

## Documentation

View the [Wiki](https://github.com/d3an/finviz/wiki) for more extensive documentation on this package.

### Screen Example

```go
package main

import (
  . "github.com/d3an/finviz"
)

func main() {
    client := NewClient()
    df, err := RunScreen(client, &ScreenInput{
        Signal: TopGainers,
        GeneralOrder: Descending,
        SpecificOrder: ChangeFromOpen,
    	Filters: []FilterInterface{
        	IndustryFilter(WasteManagement, Airlines),
    	    AverageVolumeFilter(AvgVolOver50K),
            PriceFilter(PriceOver1),
        },
        View: "overview",
    })
    if err != nil {
        panic(err)
    }

    // Print screen results dataframe
    PrintFullDataframe(df)
}
```

## Contributing

You can contribute to this project by reporting bugs, suggesting enhancements, or directly by extending and writing features.
PLEASE submit any issues that you find with the package. This project is still undergoing heavy development and any insight would be incredibly helpful.

## Disclaimer

Using this library to acquire data from Finviz is against their Terms of Service and `robots.txt`.
Use it responsively and at your own risk. This library was built purely for educational purposes.
