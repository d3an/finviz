package finviz

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
