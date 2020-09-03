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
          ExchangeFilter(AMEX, NASDAQ),
          AverageVolumeFilter(AvgVolOver50K),
          PriceFilter(PriceOver1),
        },
        View: "performance",
        
    })
    if err != nil {
        panic(err)
    }

    // Print screen results dataframe
    PrintFullDataframe(df)
}
```

### Output

```command line
[20x16] DataFrame

     No.   Ticker   Perf Week Perf Month Perf Quart Perf Half Perf Year Perf YTD Volatility W Volatility M Recom    Avg Volume Rel Volume Price     Change   Volume    
  0: 1     CORT     58.07%    26.64%     40.72%     60.24%    60.24%    64.21%   19.02%       7.65%        2.00     1.03M      85.930000  19.870000 58.45%   88,251,187
  1: 2     GOGO     85.16%    181.31%    204.82%    237.37%   122.54%   48.12%   25.21%       14.11%       2.30     3.88M      24.210000  9.480000  28.63%   94,003,561
  2: 3     NNVC     44.19%    -10.58%    -20.29%    -30.94%   31.91%    122.31%  16.99%       9.37%        -        1.11M      13.580000  5.580000  42.35%   15,050,819
  3: 4     UONE     7.95%     -67.36%    110.50%    105.75%   93.13%    108.32%  19.01%       16.46%       -        3.63M      2.530000   4.210000  22.74%   9,184,204 
  4: 5     BXRX     28.33%    6.65%      3.49%      -46.15%   -         -44.36%  11.67%       9.70%        1.70     600.74K    6.190000   3.850000  26.64%   3,718,628 
  5: 6     SBPH     8.09%     -3.29%     -12.50%    4.26%     -60.27%   -6.96%   11.42%       7.74%        3.00     230.65K    1.680000   1.470000  18.55%   387,454   
  6: 7     FMCI     30.39%    63.76%     85.66%     121.35%   130.25%   124.39%  10.59%       7.46%        -        2.32M      2.000000   22.910000 24.78%   4,644,522 
  7: 8     AMTX     32.56%    82.40%     181.48%    200.20%   186.79%   174.70%  21.96%       23.03%       2.00     1.68M      1.760000   2.280000  11.76%   2,966,951 
  8: 9     CYCN     19.20%    87.58%     32.99%     84.65%    -20.70%   183.09%  10.15%       9.57%        3.00     292.01K    1.650000   7.700000  15.10%   482,880   
  9: 10    OPES     16.38%    13.19%     25.59%     27.64%    34.07%    26.31%   6.65%        5.32%        -        527.92K    1.750000   13.300000 17.39%   923,459   
 10: 11    FTHM     17.86%    130.23%    -          -         -         101.99%  9.97%        12.22%       -        268.19K    0.900000   20.260000 18.76%   240,054   
 11: 12    RAIL     14.57%    -1.70%     34.11%     4.22%     -59.29%   -16.43%  9.06%        8.63%        3.00     196.63K    2.400000   1.730000  15.33%   471,801   
 12: 13    SDC      30.32%    11.49%     15.70%     18.79%    -         12.13%   8.61%        6.70%        2.40     4.86M      7.790000   9.800000  24.52%   37,818,856
 13: 14    BHTG     6.72%     -20.99%    5.93%      -23.53%   -30.62%   -15.88%  7.73%        7.04%        2.00     2.31M      0.100000   1.430000  14.40%   228,777   
 14: 15    VRA      50.95%    68.79%     12.26%     -10.86%   -16.59%   -39.49%  10.58%       7.92%        2.60     306.17K    53.480000  7.140000  31.98%   16,374,591
 15: 16    KRUS     16.50%    31.35%     -13.05%    -11.47%   -43.02%   -44.52%  7.87%        6.83%        1.80     74.45K     2.030000   14.120000 15.08%   151,006   
 16: 17    CPAH     16.88%    2.74%      30.43%     60.71%    305.41%   383.87%  7.33%        6.36%        2.00     83.14K     1.640000   4.500000  13.64%   135,943   
 17: 18    RPTX     19.35%    20.64%     -          -         -         -5.68%   9.85%        6.30%        2.20     218.05K    0.350000   29.050000 15.51%   75,381    
 18: 19    CLPS     -2.36%    51.22%     96.83%     89.80%    -27.20%   -25.60%  13.62%       17.72%       -        168.21K    0.710000   3.720000  11.71%   118,748   
 19: 20    ATNX     38.42%    26.78%     19.75%     14.00%    1.14%     -6.68%   10.14%       6.45%        1.70     549.67K    3.320000   14.250000 13.73%   1,824,149 
     <int> <string> <string>  <string>   <string>   <string>  <string>  <string> <string>     <string>     <string> <string>   <float>    <float>   <string> <string>  

```

## Contributing

You can contribute to this project by reporting bugs, suggesting enhancements, or directly by extending and writing features.

PLEASE submit any issues that you find with the package. This project is still undergoing heavy development and any insight would be incredibly helpful.

## Disclaimer

Using this library to acquire data from Finviz is against their Terms of Service and `robots.txt`.
Use it responsively and at your own risk. This library was built purely for educational purposes.
