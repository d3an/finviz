// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

import "fmt"

// SignalType represents the various signals to screen with
type SignalType string

// Screener Signals
const (
	AllStocks               SignalType = ""
	TopGainers              SignalType = "ta_topgainers"
	TopLosers               SignalType = "ta_toplosers"
	NewHigh                 SignalType = "ta_newhigh"
	NewLow                  SignalType = "ta_newlow"
	MostVolatile            SignalType = "ta_mostvolatile"
	MostActive              SignalType = "ta_mostactive"
	UnusualVolume           SignalType = "ta_unusualvolume"
	Overbought              SignalType = "ta_overbought"
	Oversold                SignalType = "ta_oversold"
	Downgrades              SignalType = "n_downgrades"
	Upgrades                SignalType = "n_upgrades"
	EarningsBefore          SignalType = "n_earningsbefore"
	EarningsAfter           SignalType = "n_earningsafter"
	RecentInsiderBuying     SignalType = "it_latestbuys"
	RecentInsiderSelling    SignalType = "it_latestsales"
	MajorNews               SignalType = "n_majornews"
	HorizontalSR            SignalType = "ta_p_horizontal"
	TLResistance            SignalType = "ta_p_tlresistance"
	TLSupport               SignalType = "ta_p_tlsupport"
	WedgeUp                 SignalType = "ta_p_wedgeup"
	WedgeDown               SignalType = "ta_p_wedgedown"
	TriangleAscending       SignalType = "ta_p_wedgeresistance"
	TriangleDescending      SignalType = "ta_p_wedgesupport"
	Wedge                   SignalType = "ta_p_wedge"
	ChannelUp               SignalType = "ta_p_channelup"
	ChannelDown             SignalType = "ta_p_channeldown"
	Channel                 SignalType = "ta_p_channel"
	DoubleTop               SignalType = "ta_p_doubletop"
	DoubleBottom            SignalType = "ta_p_doublebottom"
	MultipleTop             SignalType = "ta_p_multipletop"
	MultipleBottom          SignalType = "ta_p_multiplebottom"
	HeadAndShoulders        SignalType = "ta_p_headandshoulders"
	HeadAndShouldersInverse SignalType = "ta_p_headandshouldersinv"
)

// SignalLookup presents a way to find SignalTypes based on their string representation
var SignalLookup = map[string]SignalType{
	"allstocks":                  AllStocks,
	"all stocks":                 AllStocks,
	"topgainers":                 TopGainers,
	"top gainers":                TopGainers,
	"toplosers":                  TopLosers,
	"top losers":                 TopLosers,
	"newhigh":                    NewHigh,
	"new high":                   NewHigh,
	"newlow":                     NewLow,
	"new low":                    NewLow,
	"mostvolatile":               MostVolatile,
	"most volatile":              MostVolatile,
	"mostactive":                 MostActive,
	"most active":                MostActive,
	"unusualvolume":              UnusualVolume,
	"unusual volume":             UnusualVolume,
	"overbought":                 Overbought,
	"over bought":                Overbought,
	"oversold":                   Oversold,
	"over sold":                  Oversold,
	"downgrades":                 Downgrades,
	"downgrade":                  Downgrades,
	"upgrades":                   Upgrades,
	"upgrade":                    Upgrades,
	"earningsbefore":             EarningsBefore,
	"earnings before":            EarningsBefore,
	"earningsafter":              EarningsAfter,
	"earnings after":             EarningsAfter,
	"recentinsiderbuying":        RecentInsiderBuying,
	"recent insider buying":      RecentInsiderBuying,
	"recentinsiderselling":       RecentInsiderSelling,
	"recent insider selling":     RecentInsiderSelling,
	"majornews":                  MajorNews,
	"major news":                 MajorNews,
	"horizontalsr":               HorizontalSR,
	"horizontal sr":              HorizontalSR,
	"horizontal s/r":             HorizontalSR,
	"tlresistance":               TLResistance,
	"trendlineresistance":        TLResistance,
	"tl resistance":              TLResistance,
	"trendline resistance":       TLResistance,
	"tlsupport":                  TLSupport,
	"trendlinesupport":           TLSupport,
	"tl support":                 TLSupport,
	"trendline support":          TLSupport,
	"wedgeup":                    WedgeUp,
	"wedge up":                   WedgeUp,
	"wedgedown":                  WedgeDown,
	"wedge down":                 WedgeDown,
	"triangleascending":          TriangleAscending,
	"triangle ascending":         TriangleAscending,
	"triangledescending":         TriangleDescending,
	"triangle descending":        TriangleDescending,
	"wedge":                      Wedge,
	"channelup":                  ChannelUp,
	"channel up":                 ChannelUp,
	"channeldown":                ChannelDown,
	"channel down":               ChannelDown,
	"channel":                    Channel,
	"doubletop":                  DoubleTop,
	"double top":                 DoubleTop,
	"doublebottom":               DoubleBottom,
	"double bottom":              DoubleBottom,
	"multipletop":                MultipleTop,
	"multiple top":               MultipleTop,
	"multiplebottom":             MultipleBottom,
	"multiple bottom":            MultipleBottom,
	"headandshoulders":           HeadAndShoulders,
	"head&shoulders":             HeadAndShoulders,
	"head and shoulders":         HeadAndShoulders,
	"head & shoulders":           HeadAndShoulders,
	"headandshouldersinverse":    HeadAndShouldersInverse,
	"head&shouldersinverse":      HeadAndShouldersInverse,
	"head and shoulders inverse": HeadAndShouldersInverse,
	"head & shoulders inverse":   HeadAndShouldersInverse,
}

// GetSignal returns the constant of the queried signal
func GetSignal(query string) (SignalType, error) {
	if signal, exists := SignalLookup[query]; exists {
		return signal, nil
	}
	// Add logic that suggests similar matches
	return "", SignalNotFoundError(fmt.Sprintf("Signal not found. Query: \"%v\"", query))
}
