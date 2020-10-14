# Go Wrapper for Unofficial FinViz API

**WARNING:** This package is undergoing heavy development. Breaking changes may occur in upcoming releases.

![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-52%25-brightgreen.svg?longCache=true&style=flat)

[FinViz](https://finviz.com/?a=128493348) aims to provide financial visualizations and data analysis tools.
The site is well-known for its Screener, but also has other apps such as News, Economics Calendar, Maps, Groups,
Portfolio, Insider Trading, Futures, Forex, Crypto, and more services with a [FINVIZ*Elite](https://finviz.com/elite.ashx?a=128493348) subscription.

## Installation

Run `go get github.com/d3an/finviz` to install the package.

## Documentation

View the [Wiki](https://github.com/d3an/finviz/wiki) for more extensive documentation.

### Screen Example

```go
package main

import (
  . "github.com/d3an/finviz"
  . "github.com/d3an/finviz/screener"
)

func main() {
    df, err := GetScreenerData(NewClient(), &PerformanceScreenerView{}, &map[string]interface{}{
        "signal": TopGainers,
        "general_order": Descending,
        "specific_order": ChangeFromOpen,
        "filters": []FilterInterface{
          ExchangeFilter(NYSE, NASDAQ),
          AverageVolumeFilter(AvgVolOver50K),
          PriceFilter(PriceOver1),
        },
    })
    if err != nil {
        panic(err)
    }

    // Print screen results dataframe
    PrintFullDataFrame(df)
}
```

### Output

```command line
[20x16] DataFrame

     No.   Ticker   Perf Week Perf Month Perf Quart Perf Half Perf Year Perf YTD  Volatility W Volatility M Recom    Avg Volume Rel Volume Price     Change   Volume  
  0: 1     AEMD     0.503800  0.271000   0.064900   0.296100  -0.429000 1.045700  0.165500     0.086900     2.000000 334390     43.070000  1.970000  0.331100 14402780
  1: 2     IH       NaN       NaN        NaN        NaN       NaN       0.696900  NaN          NaN          NaN      3730000    0.810000   27.150000 0.209900 3017744
  2: 3     MGI      0.214500  0.287600   0.192000   1.789900  -0.056400 0.833300  0.089000     0.065300     3.300000 1500000    5.100000   3.850000  0.203100 7629931
  3: 4     LNSR     -0.200900 NaN        NaN        NaN       NaN       -0.032200 0.157300     NaN          NaN      315500     0.130000   8.710000  0.141500 40537  
  4: 5     SOS      0.430900  0.494400   -0.120900  1.423200  -0.132500 -0.271000 0.129700     0.118800     5.000000 86210      121.570000 2.690000  0.274900 10481293
  5: 6     FTHM     0.316200  0.402200   NaN        NaN       NaN       1.373900  0.141400     0.099100     2.000000 190640     4.370000   23.810000 0.181600 833313  
  6: 7     CLVS     0.141600  -0.084300  -0.007700  -0.111000 1.054500  -0.385100 0.070100     0.080200     3.300000 4670000    2.510000   6.410000  0.157000 11715925
  7: 8     BWMX     0.200900  -0.073500  1.451400   1.141400  1.260000  1.146800  0.108000     0.068000     1.000000 89950      1.600000   21.940000 0.128000 143491  
  8: 9     BVXV     0.057300  0.084000   -0.086500  3.974400  5.925300  3.390700  0.095600     0.070500     NaN      90490      2.950000   40.790000 0.199000 266682  
  9: 10    COGT     0.017600  0.070400   0.003500   6.706700  0.952700  3.012800  0.164500     0.118200     1.000000 669750     0.820000   2.890000  0.170000 551267  
 10: 11    EVOK     0.106600  -0.033700  0.370800   2.935500  5.113800  2.012300  0.080700     0.060700     2.000000 398170     1.450000   4.880000  0.127000 578439  
 11: 12    SOL      0.443200  1.396200   1.976600   2.735300  1.609600  1.692600  0.243800     0.192600     2.500000 1180000    4.650000   3.810000  0.123900 5497191
 12: 13    INSE     0.352600  0.278800   0.244800   1.048500  -0.402300 -0.374800 0.095000     0.070200     2.000000 92540      2.020000   4.220000  0.153000 187056  
 13: 14    CBAT     -0.155400 2.622200   2.462200   7.150000  4.826600  1.834800  0.149200     0.349400     NaN      8490000    0.480000   3.260000  0.105100 4097024
 14: 15    IEA      0.363200  0.700200   0.982100   2.614000  1.988500  1.413000  0.130400     0.131600     2.000000 459600     2.410000   7.770000  0.134300 1106795
 15: 16    ENLV     -0.043400 1.662200   1.617000   1.476800  0.710200  0.652900  0.114000     0.109700     NaN      1000000    1.790000   13.870000 0.227400 1792726
 16: 17    MFH      0.409500  0.292600   0.042300   1.767900  0.436900  0.804900  0.249600     0.118000     NaN      365300     6.350000   2.960000  0.184000 2320141
 17: 18    ALRN     0.544700  0.520000   0.979200   2.022600  3.418600  2.315900  0.194100     0.119100     1.400000 829030     3.320000   1.900000  0.151500 2754670
 18: 19    AFI      0.122000  0.084600   0.212000   1.820000  -0.330700 -0.009400 0.113400     0.093700     3.000000 224070     6.950000   4.230000  0.113200 1557514
 19: 20    MMLP     0.322800  0.235300   -0.147200  0.513500  -0.578900 -0.583100 0.179100     0.125800     3.300000 332750     1.320000   1.680000  0.070100 439127  
     <int> <string> <float>   <float>    <float>    <float>   <float>   <float>   <float>      <float>      <float>  <int>      <float>    <float>   <float>  <int>  

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
