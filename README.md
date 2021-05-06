# Unofficial Go API & CLI for FinViz

**WARNING:** This package is undergoing heavy development. Breaking changes may occur in upcoming releases.

![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-52%25-brightgreen.svg?longCache=true&style=flat)

[FinViz](https://finviz.com/?a=128493348) aims to provide financial visualizations and data analysis tools.
The site is well-known for its Screener, but also has other apps such as News, Economics Calendar, Maps, Groups,
Portfolio, Insider Trading, Futures, Forex, Crypto, and more services with a [FINVIZ*Elite](https://finviz.com/elite.ashx?a=128493348) subscription.

## Installation

Run `go get github.com/d3an/finviz` to install the package.

Run `go install github.com/d3an/finviz/finviz` to install the CLI.

## Documentation

View the [Wiki](https://github.com/d3an/finviz/wiki) for more extensive documentation.

### Screen Example

```go
package main

import (
  "github.com/d3an/finviz"
  "github.com/d3an/finviz/screener"
)

func main() {
    client := screener.New(nil)

    df, err := client.GetScreenerResults("https://finviz.com/screener.ashx?v=110&s=ta_unusualvolume&f=exch_nyse,cap_largeunder&o=-volume")
    if err != nil {
        panic(err)
    }

    finviz.PrintFullDataFrame(df)
}
```

### Output

```command line
[31x11] DataFrame

     No.   Ticker   Company                                         Sector                 Industry                       Country        Market Cap  P/E       Price      Change    Volume
  0: 1     NYT      The New York Times Company                      Communication Services Publishing                     USA            7290000000  72.430000 43.170000  -0.037500 9448297
  1: 2     KAR      KAR Auction Services, Inc.                      Consumer Cyclical      Specialty Retail               USA            2130000000  NaN       16.830000  0.159900  8442565
  2: 3     HAYW     Hayward Holdings, Inc.                          Industrials            Electrical Equipment & Parts   USA            5580000000  NaN       24.110000  0.243400  5881887
  3: 4     EQC      Equity Commonwealth                             Real Estate            REIT - Office                  USA            3350000000  7.890000  27.660000  -0.044600 5447721
  4: 5     MNR      Monmouth Real Estate Investment Corporation     Real Estate            REIT - Industrial              USA            1900000000  NaN       19.350000  0.060300  4936629
  5: 6     HLF      Herbalife Nutrition Ltd.                        Consumer Defensive     Packaged Foods                 USA            5630000000  17.070000 47.480000  0.057900  4708540
  6: 7     ABC      AmerisourceBergen Corporation                   Healthcare             Medical Distribution           USA            24410000000 NaN       118.820000 -0.055900 4029406
  7: 8     TUP      Tupperware Brands Corporation                   Consumer Cyclical      Packaging & Containers         USA            1310000000  11.680000 25.510000  0.062000  3063523
  8: 9     LL       Lumber Liquidators Holdings, Inc.               Consumer Cyclical      Home Improvement Retail        USA            690160000   10.760000 23.090000  -0.067400 2956701
  9: 10    RYAM     Rayonier Advanced Materials Inc.                Basic Materials        Chemicals                      USA            456170000   NaN       6.990000   -0.232700 2530357
 10: 11    ALC      Alcon Inc.                                      Healthcare             Medical Instruments & Supplies Switzerland    35000000000 NaN       70.040000  -0.060200 2442493
 11: 12    PHG      Koninklijke Philips N.V.                        Healthcare             Diagnostics & Research         Netherlands    52720000000 39.640000 57.560000  0.021700  1903034
 12: 13    DLB      Dolby Laboratories, Inc.                        Communication Services Entertainment                  USA            9910000000  31.470000 96.980000  -0.063300 1737271
 13: 14    GMED     Globus Medical, Inc.                            Healthcare             Medical Devices                USA            7410000000  73.150000 73.370000  0.053600  1409087
 14: 15    CVII     Churchill Capital Corp VII                      Financial              Shell Companies                USA            1370000000  NaN       9.940000   0.001000  1404171
 15: 16    CBT      Cabot Corporation                               Basic Materials        Specialty Chemicals            USA            3550000000  NaN       60.910000  0.061100  1398203
 16: 17    PQG      PQ Group Holdings Inc.                          Basic Materials        Specialty Chemicals            USA            2000000000  NaN       14.590000  0.003400  978600
 17: 18    ASPN     Aspen Aerogels, Inc.                            Industrials            Building Products & Equipment  USA            611760000   NaN       19.520000  0.094800  922643
 18: 19    INSP     Inspire Medical Systems, Inc.                   Healthcare             Medical Devices                USA            5610000000  NaN       197.880000 -0.126500 666032
 19: 20    FDP      Fresh Del Monte Produce Inc.                    Consumer Defensive     Farm Products                  Cayman Islands 1660000000  33.560000 34.630000  0.199900  510053
 20: 21    SYX      Systemax Inc.                                   Industrials            Industrial Distribution        USA            1290000000  19.610000 33.380000  -0.218100 405710
 21: 22    TREC     Trecora Resources                               Basic Materials        Specialty Chemicals            USA            187540000   38.720000 7.550000   -0.009200 336389
 22: 23    MSD      Morgan Stanley Emerging Markets Debt Fund, Inc. Financial              Closed-End Fund - Debt         USA            185110000   20.310000 9.080000   NaN       296496
 23: 24    LUB      Luby's, Inc.                                    Consumer Cyclical      Restaurants                    USA            113900000   NaN       3.710000   0.002700  202157
 24: 25    BOAS     BOA Acquisition Corp.                           Financial              Shell Companies                USA            223560000   NaN       9.720000   -0.002100 153135
 25: 26    SRL      Scully Royalty Ltd.                             Financial              Capital Markets                Hong Kong      170750000   NaN       14.100000  0.072200  117014
 26: 27    JEQ      Aberdeen Japan Equity Fund, Inc.                Financial              Closed-End Fund - Foreign      USA            118670000   8.120000  8.850000   0.002300  68347  
 27: 28    KFS      Kingsway Financial Services Inc.                Consumer Cyclical      Auto & Truck Dealerships       USA            119430000   NaN       5.050000   0.004000  61383  
 28: 29    SPGS     Simon Property Group Acquisition Holdings, Inc. Financial              Shell Companies                USA            338790000   NaN       9.820000   -0.008100 55895  
 29: 30    EGF      BlackRock Enhanced Government Fund, Inc.        Financial              Closed-End Fund - Debt         USA            55400000    41.530000 13.000000  NaN       31232  
 30: 31    NBO      Neuberger Berman New York Municipal Fund, Inc.  Financial              Asset Management               USA            64690000    94.370000 12.740000  0.010300  26563  
     <int> <string> <string>                                        <string>               <string>                       <string>       <int>       <float>   <float>    <float>   <int>  
```

## Testing

To generate the coverage badge, run the following:

```command line
$ go test -covermode=count -coverprofile=count.out
$ gopherbadger -manualcov=$(go tool cover -func=count.out | tail -c 6 | head -c 4) -png=false
```

For now, you might need to manually edit the coverage badge in the README.md.

## Contributing

You can contribute to this project by reporting bugs, suggesting enhancements, or directly by extending and writing features.

PLEASE submit any issues that you find with the package. This project is still undergoing heavy development and any insight would be incredibly helpful.

## Disclaimer

Using this library to acquire data from FinViz is against their Terms of Service and `robots.txt`.
Use it responsibly and at your own risk. This library was built purely for educational purposes.
