package finviz

import (
	"fmt"
)

// FilterInterface is used to query filters
type FilterInterface interface {
	GetURLKey() string
	SetMultipleValues(options ...interface{}) FilterInterface
}

/***************************************************************************
*****	DESCRIPTIVE   ******************************************************
***************************************************************************/

// ExchangeType represents the types of exchange filters
type ExchangeType string

// Default filter constants
const (
	AMEX   ExchangeType = "amex"
	NASDAQ ExchangeType = "nasd"
	NYSE   ExchangeType = "nyse"
)

// ExchangeFilter is a filter type
type ExchangeFilter struct {
	Value ExchangeType
}

// GetURLKey returns the filter's url key
func (e ExchangeFilter) GetURLKey() string {
	return fmt.Sprintf("exch_%v", e.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (e ExchangeFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	e.Value = ExchangeType(result)
	return e
}

// IndexType represents the types of index filters
type IndexType string

// Default filter constants
const (
	SP500 IndexType = "sp500"
	DJIA  IndexType = "dji"
)

// IndexFilter is a filter type
type IndexFilter struct {
	Value IndexType
}

// GetURLKey returns the filter's url key
func (i IndexFilter) GetURLKey() string {
	return fmt.Sprintf("idx_%v", i.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (i IndexFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	i.Value = IndexType(result)
	return i
}

// SectorType represents the type of sector filters
type SectorType string

// Default filter constants
const (
	BasicMaterials        SectorType = "basicmaterials"
	CommunicationServices SectorType = "communicationservices"
	ConsumerCyclical      SectorType = "consumercyclical"
	ConsumerDefensive     SectorType = "consumerdefensive"
	Energy                SectorType = "energy"
	Financial             SectorType = "financial"
	Healthcare            SectorType = "healthcare"
	Industrials           SectorType = "industrials"
	RealEstate            SectorType = "realestate"
	Technology            SectorType = "technology"
	Utilities             SectorType = "utilities"
)

// SectorFilter is a filter type
type SectorFilter struct {
	Value SectorType
}

// GetURLKey returns the filter's url key
func (s SectorFilter) GetURLKey() string {
	return fmt.Sprintf("sec_%v", s.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (s SectorFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	s.Value = SectorType(result)
	return s
}

// IndustryType represents the type of industry filters
type IndustryType string

// Default filter constants
const (
	StocksOnly                         IndustryType = "stocksonly"
	ETF                                IndustryType = "exchangetradedfund"
	AdvertisingAgencies                IndustryType = "advertisingagencies"
	AerospaceAndDefense                IndustryType = "aerospacedefense"
	AgriculturalInputs                 IndustryType = "agriculturalinputs"
	Airlines                           IndustryType = "airlines"
	AirportsAirServices                IndustryType = "airportsairservices"
	Aluminum                           IndustryType = "aluminum"
	ApparelManufacturing               IndustryType = "apparelmanufacturing"
	ApparelRetail                      IndustryType = "apparelretail"
	AssetManagement                    IndustryType = "assetmanagement"
	AutoManufacturers                  IndustryType = "automanufacturers"
	AutoParts                          IndustryType = "autoparts"
	AutoTruckDealerships               IndustryType = "autotruckdealerships"
	BanksDiversified                   IndustryType = "banksdiversified"
	BanksRegional                      IndustryType = "banksregional"
	BeveragesBrewers                   IndustryType = "beveragesbrewers"
	BeveragesNonAlcoholic              IndustryType = "beveragesnonalcoholic"
	BeveragesWineriesDistilleries      IndustryType = "beverageswineriesdistilleries"
	Biotechnology                      IndustryType = "biotechnology"
	Broadcasting                       IndustryType = "broadcasting"
	BuildingMaterials                  IndustryType = "buildingmaterials"
	BuildingProductsEquipment          IndustryType = "buildingproductsequipment"
	BusinessEquipmentSupplies          IndustryType = "businessequipmentsupplies"
	CapitalMarkets                     IndustryType = "capitalmarkets"
	Chemicals                          IndustryType = "chemicals"
	ClosedEndFundDebt                  IndustryType = "closedendfunddebt"
	ClosedEndFundEquity                IndustryType = "closedendfundequity"
	ClosedEndFundForeign               IndustryType = "closedendfundforeign"
	CokingCoal                         IndustryType = "cokingcoal"
	CommunicationsEquipment            IndustryType = "communicationsequipment"
	ComputerHardware                   IndustryType = "computerhardware"
	Confectioners                      IndustryType = "confectioners"
	Conglomerates                      IndustryType = "conglomerates"
	ConsultingServices                 IndustryType = "consultingservices"
	ConsumerElectronics                IndustryType = "consumerelectronics"
	Copper                             IndustryType = "copper"
	CreditServices                     IndustryType = "creditservices"
	DepartmentStores                   IndustryType = "departmentstores"
	DiagnosticsResearch                IndustryType = "diagnosticsresearch"
	DiscountStores                     IndustryType = "discountstores"
	DrugManufacturersGeneral           IndustryType = "drugmanufacturersgeneral"
	DrugManufacturersSpecialtyGeneric  IndustryType = "drugmanufacturersspecialtygeneric"
	EducationTrainingServices          IndustryType = "educationtrainingservices"
	ElectricalEquipmentParts           IndustryType = "electricalequipmentparts"
	ElectronicComponents               IndustryType = "electroniccomponents"
	ElectronicGamingMultimedia         IndustryType = "electronicgamingmultimedia"
	ElectronicsComputerDistribution    IndustryType = "electronicscomputerdistribution"
	EngineeringConstruction            IndustryType = "engineeringconstruction"
	Entertainment                      IndustryType = "entertainment"
	FarmHeavyConstructionMachinery     IndustryType = "farmheavyconstructionmachinery"
	FarmProducts                       IndustryType = "farmproducts"
	FinancialConglomerates             IndustryType = "financialconglomerates"
	FinancialDataStockExchanges        IndustryType = "financialdatastockexchanges"
	FoodDistribution                   IndustryType = "fooddistribution"
	FootwearAccessories                IndustryType = "footwearaccessories"
	FurnishingsFixturesAppliances      IndustryType = "furnishingsfixturesappliances"
	Gambling                           IndustryType = "gambling"
	Gold                               IndustryType = "gold"
	GroceryStores                      IndustryType = "grocerystores"
	HealthcarePlans                    IndustryType = "healthcareplans"
	HealthInformationServices          IndustryType = "healthinformationservices"
	HomeImprovementRetail              IndustryType = "homeimprovementretail"
	HouseholdPersonalProducts          IndustryType = "householdpersonalproducts"
	IndustrialDistribution             IndustryType = "industrialdistribution"
	InformationTechnologyServices      IndustryType = "informationtechnologyservices"
	InfrastructureOperations           IndustryType = "infrastructureoperations"
	InsuranceBrokers                   IndustryType = "insurancebrokers"
	InsuranceDiversified               IndustryType = "insurancediversified"
	InsuranceLife                      IndustryType = "insurancelife"
	InsurancePropertyCasualty          IndustryType = "insurancepropertycasualty"
	InsuranceReinsurance               IndustryType = "insurancereinsurance"
	InsuranceSpecialty                 IndustryType = "insurancespecialty"
	IntegratedFreightLogistics         IndustryType = "integratedfreightlogistics"
	InternetContentInformation         IndustryType = "internetcontentinformation"
	InternetRetail                     IndustryType = "internetretail"
	Leisure                            IndustryType = "leisure"
	Lodging                            IndustryType = "lodging"
	LumberWoodProduction               IndustryType = "lumberwoodproduction"
	LuxuryGoods                        IndustryType = "luxurygoods"
	MarineShipping                     IndustryType = "marineshipping"
	MedicalCareFacilities              IndustryType = "medicalcarefacilities"
	MedicalDevices                     IndustryType = "medicaldevices"
	MedicalDistribution                IndustryType = "medicaldistribution"
	MedicalInstrumentsSupplies         IndustryType = "medicalinstrumentssupplies"
	MetalFabrication                   IndustryType = "metalfabrication"
	MortgageFinance                    IndustryType = "mortgagefinance"
	OilGasDrilling                     IndustryType = "oilgasdrilling"
	OilGasEP                           IndustryType = "oilgasep"
	OilGasEquipmentServices            IndustryType = "oilgasequipmentservices"
	OilGasIntegrated                   IndustryType = "oilgasintegrated"
	OilGasMidstream                    IndustryType = "oilgasmidstream"
	OilGasRefiningMarketing            IndustryType = "oilgasrefiningmarketing"
	OtherIndustrialMetalsMining        IndustryType = "otherindustrialmetalsmining"
	OtherPreciousMetalsMining          IndustryType = "otherpreciousmetalsmining"
	PackagedFoods                      IndustryType = "packagedfoods"
	PackagingContainers                IndustryType = "packagingcontainers"
	PaperPaperProducts                 IndustryType = "paperpaperproducts"
	PersonalServices                   IndustryType = "personalservices"
	PharmaceuticalRetailers            IndustryType = "pharmaceuticalretailers"
	PollutionTreatmentControls         IndustryType = "pollutiontreatmentcontrols"
	Publishing                         IndustryType = "publishing"
	Railroads                          IndustryType = "railroads"
	RealEstateDevelopment              IndustryType = "realestatedevelopment"
	RealEstateDiversified              IndustryType = "realestatediversified"
	RealEstateServices                 IndustryType = "realestateservices"
	RecreationalVehicles               IndustryType = "recreationalvehicles"
	REITDiversified                    IndustryType = "reitdiversified"
	REITHealthcareFacilities           IndustryType = "reithealthcarefacilities"
	REITHotelMotel                     IndustryType = "reithotelmotel"
	REITIndustrial                     IndustryType = "reitindustrial"
	REITMortgage                       IndustryType = "reitmortgage"
	REITOffice                         IndustryType = "reitoffice"
	REITResidential                    IndustryType = "reitresidential"
	REITRetail                         IndustryType = "reitretail"
	REITSpecialty                      IndustryType = "reitspecialty"
	RentalLeasingServices              IndustryType = "rentalleasingservices"
	ResidentialConstruction            IndustryType = "residentialconstruction"
	ResortsCasinos                     IndustryType = "resortscasinos"
	Restaurants                        IndustryType = "restaurants"
	ScientificTechnicalInstruments     IndustryType = "scientifictechnicalinstruments"
	SecurityProtectionServices         IndustryType = "securityprotectionservices"
	SemiconductorEquipmentMaterials    IndustryType = "semiconductorequipmentmaterials"
	Semiconductors                     IndustryType = "semiconductors"
	ShellCompanies                     IndustryType = "shellcompanies"
	Silver                             IndustryType = "silver"
	SoftwareApplication                IndustryType = "softwareapplication"
	SoftwareInfrastructure             IndustryType = "softwareinfrastructure"
	Solar                              IndustryType = "solar"
	SpecialtyBusinessServices          IndustryType = "specialtybusinessservices"
	SpecialtyChemicals                 IndustryType = "specialtychemicals"
	SpecialtyIndustrialMachinery       IndustryType = "specialtyindustrialmachinery"
	SpecialtyRetail                    IndustryType = "specialtyretail"
	StaffingEmploymentServices         IndustryType = "staffingemploymentservices"
	Steel                              IndustryType = "steel"
	TelecomServices                    IndustryType = "telecomservices"
	TextileManufacturing               IndustryType = "textilemanufacturing"
	ThermalCoal                        IndustryType = "thermalcoal"
	Tobacco                            IndustryType = "tobacco"
	ToolsAccessories                   IndustryType = "toolsaccessories"
	TravelServices                     IndustryType = "travelservices"
	Trucking                           IndustryType = "trucking"
	Uranium                            IndustryType = "uranium"
	UtilitiesDiversified               IndustryType = "utilitiesdiversified"
	UtilitiesIndependentPowerProducers IndustryType = "utilitiesindependentpowerproducers"
	UtilitiesRegulatedElectric         IndustryType = "utilitiesregulatedelectric"
	UtilitiesRegulatedGas              IndustryType = "utilitiesregulatedgas"
	UtilitiesRegulatedWater            IndustryType = "utilitiesregulatedwater"
	UtilitiesRenewable                 IndustryType = "utilitiesrenewable"
	WasteManagement                    IndustryType = "wastemanagement"
)

// IndustryFilter is a filter type
type IndustryFilter struct {
	Value IndustryType
}

// GetURLKey returns the filter's url key
func (i IndustryFilter) GetURLKey() string {
	return fmt.Sprintf("ind_%v", i.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (i IndustryFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	i.Value = IndustryType(result)
	return i
}

// CountryType represents the type of country filters
type CountryType string

// Default filter constants
const (
	USA           CountryType = "usa"
	NotUSA        CountryType = "notusa"
	Asia          CountryType = "asia"
	Europe        CountryType = "europe"
	LatinAmerica  CountryType = "latinamerica"
	BRIC          CountryType = "bric"
	Argentina     CountryType = "argentina"
	Australia     CountryType = "australia"
	Bahamas       CountryType = "bahamas"
	Belgium       CountryType = "belgium"
	BeNeLux       CountryType = "benelux"
	Bermuda       CountryType = "bermuda"
	Brazil        CountryType = "brazil"
	Canada        CountryType = "canada"
	CaymanIslands CountryType = "caymanislands"
	Chile         CountryType = "chile"
	China         CountryType = "china"
	ChinaHongKong CountryType = "chinahongkong"
	Colombia      CountryType = "colombia"
	Cyprus        CountryType = "cyprus"
	Denmark       CountryType = "denmark"
	Finland       CountryType = "finland"
	France        CountryType = "france"
	Germany       CountryType = "germany"
	Greece        CountryType = "greece"
	HongKong      CountryType = "hongkong"
	Hungary       CountryType = "hungary"
	Iceland       CountryType = "iceland"
	India         CountryType = "india"
	Indonesia     CountryType = "indonesia"
	Ireland       CountryType = "ireland"
	Israel        CountryType = "israel"
	Italy         CountryType = "italy"
	Japan         CountryType = "japan"
	Kazakhstan    CountryType = "kazakhstan"
	Luxembourg    CountryType = "luxembourg"
	Malaysia      CountryType = "malaysia"
	Malta         CountryType = "malta"
	Mexico        CountryType = "mexico"
	Monaco        CountryType = "monaco"
	Netherlands   CountryType = "netherlands"
	NewZealand    CountryType = "newzealand"
	Norway        CountryType = "norway"
	Panama        CountryType = "panama"
	Peru          CountryType = "peru"
	Philippines   CountryType = "philippines"
	Portugal      CountryType = "portugal"
	Russia        CountryType = "russia"
	Singapore     CountryType = "singapore"
	SouthAfrica   CountryType = "southafrica"
	SouthKorea    CountryType = "southkorea"
	Spain         CountryType = "spain"
	Sweden        CountryType = "sweden"
	Switzerland   CountryType = "switzerland"
	Taiwan        CountryType = "taiwan"
	Turkey        CountryType = "turkey"
	UAE           CountryType = "unitedarabemirates"
	UK            CountryType = "unitedkingdom"
	Uruguay       CountryType = "uruguay"
)

// CountryFilter is a filter type
type CountryFilter struct {
	Value CountryType
}

// GetURLKey returns the filter's url key
func (c CountryFilter) GetURLKey() string {
	return fmt.Sprintf("geo_%v", c.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (c CountryFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	c.Value = CountryType(result)
	return c
}

// MarketCapType represents the types of market cap filters
type MarketCapType string

// Default filter constants
const (
	MegaOver200B   MarketCapType = "mega"
	Large10to200B  MarketCapType = "large"
	Mid2to10B      MarketCapType = "mid"
	Small300Mto2B  MarketCapType = "small"
	Micro50to300M  MarketCapType = "micro"
	NanoUnder50M   MarketCapType = "nano"
	LargeOver10B   MarketCapType = "largeover"
	MidOver2B      MarketCapType = "midover"
	SmallOver300M  MarketCapType = "smallover"
	MicroOver50M   MarketCapType = "microover"
	LargeUnder200B MarketCapType = "largeunder"
	MidUnder10B    MarketCapType = "midunder"
	SmallUnder2B   MarketCapType = "smallunder"
	MicroUnder300M MarketCapType = "microunder"
)

// MarketCapFilter is a filter type
type MarketCapFilter struct {
	Value MarketCapType
}

// GetURLKey returns the filter's url key
func (m MarketCapFilter) GetURLKey() string {
	return fmt.Sprintf("cap_%v", m.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (m MarketCapFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	m.Value = MarketCapType(result)
	return m
}

// DividendYieldType represents the type of dividend yield filters
type DividendYieldType string

// Default filter constants
const (
	DYNone     DividendYieldType = "none"
	DYPositive DividendYieldType = "pos"
	DYHigh     DividendYieldType = "high"
	DYVeryHigh DividendYieldType = "veryhigh"
	DYOver1    DividendYieldType = "o1"
	DYOver2    DividendYieldType = "o2"
	DYOver3    DividendYieldType = "o3"
	DYOver4    DividendYieldType = "o4"
	DYOver5    DividendYieldType = "o5"
	DYOver6    DividendYieldType = "o6"
	DYOver7    DividendYieldType = "o7"
	DYOver8    DividendYieldType = "o8"
	DYOver9    DividendYieldType = "o9"
	DYOver10   DividendYieldType = "o10"
)

// DividendYieldFilter is a filter type
type DividendYieldFilter struct {
	Value DividendYieldType
}

// GetURLKey returns the filter's url key
func (d DividendYieldFilter) GetURLKey() string {
	return fmt.Sprintf("fa_div_%v", d.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (d DividendYieldFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	d.Value = DividendYieldType(result)
	return d
}

// FloatShortType represents the type of short selling filters
type FloatShortType string

// Default filter constants
const (
	FSLow     FloatShortType = "low"
	FSHigh    FloatShortType = "high"
	FSUnder5  FloatShortType = "u5"
	FSUnder10 FloatShortType = "u10"
	FSUnder15 FloatShortType = "u15"
	FSUnder20 FloatShortType = "u20"
	FSUnder25 FloatShortType = "u25"
	FSUnder30 FloatShortType = "u30"
	FSOver5   FloatShortType = "o5"
	FSOver10  FloatShortType = "o10"
	FSOver15  FloatShortType = "o15"
	FSOver20  FloatShortType = "o20"
	FSOver25  FloatShortType = "o25"
	FSOver30  FloatShortType = "o30"
)

// FloatShortFilter is a filter type
type FloatShortFilter struct {
	Value FloatShortType
}

// GetURLKey returns the filter's url key
func (f FloatShortFilter) GetURLKey() string {
	return fmt.Sprintf("sh_short_%v", f.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (f FloatShortFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	f.Value = FloatShortType(result)
	return f
}

// RecommendationType represents the type of analyst recommendation filters
type RecommendationType string

// Default filter constants
const (
	StrongBuy  RecommendationType = "strongbuy"
	BuyBetter  RecommendationType = "buybetter"
	Buy        RecommendationType = "buy"
	HoldBetter RecommendationType = "holdbetter"
	Hold       RecommendationType = "hold"
	HoldWorse  RecommendationType = "holdworse"
	Sell       RecommendationType = "sell"
	SellWorse  RecommendationType = "sellworse"
	StrongSell RecommendationType = "strongsell"
)

// RecommendationFilter is a filter type
type RecommendationFilter struct {
	Value RecommendationType
}

// GetURLKey returns the filter's url key
func (r RecommendationFilter) GetURLKey() string {
	return fmt.Sprintf("an_recom_%v", r.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (r RecommendationFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	r.Value = RecommendationType(result)
	return r
}

// OptionShortType represents the type of optionable/shortable filters
type OptionShortType string

// Default filter constants
const (
	Option      OptionShortType = "option"
	Short       OptionShortType = "short"
	OptionShort OptionShortType = "optionshort"
)

// OptionShortFilter is a filter type
type OptionShortFilter struct {
	Value OptionShortType
}

// GetURLKey returns the filter's url key
func (o OptionShortFilter) GetURLKey() string {
	return fmt.Sprintf("sh_opt_%v", o.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (o OptionShortFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	o.Value = OptionShortType(result)
	return o
}

// EarningsDateType represents the type of earnings date filters
type EarningsDateType string

// Default filter constants
const (
	EDToday                     EarningsDateType = "today"
	EDTodayBeforeMarketOpen     EarningsDateType = "todaybefore"
	EDTodayAfterMarketClose     EarningsDateType = "todayafter"
	EDTomorrow                  EarningsDateType = "tomorrow"
	EDTomorrowBeforeMarketOpen  EarningsDateType = "tomorrowbefore"
	EDTomorrowAfterMarketClose  EarningsDateType = "tomorrowafter"
	EDYesterday                 EarningsDateType = "yesterday"
	EDYesterdayBeforeMarketOpen EarningsDateType = "yesterdaybefore"
	EDYesterdayAfterMarketClose EarningsDateType = "yesterdayafter"
	EDNext5Days                 EarningsDateType = "nextdays5"
	EDPrevious5Days             EarningsDateType = "prevdays5"
	EDThisWeek                  EarningsDateType = "thisweek"
	EDNextWeek                  EarningsDateType = "nextweek"
	EDPrevWeek                  EarningsDateType = "prevweek"
	EDThisMonth                 EarningsDateType = "thismonth"
)

// EarningsDateFilter is a filter type
type EarningsDateFilter struct {
	Value EarningsDateType
}

// GetURLKey returns the filter's url key
func (e EarningsDateFilter) GetURLKey() string {
	return fmt.Sprintf("earningsdate_%v", e.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (e EarningsDateFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	e.Value = EarningsDateType(result)
	return e
}

// AverageVolumeType represents the type of average volume filters
type AverageVolumeType string

// Default filter constants
const (
	AvgVolUnder50K         AverageVolumeType = "u50"
	AvgVolUnder100K        AverageVolumeType = "u100"
	AvgVolUnder500K        AverageVolumeType = "u500"
	AvgVolUnder750K        AverageVolumeType = "u750"
	AvgVolUnder1M          AverageVolumeType = "u1000"
	AvgVolOver50K          AverageVolumeType = "o50"
	AvgVolOver100K         AverageVolumeType = "o100"
	AvgVolOver200K         AverageVolumeType = "o200"
	AvgVolOver300K         AverageVolumeType = "o300"
	AvgVolOver400K         AverageVolumeType = "o400"
	AvgVolOver500K         AverageVolumeType = "o500"
	AvgVolOver750K         AverageVolumeType = "o750"
	AvgVolOver1M           AverageVolumeType = "o1000"
	AvgVolOver2M           AverageVolumeType = "o2000"
	AvgVolBetween100to500K AverageVolumeType = "100to500"
	AvgVolBetween100Kto1M  AverageVolumeType = "100to1000"
	AvgVolBetween500Kto1M  AverageVolumeType = "500to1000"
	AvgVolBetween500Kto10M AverageVolumeType = "500to10000"
)

// AverageVolumeFilter is a filter type
type AverageVolumeFilter struct {
	Value AverageVolumeType
}

// GetURLKey returns the filter's url key
func (a AverageVolumeFilter) GetURLKey() string {
	return fmt.Sprintf("sh_avgvol_%v", a.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (a AverageVolumeFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	a.Value = AverageVolumeType(result)
	return a
}

// RelativeVolumeType represents the type of relative volume filters
type RelativeVolumeType string

// Default filter constants
const (
	RVOver10        RelativeVolumeType = "o10"
	RVOver5         RelativeVolumeType = "o5"
	RVOver3         RelativeVolumeType = "o3"
	RVOver2         RelativeVolumeType = "o2"
	RVOver1point5   RelativeVolumeType = "o1.5"
	RVOver1         RelativeVolumeType = "o1"
	RVOver0point75  RelativeVolumeType = "o0.75"
	RVOver0point5   RelativeVolumeType = "o0.5"
	RVOver0point25  RelativeVolumeType = "o0.25"
	RVUnder2        RelativeVolumeType = "u2"
	RVUnder1point5  RelativeVolumeType = "u1.5"
	RVUnder1        RelativeVolumeType = "u1"
	RVUnder0point75 RelativeVolumeType = "u0.75"
	RVUnder0point5  RelativeVolumeType = "u0.5"
	RVUnder0point25 RelativeVolumeType = "u0.25"
	RVUnder0point1  RelativeVolumeType = "u0.1"
)

// RelativeVolumeFilter is a filter type
type RelativeVolumeFilter struct {
	Value RelativeVolumeType
}

// GetURLKey returns the filter's url key
func (r RelativeVolumeFilter) GetURLKey() string {
	return fmt.Sprintf("sh_relvol_%v", r.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (r RelativeVolumeFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	r.Value = RelativeVolumeType(result)
	return r
}

// CurrentVolumeType represents the type of current volume filters
type CurrentVolumeType string

// Default filter constants
const (
	CurVolUnder50K  CurrentVolumeType = "u50"
	CurVolUnder100K CurrentVolumeType = "u100"
	CurVolUnder500K CurrentVolumeType = "u500"
	CurVolUnder750K CurrentVolumeType = "u750"
	CurVolUnder1M   CurrentVolumeType = "u1000"
	CurVolOver0     CurrentVolumeType = "o0"
	CurVolOver50K   CurrentVolumeType = "o50"
	CurVolOver100K  CurrentVolumeType = "o100"
	CurVolOver200K  CurrentVolumeType = "o200"
	CurVolOver300K  CurrentVolumeType = "o300"
	CurVolOver400K  CurrentVolumeType = "o400"
	CurVolOver500K  CurrentVolumeType = "o500"
	CurVolOver750K  CurrentVolumeType = "o750"
	CurVolOver1M    CurrentVolumeType = "o1000"
	CurVolOver2M    CurrentVolumeType = "o2000"
	CurVolOver5M    CurrentVolumeType = "o5000"
	CurVolOver10M   CurrentVolumeType = "o10000"
	CurVolOver20M   CurrentVolumeType = "o20000"
)

// CurrentVolumeFilter is a filter type
type CurrentVolumeFilter struct {
	Value CurrentVolumeType
}

// GetURLKey returns the filter's url key
func (c CurrentVolumeFilter) GetURLKey() string {
	return fmt.Sprintf("sh_curvol_%v", c.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (c CurrentVolumeFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	c.Value = CurrentVolumeType(result)
	return c
}

// PriceType represents the type of price filters
type PriceType string

// Default filter constants
const (
	PriceUnder1  PriceType = "u1"
	PriceUnder2  PriceType = "u2"
	PriceUnder3  PriceType = "u3"
	PriceUnder4  PriceType = "u4"
	PriceUnder5  PriceType = "u5"
	PriceUnder7  PriceType = "u7"
	PriceUnder10 PriceType = "u10"
	PriceUnder15 PriceType = "u15"
	PriceUnder20 PriceType = "u20"
	PriceUnder30 PriceType = "u30"
	PriceUnder40 PriceType = "u40"
	PriceUnder50 PriceType = "u50"
	PriceOver1   PriceType = "o1"
	PriceOver2   PriceType = "o2"
	PriceOver3   PriceType = "o3"
	PriceOver4   PriceType = "o4"
	PriceOver5   PriceType = "o5"
	PriceOver7   PriceType = "o7"
	PriceOver10  PriceType = "o10"
	PriceOver15  PriceType = "o15"
	PriceOver20  PriceType = "o20"
	PriceOver30  PriceType = "o30"
	PriceOver40  PriceType = "o40"
	PriceOver50  PriceType = "o50"
	PriceOver60  PriceType = "o60"
	PriceOver70  PriceType = "o70"
	PriceOver80  PriceType = "o80"
	PriceOver90  PriceType = "o90"
	PriceOver100 PriceType = "o100"
	Price1to5    PriceType = "1to5"
	Price1to10   PriceType = "1to10"
	Price1to20   PriceType = "1to20"
	Price5to10   PriceType = "5to10"
	Price5to20   PriceType = "5to20"
	Price5to50   PriceType = "5to50"
	Price10to20  PriceType = "10to20"
	Price10to50  PriceType = "10to50"
	Price20to50  PriceType = "20to50"
	Price50to100 PriceType = "50to100"
)

// PriceFilter is a filter type
type PriceFilter struct {
	Value PriceType
}

// GetURLKey returns the filter's url key
func (p PriceFilter) GetURLKey() string {
	return fmt.Sprintf("sh_price_%v", p.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (p PriceFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	p.Value = PriceType(result)
	return p
}

// TargetPriceType represents the type of target price filters
type TargetPriceType string

// Default filter constants
const (
	TargetAbovePriceBy50Percent TargetPriceType = "a50"
	TargetAbovePriceBy40Percent TargetPriceType = "a40"
	TargetAbovePriceBy30Percent TargetPriceType = "a30"
	TargetAbovePriceBy20Percent TargetPriceType = "a20"
	TargetAbovePriceBy10Percent TargetPriceType = "a10"
	TargetAbovePriceBy5Percent  TargetPriceType = "a5"
	TargetAbovePrice            TargetPriceType = "above"
	TargetBelowPrice            TargetPriceType = "below"
	TargetBelowPriceBy50Percent TargetPriceType = "b50"
	TargetBelowPriceBy40Percent TargetPriceType = "b40"
	TargetBelowPriceBy30Percent TargetPriceType = "b30"
	TargetBelowPriceBy20Percent TargetPriceType = "b20"
	TargetBelowPriceBy10Percent TargetPriceType = "b10"
	TargetBelowPriceBy5Percent  TargetPriceType = "b5"
)

// TargetPriceFilter is a filter type
type TargetPriceFilter struct {
	Value TargetPriceType
}

// GetURLKey returns the filter's url key
func (t TargetPriceFilter) GetURLKey() string {
	return fmt.Sprintf("targetprice_%v", t.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (t TargetPriceFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	t.Value = TargetPriceType(result)
	return t
}

// IPODateType represents the type of IPO Date filters
type IPODateType string

// Default filter constants
const (
	IDToday              IPODateType = "today"
	IDYesterday          IPODateType = "yesterday"
	IDPreviousWeek       IPODateType = "prevweek"
	IDPreviousMonth      IPODateType = "prevmonth"
	IDPreviousQuarter    IPODateType = "prevquarter"
	IDPreviousYear       IPODateType = "prevyear"
	IDPrevious2Years     IPODateType = "prev2yrs"
	IDPrevious3Years     IPODateType = "prev3yrs"
	IDPrevious5Years     IPODateType = "prev5yrs"
	IDMoreThan1YearAgo   IPODateType = "more1"
	IDMoreThan5YearsAgo  IPODateType = "more5"
	IDMoreThan10YearsAgo IPODateType = "more10"
	IDMoreThan15YearsAgo IPODateType = "more15"
	IDMoreThan20YearsAgo IPODateType = "more20"
	IDMoreThan25YearsAgo IPODateType = "more25"
)

// IPODateFilter is a filter type
type IPODateFilter struct {
	Value IPODateType
}

// GetURLKey returns the filter's url key
func (i IPODateFilter) GetURLKey() string {
	return fmt.Sprintf("ipodate_%v", i.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (i IPODateFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	i.Value = IPODateType(result)
	return i
}

// SharesOutstandingType represents the type of shares outstanding filters
type SharesOutstandingType string

// Default filter constants
const (
	SOUnder1M   SharesOutstandingType = "u1"
	SOUnder5M   SharesOutstandingType = "u5"
	SOUnder10M  SharesOutstandingType = "u10"
	SOUnder20M  SharesOutstandingType = "u20"
	SOUnder50M  SharesOutstandingType = "u50"
	SOUnder100M SharesOutstandingType = "u100"
	SOOver1M    SharesOutstandingType = "o1"
	SOOver2M    SharesOutstandingType = "o2"
	SOOver5M    SharesOutstandingType = "o5"
	SOOver10M   SharesOutstandingType = "o10"
	SOOver20M   SharesOutstandingType = "o20"
	SOOver50M   SharesOutstandingType = "o50"
	SOOver100M  SharesOutstandingType = "o100"
	SOOver200M  SharesOutstandingType = "o200"
	SOOver500M  SharesOutstandingType = "o500"
	SOOver1000M SharesOutstandingType = "o1000"
)

// SharesOutstandingFilter is a filter type
type SharesOutstandingFilter struct {
	Value SharesOutstandingType
}

// GetURLKey returns the filter's url key
func (s SharesOutstandingFilter) GetURLKey() string {
	return fmt.Sprintf("sh_outstanding_%v", s.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (s SharesOutstandingFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	s.Value = SharesOutstandingType(result)
	return s
}

// FloatType represents the type of float filters
type FloatType = SharesOutstandingType

// FloatFilter is a filter type
type FloatFilter struct {
	Value FloatType
}

// GetURLKey returns the filter's url key
func (f FloatFilter) GetURLKey() string {
	return fmt.Sprintf("sh_float_%v", f.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (f FloatFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	f.Value = FloatType(result)
	return f
}

/***************************************************************************
*****	FUNDAMENTALS   *****************************************************
***************************************************************************/

// PEType represents the types of PE filters
type PEType string

// Default filter constants
const (
	PELow        PEType = "low"
	PEProfitable PEType = "profitable"
	PEHigh       PEType = "high"
	PEUnder5     PEType = "u5"
	PEUnder10    PEType = "u10"
	PEUnder15    PEType = "u15"
	PEUnder20    PEType = "u20"
	PEUnder25    PEType = "u25"
	PEUnder30    PEType = "u30"
	PEUnder35    PEType = "u35"
	PEUnder40    PEType = "u40"
	PEUnder45    PEType = "u45"
	PEUnder50    PEType = "u50"
	PEOver5      PEType = "o5"
	PEOver10     PEType = "o10"
	PEOver15     PEType = "o15"
	PEOver20     PEType = "o20"
	PEOver25     PEType = "o25"
	PEOver30     PEType = "o30"
	PEOver35     PEType = "o35"
	PEOver40     PEType = "o40"
	PEOver45     PEType = "o45"
	PEOver50     PEType = "o50"
)

// PEFilter is a filter type
type PEFilter struct {
	Value PEType
}

// GetURLKey returns the filter's url key
func (p PEFilter) GetURLKey() string {
	return fmt.Sprintf("fa_pe_%v", p.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (p PEFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	p.Value = PEType(result)
	return p
}

// ForwardPEType represents the types of ForwardPE filters
// Same categories as PEType, so those will be reused
type ForwardPEType = PEType

// ForwardPEFilter is a filter type
type ForwardPEFilter struct {
	Value ForwardPEType
}

// GetURLKey returns the filter's url key
func (f ForwardPEFilter) GetURLKey() string {
	return fmt.Sprintf("fa_fpe_%v", f.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (f ForwardPEFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	f.Value = ForwardPEType(result)
	return f
}

// PEGType represents the types of PEG filters
type PEGType string

// Default filter constants
const (
	PEGLow    PEGType = "low"
	PEGHigh   PEGType = "high"
	PEGUnder1 PEGType = "u1"
	PEGUnder2 PEGType = "u2"
	PEGUnder3 PEGType = "u3"
	PEGOver1  PEGType = "o1"
	PEGOver2  PEGType = "o2"
	PEGOver3  PEGType = "o3"
)

// PEGFilter is a filter type
type PEGFilter struct {
	Value PEGType
}

// GetURLKey returns the filter's url key
func (p PEGFilter) GetURLKey() string {
	return fmt.Sprintf("fa_peg_%v", p.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (p PEGFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	p.Value = PEGType(result)
	return p
}

// PriceSalesType represents the types of price/sales filters
type PriceSalesType string

// Default filter constants
const (
	PSLow     PriceSalesType = "low"
	PSHigh    PriceSalesType = "high"
	PSUnder1  PriceSalesType = "u1"
	PSUnder2  PriceSalesType = "u2"
	PSUnder3  PriceSalesType = "u3"
	PSUnder4  PriceSalesType = "u4"
	PSUnder5  PriceSalesType = "u5"
	PSUnder6  PriceSalesType = "u6"
	PSUnder7  PriceSalesType = "u7"
	PSUnder8  PriceSalesType = "u8"
	PSUnder9  PriceSalesType = "u9"
	PSUnder10 PriceSalesType = "u10"
	PSOver1   PriceSalesType = "o1"
	PSOver2   PriceSalesType = "o2"
	PSOver3   PriceSalesType = "o3"
	PSOver4   PriceSalesType = "o4"
	PSOver5   PriceSalesType = "o5"
	PSOver6   PriceSalesType = "o6"
	PSOver7   PriceSalesType = "o7"
	PSOver8   PriceSalesType = "o8"
	PSOver9   PriceSalesType = "o9"
	PSOver10  PriceSalesType = "o10"
)

// PriceSalesFilter is a filter type
type PriceSalesFilter struct {
	Value PriceSalesType
}

// GetURLKey returns the filter's url key
func (p PriceSalesFilter) GetURLKey() string {
	return fmt.Sprintf("fa_ps_%v", p.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (p PriceSalesFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	p.Value = PriceSalesType(result)
	return p
}

// PriceBookType represents the types of price/book filters
type PriceBookType = PriceSalesType

// PriceBookFilter is a filter type
type PriceBookFilter struct {
	Value PriceBookType
}

// GetURLKey returns the filter's url key
func (p PriceBookFilter) GetURLKey() string {
	return fmt.Sprintf("fa_pb_%v", p.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (p PriceBookFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	p.Value = PriceBookType(result)
	return p
}

// PriceCashType represents the types of price/cash filters
type PriceCashType string

// Default filter constants
const (
	PCLow     PriceCashType = "low"
	PCHigh    PriceCashType = "high"
	PCUnder1  PriceCashType = "u1"
	PCUnder2  PriceCashType = "u2"
	PCUnder3  PriceCashType = "u3"
	PCUnder4  PriceCashType = "u4"
	PCUnder5  PriceCashType = "u5"
	PCUnder6  PriceCashType = "u6"
	PCUnder7  PriceCashType = "u7"
	PCUnder8  PriceCashType = "u8"
	PCUnder9  PriceCashType = "u9"
	PCUnder10 PriceCashType = "u10"
	PCOver1   PriceCashType = "o1"
	PCOver2   PriceCashType = "o2"
	PCOver3   PriceCashType = "o3"
	PCOver4   PriceCashType = "o4"
	PCOver5   PriceCashType = "o5"
	PCOver6   PriceCashType = "o6"
	PCOver7   PriceCashType = "o7"
	PCOver8   PriceCashType = "o8"
	PCOver9   PriceCashType = "o9"
	PCOver10  PriceCashType = "o10"
	PCOver20  PriceCashType = "o20"
	PCOver30  PriceCashType = "o30"
	PCOver40  PriceCashType = "o40"
	PCOver50  PriceCashType = "o50"
)

// PriceCashFilter is a filter type
type PriceCashFilter struct {
	Value PriceCashType
}

// GetURLKey returns the filter's url key
func (p PriceCashFilter) GetURLKey() string {
	return fmt.Sprintf("fa_pc_%v", p.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (p PriceCashFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	p.Value = PriceCashType(result)
	return p
}

// PriceFCFType represents the types of price/FCF (free cash flow) filters
type PriceFCFType string

// Default filter constants
const (
	PFCFLow      PriceFCFType = "low"
	PFCFHigh     PriceFCFType = "high"
	PFCFUnder5   PriceFCFType = "u5"
	PFCFUnder10  PriceFCFType = "u10"
	PFCFUnder15  PriceFCFType = "u15"
	PFCFUnder20  PriceFCFType = "u20"
	PFCFUnder25  PriceFCFType = "u25"
	PFCFUnder30  PriceFCFType = "u30"
	PFCFUnder35  PriceFCFType = "u35"
	PFCFUnder40  PriceFCFType = "u40"
	PFCFUnder45  PriceFCFType = "u45"
	PFCFUnder50  PriceFCFType = "u50"
	PFCFUnder60  PriceFCFType = "u60"
	PFCFUnder70  PriceFCFType = "u70"
	PFCFUnder80  PriceFCFType = "u80"
	PFCFUnder90  PriceFCFType = "u90"
	PFCFUnder100 PriceFCFType = "u100"
	PFCFOver5    PriceFCFType = "o5"
	PFCFOver10   PriceFCFType = "o10"
	PFCFOver15   PriceFCFType = "o15"
	PFCFOver20   PriceFCFType = "o20"
	PFCFOver25   PriceFCFType = "o25"
	PFCFOver30   PriceFCFType = "o30"
	PFCFOver35   PriceFCFType = "o35"
	PFCFOver40   PriceFCFType = "o40"
	PFCFOver45   PriceFCFType = "o45"
	PFCFOver50   PriceFCFType = "o50"
	PFCFOver60   PriceFCFType = "o60"
	PFCFOver70   PriceFCFType = "o70"
	PFCFOver80   PriceFCFType = "o80"
	PFCFOver90   PriceFCFType = "o90"
	PFCFOver100  PriceFCFType = "o100"
)

// PriceFCFFilter is a filter type
type PriceFCFFilter struct {
	Value PriceFCFType
}

// GetURLKey returns the filter's url key
func (p PriceFCFFilter) GetURLKey() string {
	return fmt.Sprintf("fa_pfcf_%v", p.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (p PriceFCFFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	p.Value = PriceFCFType(result)
	return p
}

// GrowthType represents the types of growth (%) filters
type GrowthType string

// Default filter constants
const (
	GrowthNegative    GrowthType = "neg"
	GrowthPositive    GrowthType = "pos"
	GrowthPositiveLow GrowthType = "poslow"
	GrowthHigh        GrowthType = "high"
	GrowthUnder5      GrowthType = "u5"
	GrowthUnder10     GrowthType = "u10"
	GrowthUnder15     GrowthType = "u15"
	GrowthUnder20     GrowthType = "u20"
	GrowthUnder25     GrowthType = "u25"
	GrowthUnder30     GrowthType = "u30"
	GrowthOver5       GrowthType = "o5"
	GrowthOver10      GrowthType = "o10"
	GrowthOver15      GrowthType = "o15"
	GrowthOver20      GrowthType = "o20"
	GrowthOver25      GrowthType = "o25"
	GrowthOver30      GrowthType = "o30"
)

// EPSGrowthThisYearFilter is a filter type
type EPSGrowthThisYearFilter struct {
	Value GrowthType
}

// GetURLKey returns the filter's url key
func (e EPSGrowthThisYearFilter) GetURLKey() string {
	return fmt.Sprintf("fa_epsyoy_%v", e.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (e EPSGrowthThisYearFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	e.Value = GrowthType(result)
	return e
}

// EPSGrowthNextYearFilter is a filter type
type EPSGrowthNextYearFilter struct {
	Value GrowthType
}

// GetURLKey returns the filter's url key
func (e EPSGrowthNextYearFilter) GetURLKey() string {
	return fmt.Sprintf("fa_epsyoy1_%v", e.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (e EPSGrowthNextYearFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	e.Value = GrowthType(result)
	return e
}

// EPSGrowthPast5YearsFilter is a filter type
type EPSGrowthPast5YearsFilter struct {
	Value GrowthType
}

// GetURLKey returns the filter's url key
func (e EPSGrowthPast5YearsFilter) GetURLKey() string {
	return fmt.Sprintf("fa_eps5years_%v", e.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (e EPSGrowthPast5YearsFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	e.Value = GrowthType(result)
	return e
}

// EPSGrowthNext5YearsFilter is a filter type
type EPSGrowthNext5YearsFilter struct {
	Value GrowthType
}

// GetURLKey returns the filter's url key
func (e EPSGrowthNext5YearsFilter) GetURLKey() string {
	return fmt.Sprintf("fa_estltgrowth_%v", e.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (e EPSGrowthNext5YearsFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	e.Value = GrowthType(result)
	return e
}

// SalesGrowthPast5YearsFilter is a filter type
type SalesGrowthPast5YearsFilter struct {
	Value GrowthType
}

// GetURLKey returns the filter's url key
func (s SalesGrowthPast5YearsFilter) GetURLKey() string {
	return fmt.Sprintf("fa_sales5years_%v", s.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (s SalesGrowthPast5YearsFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	s.Value = GrowthType(result)
	return s
}

// EPSGrowthQtrOverQtrFilter is a filter type
type EPSGrowthQtrOverQtrFilter struct {
	Value GrowthType
}

// GetURLKey returns the filter's url key
func (e EPSGrowthQtrOverQtrFilter) GetURLKey() string {
	return fmt.Sprintf("fa_epsqoq_%v", e.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (e EPSGrowthQtrOverQtrFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	e.Value = GrowthType(result)
	return e
}

// SalesGrowthQtrOverQtrFilter is a filter type
type SalesGrowthQtrOverQtrFilter struct {
	Value GrowthType
}

// GetURLKey returns the filter's url key
func (s SalesGrowthQtrOverQtrFilter) GetURLKey() string {
	return fmt.Sprintf("fa_salesqoq_%v", s.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (s SalesGrowthQtrOverQtrFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	s.Value = GrowthType(result)
	return s
}

// ReturnType represents the types of % return on assets, equity, and investment filters
type ReturnType string

// Default filter constants
const (
	ReturnPositive     ReturnType = "pos"
	ReturnNegative     ReturnType = "neg"
	ReturnVeryPositive ReturnType = "verypos"
	ReturnVeryNegative ReturnType = "veryneg"
	ReturnUnderNeg50   ReturnType = "u-50"
	ReturnUnderNeg45   ReturnType = "u-45"
	ReturnUnderNeg40   ReturnType = "u-40"
	ReturnUnderNeg35   ReturnType = "u-35"
	ReturnUnderNeg30   ReturnType = "u-30"
	ReturnUnderNeg25   ReturnType = "u-25"
	ReturnUnderNeg20   ReturnType = "u-20"
	ReturnUnderNeg15   ReturnType = "u-15"
	ReturnUnderNeg10   ReturnType = "u-10"
	ReturnUnderNeg5    ReturnType = "u-5"
	ReturnOver50       ReturnType = "o50"
	ReturnOver45       ReturnType = "o45"
	ReturnOver40       ReturnType = "o40"
	ReturnOver35       ReturnType = "o35"
	ReturnOver30       ReturnType = "o30"
	ReturnOver25       ReturnType = "o25"
	ReturnOver20       ReturnType = "o20"
	ReturnOver15       ReturnType = "o15"
	ReturnOver10       ReturnType = "o10"
	ReturnOver5        ReturnType = "o5"
)

// ROAFilter is a filter type
type ROAFilter struct {
	Value ReturnType
}

// GetURLKey returns the filter's url key
func (r ROAFilter) GetURLKey() string {
	return fmt.Sprintf("fa_roa_%v", r.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (r ROAFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	r.Value = ReturnType(result)
	return r
}

// ROEFilter is a filter type
type ROEFilter struct {
	Value ReturnType
}

// GetURLKey returns the filter's url key
func (r ROEFilter) GetURLKey() string {
	return fmt.Sprintf("fa_roe_%v", r.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (r ROEFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	r.Value = ReturnType(result)
	return r
}

// ROIFilter is a filter type
type ROIFilter struct {
	Value ReturnType
}

// GetURLKey returns the filter's url key
func (r ROIFilter) GetURLKey() string {
	return fmt.Sprintf("fa_roi_%v", r.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (r ROIFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	r.Value = ReturnType(result)
	return r
}

// AssetRatioType represents the types of asset ratio filters
type AssetRatioType string

// Default filter constants
const (
	RatioHigh         AssetRatioType = "high"
	RatioLow          AssetRatioType = "low"
	RatioUnder1       AssetRatioType = "u1"
	RatioUnder0point5 AssetRatioType = "u0.5"
	RatioOver0point5  AssetRatioType = "o0.5"
	RatioOver1        AssetRatioType = "o1"
	RatioOver1point5  AssetRatioType = "o1.5"
	RatioOver2        AssetRatioType = "o2"
	RatioOver3        AssetRatioType = "o3"
	RatioOver4        AssetRatioType = "o4"
	RatioOver5        AssetRatioType = "o5"
	RatioOver10       AssetRatioType = "o10"
)

// CurrentRatioFilter is a filter type
type CurrentRatioFilter struct {
	Value AssetRatioType
}

// GetURLKey returns the filter's url key
func (c CurrentRatioFilter) GetURLKey() string {
	return fmt.Sprintf("fa_curratio_%v", c.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (c CurrentRatioFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	c.Value = AssetRatioType(result)
	return c
}

// QuickRatioFilter is a filter type
type QuickRatioFilter struct {
	Value AssetRatioType
}

// GetURLKey returns the filter's url key
func (q QuickRatioFilter) GetURLKey() string {
	return fmt.Sprintf("fa_quickratio_%v", q.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (q QuickRatioFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	q.Value = AssetRatioType(result)
	return q
}

// DebtEquityType represents the types of debt/equity filters
type DebtEquityType string

// Default filter constants
const (
	DEHigh         DebtEquityType = "high"
	DELow          DebtEquityType = "low"
	DEUnder1       DebtEquityType = "u1"
	DEUnder0point9 DebtEquityType = "u0.9"
	DEUnder0point8 DebtEquityType = "u0.8"
	DEUnder0point7 DebtEquityType = "u0.7"
	DEUnder0point6 DebtEquityType = "u0.6"
	DEUnder0point5 DebtEquityType = "u0.5"
	DEUnder0point4 DebtEquityType = "u0.4"
	DEUnder0point3 DebtEquityType = "u0.3"
	DEUnder0point2 DebtEquityType = "u0.2"
	DEUnder0point1 DebtEquityType = "u0.1"
	DEOver1        DebtEquityType = "o1"
	DEOver0point9  DebtEquityType = "o0.9"
	DEOver0point8  DebtEquityType = "o0.8"
	DEOver0point7  DebtEquityType = "o0.7"
	DEOver0point6  DebtEquityType = "o0.6"
	DEOver0point5  DebtEquityType = "o0.5"
	DEOver0point4  DebtEquityType = "o0.4"
	DEOver0point3  DebtEquityType = "o0.3"
	DEOver0point2  DebtEquityType = "o0.2"
	DEOver0point1  DebtEquityType = "o0.1"
)

// LTDebtEquityFilter is a filter type
type LTDebtEquityFilter struct {
	Value DebtEquityType
}

// GetURLKey returns the filter's url key
func (l LTDebtEquityFilter) GetURLKey() string {
	return fmt.Sprintf("fa_ltdebteq_%v", l.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (l LTDebtEquityFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	l.Value = DebtEquityType(result)
	return l
}

// DebtEquityFilter is a filter type
type DebtEquityFilter struct {
	Value DebtEquityType
}

// GetURLKey returns the filter's url key
func (d DebtEquityFilter) GetURLKey() string {
	return fmt.Sprintf("fa_debteq_%v", d.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (d DebtEquityFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	d.Value = DebtEquityType(result)
	return d
}

// GrossMarginType represents the types of gross margin filters
type GrossMarginType string

// Default filter constants
const (
	GMPositive    GrossMarginType = "pos"
	GMNegative    GrossMarginType = "neg"
	GMHigh        GrossMarginType = "high"
	GMUnder90     GrossMarginType = "u90"
	GMUnder80     GrossMarginType = "u80"
	GMUnder70     GrossMarginType = "u70"
	GMUnder60     GrossMarginType = "u60"
	GMUnder50     GrossMarginType = "u50"
	GMUnder45     GrossMarginType = "u45"
	GMUnder40     GrossMarginType = "u40"
	GMUnder35     GrossMarginType = "u35"
	GMUnder30     GrossMarginType = "u30"
	GMUnder25     GrossMarginType = "u25"
	GMUnder20     GrossMarginType = "u20"
	GMUnder15     GrossMarginType = "u15"
	GMUnder10     GrossMarginType = "u10"
	GMUnder5      GrossMarginType = "u5"
	GMUnder0      GrossMarginType = "u0"
	GMUnderNeg10  GrossMarginType = "u-10"
	GMUnderNeg20  GrossMarginType = "u-20"
	GMUnderNeg30  GrossMarginType = "u-30"
	GMUnderNeg50  GrossMarginType = "u-50"
	GMUnderNeg70  GrossMarginType = "u-70"
	GMUnderNeg100 GrossMarginType = "u-100"
	GMOver0       GrossMarginType = "o0"
	GMOver5       GrossMarginType = "o5"
	GMOver10      GrossMarginType = "o10"
	GMOver15      GrossMarginType = "o15"
	GMOver20      GrossMarginType = "o20"
	GMOver25      GrossMarginType = "o25"
	GMOver30      GrossMarginType = "o30"
	GMOver35      GrossMarginType = "o35"
	GMOver40      GrossMarginType = "o40"
	GMOver45      GrossMarginType = "o45"
	GMOver50      GrossMarginType = "o50"
	GMOver60      GrossMarginType = "o60"
	GMOver70      GrossMarginType = "o70"
	GMOver80      GrossMarginType = "o80"
	GMOver90      GrossMarginType = "o90"
)

// GrossMarginFilter is a filter type
type GrossMarginFilter struct {
	Value GrossMarginType
}

// GetURLKey returns the filter's url key
func (g GrossMarginFilter) GetURLKey() string {
	return fmt.Sprintf("fa_grossmargin_%v", g.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (g GrossMarginFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	g.Value = GrossMarginType(result)
	return g
}

// OperatingMarginType represents the types of operating margin filters
type OperatingMarginType string

// Default filter constants
const (
	OMPositive     OperatingMarginType = "pos"
	OMNegative     OperatingMarginType = "neg"
	OMVeryNegative OperatingMarginType = "veryneg"
	OMHigh         OperatingMarginType = "high"
	OMUnder90      OperatingMarginType = "u90"
	OMUnder80      OperatingMarginType = "u80"
	OMUnder70      OperatingMarginType = "u70"
	OMUnder60      OperatingMarginType = "u60"
	OMUnder50      OperatingMarginType = "u50"
	OMUnder45      OperatingMarginType = "u45"
	OMUnder40      OperatingMarginType = "u40"
	OMUnder35      OperatingMarginType = "u35"
	OMUnder30      OperatingMarginType = "u30"
	OMUnder25      OperatingMarginType = "u25"
	OMUnder20      OperatingMarginType = "u20"
	OMUnder15      OperatingMarginType = "u15"
	OMUnder10      OperatingMarginType = "u10"
	OMUnder5       OperatingMarginType = "u5"
	OMUnder0       OperatingMarginType = "u0"
	OMUnderNeg10   OperatingMarginType = "u-10"
	OMUnderNeg20   OperatingMarginType = "u-20"
	OMUnderNeg30   OperatingMarginType = "u-30"
	OMUnderNeg50   OperatingMarginType = "u-50"
	OMUnderNeg70   OperatingMarginType = "u-70"
	OMUnderNeg100  OperatingMarginType = "u-100"
	OMOver0        OperatingMarginType = "o0"
	OMOver5        OperatingMarginType = "o5"
	OMOver10       OperatingMarginType = "o10"
	OMOver15       OperatingMarginType = "o15"
	OMOver20       OperatingMarginType = "o20"
	OMOver25       OperatingMarginType = "o25"
	OMOver30       OperatingMarginType = "o30"
	OMOver35       OperatingMarginType = "o35"
	OMOver40       OperatingMarginType = "o40"
	OMOver45       OperatingMarginType = "o45"
	OMOver50       OperatingMarginType = "o50"
	OMOver60       OperatingMarginType = "o60"
	OMOver70       OperatingMarginType = "o70"
	OMOver80       OperatingMarginType = "o80"
	OMOver90       OperatingMarginType = "o90"
)

// OperatingMarginFilter is a filter type
type OperatingMarginFilter struct {
	Value OperatingMarginType
}

// GetURLKey returns the filter's url key
func (o OperatingMarginFilter) GetURLKey() string {
	return fmt.Sprintf("fa_opermargin_%v", o.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (o OperatingMarginFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	o.Value = OperatingMarginType(result)
	return o
}

// NetProfitMarginFilter is a filter type
type NetProfitMarginFilter struct {
	Value OperatingMarginType
}

// GetURLKey returns the filter's url key
func (n NetProfitMarginFilter) GetURLKey() string {
	return fmt.Sprintf("fa_netmargin_%v", n.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (n NetProfitMarginFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	n.Value = OperatingMarginType(result)
	return n
}

// PayoutRatioType represents the types of payout ratio filters
type PayoutRatioType string

// Default filter constants
const (
	PRNone     PayoutRatioType = "none"
	PRPositive PayoutRatioType = "pos"
	PRLow      PayoutRatioType = "low"
	PRHigh     PayoutRatioType = "high"
	PROver1    PayoutRatioType = "o10"
	PROver2    PayoutRatioType = "o20"
	PROver3    PayoutRatioType = "o30"
	PROver4    PayoutRatioType = "o40"
	PROver5    PayoutRatioType = "o50"
	PROver6    PayoutRatioType = "o60"
	PROver7    PayoutRatioType = "o70"
	PROver8    PayoutRatioType = "o80"
	PROver9    PayoutRatioType = "o90"
	PROver10   PayoutRatioType = "o100"
	PRUnder1   PayoutRatioType = "u10"
	PRUnder2   PayoutRatioType = "u20"
	PRUnder3   PayoutRatioType = "u30"
	PRUnder4   PayoutRatioType = "u40"
	PRUnder5   PayoutRatioType = "u50"
	PRUnder6   PayoutRatioType = "u60"
	PRUnder7   PayoutRatioType = "u70"
	PRUnder8   PayoutRatioType = "u80"
	PRUnder9   PayoutRatioType = "u90"
	PRUnder10  PayoutRatioType = "u100"
)

// PayoutRatioFilter is a filter type
type PayoutRatioFilter struct {
	Value PayoutRatioType
}

// GetURLKey returns the filter's url key
func (p PayoutRatioFilter) GetURLKey() string {
	return fmt.Sprintf("fa_payoutratio_%v", p.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (p PayoutRatioFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	p.Value = PayoutRatioType(result)
	return p
}

// InsiderOwnershipType represents the types of insider ownership filters
type InsiderOwnershipType string

// Default filter constants
const (
	InsdrOwnLow      InsiderOwnershipType = "low"
	InsdrOwnHigh     InsiderOwnershipType = "high"
	InsdrOwnVeryHigh InsiderOwnershipType = "veryhigh"
	InsdrOwnOver10   InsiderOwnershipType = "o10"
	InsdrOwnOver20   InsiderOwnershipType = "o20"
	InsdrOwnOver30   InsiderOwnershipType = "o30"
	InsdrOwnOver40   InsiderOwnershipType = "o40"
	InsdrOwnOver50   InsiderOwnershipType = "o50"
	InsdrOwnOver60   InsiderOwnershipType = "o60"
	InsdrOwnOver70   InsiderOwnershipType = "o70"
	InsdrOwnOver80   InsiderOwnershipType = "o80"
	InsdrOwnOver90   InsiderOwnershipType = "o90"
)

// InsiderOwnershipFilter is a filter type
type InsiderOwnershipFilter struct {
	Value InsiderOwnershipType
}

// GetURLKey returns the filter's url key
func (i InsiderOwnershipFilter) GetURLKey() string {
	return fmt.Sprintf("sh_insiderown_%v", i.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (i InsiderOwnershipFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	i.Value = InsiderOwnershipType(result)
	return i
}

// InsiderTransactionsType represents the types of insider transactions filters
type InsiderTransactionsType string

// Default filter constants
const (
	InsdrTransVeryNeg    InsiderTransactionsType = "veryneg"
	InsdrTransNeg        InsiderTransactionsType = "neg"
	InsdrTransPos        InsiderTransactionsType = "pos"
	InsdrTransVeryPos    InsiderTransactionsType = "verypos"
	InsdrTransOver5      InsiderTransactionsType = "o5"
	InsdrTransOver10     InsiderTransactionsType = "o10"
	InsdrTransOver15     InsiderTransactionsType = "o15"
	InsdrTransOver20     InsiderTransactionsType = "o20"
	InsdrTransOver25     InsiderTransactionsType = "o25"
	InsdrTransOver30     InsiderTransactionsType = "o30"
	InsdrTransOver35     InsiderTransactionsType = "o35"
	InsdrTransOver40     InsiderTransactionsType = "o40"
	InsdrTransOver45     InsiderTransactionsType = "o45"
	InsdrTransOver50     InsiderTransactionsType = "o50"
	InsdrTransOver60     InsiderTransactionsType = "o60"
	InsdrTransOver70     InsiderTransactionsType = "o70"
	InsdrTransOver80     InsiderTransactionsType = "o80"
	InsdrTransOver90     InsiderTransactionsType = "o90"
	InsdrTransUnderNeg5  InsiderTransactionsType = "u-5"
	InsdrTransUnderNeg10 InsiderTransactionsType = "u-10"
	InsdrTransUnderNeg15 InsiderTransactionsType = "u-15"
	InsdrTransUnderNeg20 InsiderTransactionsType = "u-20"
	InsdrTransUnderNeg25 InsiderTransactionsType = "u-25"
	InsdrTransUnderNeg30 InsiderTransactionsType = "u-30"
	InsdrTransUnderNeg35 InsiderTransactionsType = "u-35"
	InsdrTransUnderNeg40 InsiderTransactionsType = "u-40"
	InsdrTransUnderNeg45 InsiderTransactionsType = "u-45"
	InsdrTransUnderNeg50 InsiderTransactionsType = "u-50"
	InsdrTransUnderNeg60 InsiderTransactionsType = "u-60"
	InsdrTransUnderNeg70 InsiderTransactionsType = "u-70"
	InsdrTransUnderNeg80 InsiderTransactionsType = "u-80"
	InsdrTransUnderNeg90 InsiderTransactionsType = "u-90"
)

// InsiderTransactionsFilter is a filter type
type InsiderTransactionsFilter struct {
	Value InsiderTransactionsType
}

// GetURLKey returns the filter's url key
func (i InsiderTransactionsFilter) GetURLKey() string {
	return fmt.Sprintf("sh_insidertrans_%v", i.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (i InsiderTransactionsFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	i.Value = InsiderTransactionsType(result)
	return i
}

// InstitutionalOwnershipType represents the types of institutional ownership filters
type InstitutionalOwnershipType string

// Default filter constants
const (
	InstOwnLow     InstitutionalOwnershipType = "low"
	InstOwnHigh    InstitutionalOwnershipType = "high"
	InstOwnOver10  InstitutionalOwnershipType = "o10"
	InstOwnOver20  InstitutionalOwnershipType = "o20"
	InstOwnOver30  InstitutionalOwnershipType = "o30"
	InstOwnOver40  InstitutionalOwnershipType = "o40"
	InstOwnOver50  InstitutionalOwnershipType = "o50"
	InstOwnOver60  InstitutionalOwnershipType = "o60"
	InstOwnOver70  InstitutionalOwnershipType = "o70"
	InstOwnOver80  InstitutionalOwnershipType = "o80"
	InstOwnOver90  InstitutionalOwnershipType = "o90"
	InstOwnUnder10 InstitutionalOwnershipType = "u10"
	InstOwnUnder20 InstitutionalOwnershipType = "u20"
	InstOwnUnder30 InstitutionalOwnershipType = "u30"
	InstOwnUnder40 InstitutionalOwnershipType = "u40"
	InstOwnUnder50 InstitutionalOwnershipType = "u50"
	InstOwnUnder60 InstitutionalOwnershipType = "u60"
	InstOwnUnder70 InstitutionalOwnershipType = "u70"
	InstOwnUnder80 InstitutionalOwnershipType = "u80"
	InstOwnUnder90 InstitutionalOwnershipType = "u90"
)

// InstitutionalOwnershipFilter is a filter type
type InstitutionalOwnershipFilter struct {
	Value InstitutionalOwnershipType
}

// GetURLKey returns the filter's url key
func (i InstitutionalOwnershipFilter) GetURLKey() string {
	return fmt.Sprintf("sh_instown_%v", i.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (i InstitutionalOwnershipFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	i.Value = InstitutionalOwnershipType(result)
	return i
}

// InstitutionalTransactionsType represents the types of institutional transactions filters
type InstitutionalTransactionsType string

// Default filter constants
const (
	InstTransVeryNeg    InstitutionalTransactionsType = "veryneg"
	InstTransNeg        InstitutionalTransactionsType = "neg"
	InstTransPos        InstitutionalTransactionsType = "pos"
	InstTransVeryPos    InstitutionalTransactionsType = "verypos"
	InstTransOver5      InstitutionalTransactionsType = "o5"
	InstTransOver10     InstitutionalTransactionsType = "o10"
	InstTransOver15     InstitutionalTransactionsType = "o15"
	InstTransOver20     InstitutionalTransactionsType = "o20"
	InstTransOver25     InstitutionalTransactionsType = "o25"
	InstTransOver30     InstitutionalTransactionsType = "o30"
	InstTransOver35     InstitutionalTransactionsType = "o35"
	InstTransOver40     InstitutionalTransactionsType = "o40"
	InstTransOver45     InstitutionalTransactionsType = "o45"
	InstTransOver50     InstitutionalTransactionsType = "o50"
	InstTransUnderNeg5  InstitutionalTransactionsType = "u-5"
	InstTransUnderNeg10 InstitutionalTransactionsType = "u-10"
	InstTransUnderNeg15 InstitutionalTransactionsType = "u-15"
	InstTransUnderNeg20 InstitutionalTransactionsType = "u-20"
	InstTransUnderNeg25 InstitutionalTransactionsType = "u-25"
	InstTransUnderNeg30 InstitutionalTransactionsType = "u-30"
	InstTransUnderNeg35 InstitutionalTransactionsType = "u-35"
	InstTransUnderNeg40 InstitutionalTransactionsType = "u-40"
	InstTransUnderNeg45 InstitutionalTransactionsType = "u-45"
	InstTransUnderNeg50 InstitutionalTransactionsType = "u-50"
)

// InstitutionalTransactionsFilter is a filter type
type InstitutionalTransactionsFilter struct {
	Value InstitutionalTransactionsType
}

// GetURLKey returns the filter's url key
func (i InstitutionalTransactionsFilter) GetURLKey() string {
	return fmt.Sprintf("sh_insttrans_%v", i.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (i InstitutionalTransactionsFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	i.Value = InstitutionalTransactionsType(result)
	return i
}

/***************************************************************************
*****	TECHNICAL   ********************************************************
***************************************************************************/

// PerformanceType represents the types of performance filters
type PerformanceType string

// Default filter constants
const (
	PerfTodayUp                PerformanceType = "dup"
	PerfTodayDown              PerformanceType = "down"
	PerfTodayDown15Percent     PerformanceType = "d15u"
	PerfTodayDown10Percent     PerformanceType = "d10u"
	PerfTodayDown5Percent      PerformanceType = "d5u"
	PerfTodayUp5Percent        PerformanceType = "d5o"
	PerfTodayUp10Percent       PerformanceType = "d10o"
	PerfTodayUp15Percent       PerformanceType = "d15o"
	PerfWeekDown30Percent      PerformanceType = "1w30u"
	PerfWeekDown20Percent      PerformanceType = "1w20u"
	PerfWeekDown10Percent      PerformanceType = "1w10u"
	PerfWeekDown               PerformanceType = "1wdown"
	PerfWeekUp                 PerformanceType = "1wup"
	PerfWeekUp10Percent        PerformanceType = "1w10o"
	PerfWeekUp20Percent        PerformanceType = "1w20o"
	PerfWeekUp30Percent        PerformanceType = "1w30o"
	PerfMonthDown50Percent     PerformanceType = "4w50u"
	PerfMonthDown30Percent     PerformanceType = "4w30u"
	PerfMonthDown20Percent     PerformanceType = "4w20u"
	PerfMonthDown10Percent     PerformanceType = "4w10u"
	PerfMonthDown              PerformanceType = "4wdown"
	PerfMonthUp                PerformanceType = "4wup"
	PerfMonthUp10Percent       PerformanceType = "4w10o"
	PerfMonthUp20Percent       PerformanceType = "4w20o"
	PerfMonthUp30Percent       PerformanceType = "4w30o"
	PerfMonthUp50Percent       PerformanceType = "4w50o"
	PerfQuarterDown50Percent   PerformanceType = "13w50u"
	PerfQuarterDown30Percent   PerformanceType = "13w30u"
	PerfQuarterDown20Percent   PerformanceType = "13w20u"
	PerfQuarterDown10Percent   PerformanceType = "13w10u"
	PerfQuarterDown            PerformanceType = "13wdown"
	PerfQuarterUp              PerformanceType = "13wup"
	PerfQuarterUp10Percent     PerformanceType = "13w10o"
	PerfQuarterUp20Percent     PerformanceType = "13w20o"
	PerfQuarterUp30Percent     PerformanceType = "13w30o"
	PerfQuarterUp50Percent     PerformanceType = "13w50o"
	PerfSixMonthsDown75Percent PerformanceType = "26w75u"
	PerfSixMonthsDown50Percent PerformanceType = "26w50u"
	PerfSixMonthsDown30Percent PerformanceType = "26w30u"
	PerfSixMonthsDown20Percent PerformanceType = "26w20u"
	PerfSixMonthsDown10Percent PerformanceType = "26w10u"
	PerfSixMonthsDown          PerformanceType = "26wdown"
	PerfSixMonthsUp            PerformanceType = "26wup"
	PerfSixMonthsUp10Percent   PerformanceType = "26w10o"
	PerfSixMonthsUp20Percent   PerformanceType = "26w20o"
	PerfSixMonthsUp30Percent   PerformanceType = "26w30o"
	PerfSixMonthsUp50Percent   PerformanceType = "26w50o"
	PerfSixMonthsUp100Percent  PerformanceType = "26w100o"
	PerfYearDown75Percent      PerformanceType = "52w75u"
	PerfYearDown50Percent      PerformanceType = "52w50u"
	PerfYearDown30Percent      PerformanceType = "52w30u"
	PerfYearDown20Percent      PerformanceType = "52w20u"
	PerfYearDown10Percent      PerformanceType = "52w10u"
	PerfYearDown               PerformanceType = "52wdown"
	PerfYearUp                 PerformanceType = "52wup"
	PerfYearUp10Percent        PerformanceType = "52w10o"
	PerfYearUp20Percent        PerformanceType = "52w20o"
	PerfYearUp30Percent        PerformanceType = "52w30o"
	PerfYearUp50Percent        PerformanceType = "52w50o"
	PerfYearUp100Percent       PerformanceType = "52w100o"
	PerfYearUp200Percent       PerformanceType = "52w200o"
	PerfYearUp300Percent       PerformanceType = "52w300o"
	PerfYearUp500Percent       PerformanceType = "52w500o"
	PerfYTDDown75Percent       PerformanceType = "ytd75u"
	PerfYTDDown50Percent       PerformanceType = "ytd50u"
	PerfYTDDown30Percent       PerformanceType = "ytd30u"
	PerfYTDDown20Percent       PerformanceType = "ytd20u"
	PerfYTDDown10Percent       PerformanceType = "ytd10u"
	PerfYTDDown5Percent        PerformanceType = "ytd5u"
	PerfYTDDown                PerformanceType = "ytddown"
	PerfYTDUp                  PerformanceType = "ytdup"
	PerfYTDUp5Percent          PerformanceType = "ytd5o"
	PerfYTDUp10Percent         PerformanceType = "ytd10o"
	PerfYTDUp20Percent         PerformanceType = "ytd20o"
	PerfYTDUp30Percent         PerformanceType = "ytd30o"
	PerfYTDUp50Percent         PerformanceType = "ytd50o"
	PerfYTDUp100Percent        PerformanceType = "ytd100o"
)

// PerformanceFilter is a filter type
type PerformanceFilter struct {
	Value PerformanceType
}

// GetURLKey returns the filter's url key
func (p PerformanceFilter) GetURLKey() string {
	return fmt.Sprintf("ta_perf_%v", p.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (p PerformanceFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	p.Value = PerformanceType(result)
	return p
}

// Performance2Filter is a filter type
type Performance2Filter struct {
	Value PerformanceType
}

// GetURLKey returns the filter's url key
func (p Performance2Filter) GetURLKey() string {
	return fmt.Sprintf("ta_perf2_%v", p.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (p Performance2Filter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	p.Value = PerformanceType(result)
	return p
}

// VolatilityType represents the types of volatility filters
type VolatilityType string

// Default filter constants
const (
	VolWeekOver3   VolatilityType = "wo3"
	VolWeekOver4   VolatilityType = "wo4"
	VolWeekOver5   VolatilityType = "wo5"
	VolWeekOver6   VolatilityType = "wo6"
	VolWeekOver7   VolatilityType = "wo7"
	VolWeekOver8   VolatilityType = "wo8"
	VolWeekOver9   VolatilityType = "wo9"
	VolWeekOver10  VolatilityType = "wo10"
	VolWeekOver12  VolatilityType = "wo12"
	VolWeekOver15  VolatilityType = "wo15"
	VolMonthOver2  VolatilityType = "mo2"
	VolMonthOver3  VolatilityType = "mo3"
	VolMonthOver4  VolatilityType = "mo4"
	VolMonthOver5  VolatilityType = "mo5"
	VolMonthOver6  VolatilityType = "mo6"
	VolMonthOver7  VolatilityType = "mo7"
	VolMonthOver8  VolatilityType = "mo8"
	VolMonthOver9  VolatilityType = "mo9"
	VolMonthOver10 VolatilityType = "mo10"
	VolMonthOver12 VolatilityType = "mo12"
	VolMonthOver15 VolatilityType = "mo15"
)

// VolatilityFilter is a filter type
type VolatilityFilter struct {
	Value VolatilityType
}

// GetURLKey returns the filter's url key
func (v VolatilityFilter) GetURLKey() string {
	return fmt.Sprintf("ta_volatility_%v", v.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (v VolatilityFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	v.Value = VolatilityType(result)
	return v
}

// RSIType represents the types of RSI 14-day filters
type RSIType string

// Default filter constants
const (
	Overbought90         RSIType = "ob90"
	Overbought80         RSIType = "ob80"
	Overbought70         RSIType = "ob70"
	Overbought60         RSIType = "ob60"
	Oversold40           RSIType = "os40"
	Oversold30           RSIType = "os30"
	Oversold20           RSIType = "os20"
	Oversold10           RSIType = "os10"
	NotOverboughtUnder60 RSIType = "nob60"
	NotOverboughtUnder50 RSIType = "nob50"
	NotOversoldOver50    RSIType = "nos50"
	NotOversoldOver40    RSIType = "nos40"
)

// RSIFilter is a filter type
type RSIFilter struct {
	Value RSIType
}

// GetURLKey returns the filter's url key
func (r RSIFilter) GetURLKey() string {
	return fmt.Sprintf("ta_rsi_%v", r.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (r RSIFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	r.Value = RSIType(result)
	return r
}

// GapType represents the types of gap filters
type GapType string

// Default filter constants
const (
	GapUp            GapType = "u"
	GapUp0Percent    GapType = "u0"
	GapUp1Percent    GapType = "u1"
	GapUp2Percent    GapType = "u2"
	GapUp3Percent    GapType = "u3"
	GapUp4Percent    GapType = "u4"
	GapUp5Percent    GapType = "u5"
	GapUp6Percent    GapType = "u6"
	GapUp7Percent    GapType = "u7"
	GapUp8Percent    GapType = "u8"
	GapUp9Percent    GapType = "u9"
	GapUp10Percent   GapType = "u10"
	GapUp15Percent   GapType = "u15"
	GapUp20Percent   GapType = "u20"
	GapDown          GapType = "d"
	GapDown0Percent  GapType = "d0"
	GapDown1Percent  GapType = "d1"
	GapDown2Percent  GapType = "d2"
	GapDown3Percent  GapType = "d3"
	GapDown4Percent  GapType = "d4"
	GapDown5Percent  GapType = "d5"
	GapDown6Percent  GapType = "d6"
	GapDown7Percent  GapType = "d7"
	GapDown8Percent  GapType = "d8"
	GapDown9Percent  GapType = "d9"
	GapDown10Percent GapType = "d10"
	GapDown15Percent GapType = "d15"
	GapDown20Percent GapType = "d20"
)

// GapFilter is a filter type
type GapFilter struct {
	Value GapType
}

// GetURLKey returns the filter's url key
func (g GapFilter) GetURLKey() string {
	return fmt.Sprintf("ta_gap_%v", g.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (g GapFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	g.Value = GapType(result)
	return g
}

// SMA20Type represents the types of 20-Day Simple Moving Average filters
type SMA20Type string

// Default filter constants
const (
	Mavg20PriceBelowSMA20          SMA20Type = "pb"
	Mavg20Price10PercentBelowSMA20 SMA20Type = "pb10"
	Mavg20Price20PercentBelowSMA20 SMA20Type = "pb20"
	Mavg20Price30PercentBelowSMA20 SMA20Type = "pb30"
	Mavg20Price40PercentBelowSMA20 SMA20Type = "pb40"
	Mavg20Price50PercentBelowSMA20 SMA20Type = "pb50"
	Mavg20PriceAboveSMA20          SMA20Type = "pa"
	Mavg20Price10PercentAboveSMA20 SMA20Type = "pa10"
	Mavg20Price20PercentAboveSMA20 SMA20Type = "pa20"
	Mavg20Price30PercentAboveSMA20 SMA20Type = "pa30"
	Mavg20Price40PercentAboveSMA20 SMA20Type = "pa40"
	Mavg20Price50PercentAboveSMA20 SMA20Type = "pa50"
	Mavg20PriceCrossedSMA20        SMA20Type = "pc"
	Mavg20PriceCrossedSMA20Above   SMA20Type = "pca"
	Mavg20PriceCrossedSMA20Below   SMA20Type = "pcb"
	Mavg20SMA20CrossedSMA50        SMA20Type = "cross50"
	Mavg20SMA20CrossedSMA50Above   SMA20Type = "cross50a"
	Mavg20SMA20CrossedSMA50Below   SMA20Type = "cross50b"
	Mavg20SMA20CrossedSMA200       SMA20Type = "cross200"
	Mavg20SMA20CrossedSMA200Above  SMA20Type = "cross200a"
	Mavg20SMA20CrossedSMA200Below  SMA20Type = "cross200b"
	Mavg20SMA20AboveSMA50          SMA20Type = "sa50"
	Mavg20SMA20BelowSMA50          SMA20Type = "sb50"
	Mavg20SMA20AboveSMA200         SMA20Type = "sa200"
	Mavg20SMA20BelowSMA200         SMA20Type = "sb200"
)

// SMA20Filter is a filter type
type SMA20Filter struct {
	Value SMA20Type
}

// GetURLKey returns the filter's url key
func (s SMA20Filter) GetURLKey() string {
	return fmt.Sprintf("ta_sma20_%v", s.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (s SMA20Filter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	s.Value = SMA20Type(result)
	return s
}

// SMA50Type represents the types of 50-Day Simple Moving Average filters
type SMA50Type string

// Default filter constants
const (
	Mavg50PriceBelowSMA50          SMA50Type = "pb"
	Mavg50Price10PercentBelowSMA50 SMA50Type = "pb10"
	Mavg50Price20PercentBelowSMA50 SMA50Type = "pb20"
	Mavg50Price30PercentBelowSMA50 SMA50Type = "pb30"
	Mavg50Price40PercentBelowSMA50 SMA50Type = "pb40"
	Mavg50Price50PercentBelowSMA50 SMA50Type = "pb50"
	Mavg50PriceAboveSMA50          SMA50Type = "pa"
	Mavg50Price10PercentAboveSMA50 SMA50Type = "pa10"
	Mavg50Price20PercentAboveSMA50 SMA50Type = "pa20"
	Mavg50Price30PercentAboveSMA50 SMA50Type = "pa30"
	Mavg50Price40PercentAboveSMA50 SMA50Type = "pa40"
	Mavg50Price50PercentAboveSMA50 SMA50Type = "pa50"
	Mavg50PriceCrossedSMA50        SMA50Type = "pc"
	Mavg50PriceCrossedSMA50Above   SMA50Type = "pca"
	Mavg50PriceCrossedSMA50Below   SMA50Type = "pcb"
	Mavg50SMA50CrossedSMA20        SMA50Type = "cross20"
	Mavg50SMA50CrossedSMA20Above   SMA50Type = "cross20a"
	Mavg50SMA50CrossedSMA20Below   SMA50Type = "cross20b"
	Mavg50SMA50CrossedSMA200       SMA50Type = "cross200"
	Mavg50SMA50CrossedSMA200Above  SMA50Type = "cross200a"
	Mavg50SMA50CrossedSMA200Below  SMA50Type = "cross200b"
	Mavg50SMA50AboveSMA20          SMA50Type = "sa20"
	Mavg50SMA50BelowSMA20          SMA50Type = "sb20"
	Mavg50SMA50AboveSMA200         SMA50Type = "sa200"
	Mavg50SMA50BelowSMA200         SMA50Type = "sb200"
)

// SMA50Filter is a filter type
type SMA50Filter struct {
	Value SMA50Type
}

// GetURLKey returns the filter's url key
func (s SMA50Filter) GetURLKey() string {
	return fmt.Sprintf("ta_sma50_%v", s.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (s SMA50Filter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	s.Value = SMA50Type(result)
	return s
}

// SMA200Type represents the types of 200-Day Simple Moving Average filters
type SMA200Type string

// Default filter constants
const (
	Mavg200PriceBelowSMA200           SMA200Type = "pb"
	Mavg200Price10PercentBelowSMA200  SMA200Type = "pb10"
	Mavg200Price20PercentBelowSMA200  SMA200Type = "pb20"
	Mavg200Price30PercentBelowSMA200  SMA200Type = "pb30"
	Mavg200Price40PercentBelowSMA200  SMA200Type = "pb40"
	Mavg200Price50PercentBelowSMA200  SMA200Type = "pb50"
	Mavg200Price60PercentBelowSMA200  SMA200Type = "pb60"
	Mavg200Price70PercentBelowSMA200  SMA200Type = "pb70"
	Mavg200Price80PercentBelowSMA200  SMA200Type = "pb80"
	Mavg200Price90PercentBelowSMA200  SMA200Type = "pb90"
	Mavg200PriceAboveSMA200           SMA200Type = "pa"
	Mavg200Price10PercentAboveSMA200  SMA200Type = "pa10"
	Mavg200Price20PercentAboveSMA200  SMA200Type = "pa20"
	Mavg200Price30PercentAboveSMA200  SMA200Type = "pa30"
	Mavg200Price40PercentAboveSMA200  SMA200Type = "pa40"
	Mavg200Price50PercentAboveSMA200  SMA200Type = "pa50"
	Mavg200Price60PercentAboveSMA200  SMA200Type = "pa60"
	Mavg200Price70PercentAboveSMA200  SMA200Type = "pa70"
	Mavg200Price80PercentAboveSMA200  SMA200Type = "pa80"
	Mavg200Price90PercentAboveSMA200  SMA200Type = "pa90"
	Mavg200Price100PercentAboveSMA200 SMA200Type = "pa100"
	Mavg200PriceCrossedSMA200         SMA200Type = "pc"
	Mavg200PriceCrossedSMA200Above    SMA200Type = "pca"
	Mavg200PriceCrossedSMA200Below    SMA200Type = "pcb"
	Mavg200SMA200CrossedSMA20         SMA200Type = "cross20"
	Mavg200SMA200CrossedSMA20Above    SMA200Type = "cross20a"
	Mavg200SMA200CrossedSMA20Below    SMA200Type = "cross20b"
	Mavg200SMA200CrossedSMA50         SMA200Type = "cross50"
	Mavg200SMA200CrossedSMA50Above    SMA200Type = "cross50a"
	Mavg200SMA200CrossedSMA50Below    SMA200Type = "cross50b"
	Mavg200SMA200AboveSMA20           SMA200Type = "sa20"
	Mavg200SMA200BelowSMA20           SMA200Type = "sb20"
	Mavg200SMA200AboveSMA50           SMA200Type = "sa50"
	Mavg200SMA200BelowSMA50           SMA200Type = "sb50"
)

// SMA200Filter is a filter type
type SMA200Filter struct {
	Value SMA200Type
}

// GetURLKey returns the filter's url key
func (s SMA200Filter) GetURLKey() string {
	return fmt.Sprintf("ta_sma200_%v", s.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (s SMA200Filter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	s.Value = SMA200Type(result)
	return s
}

// ChangeType represents the types of change (%) filters
type ChangeType string

// Default filter constants
const (
	ChangeUp            ChangeType = "u"
	ChangeUp1Percent    ChangeType = "u1"
	ChangeUp2Percent    ChangeType = "u2"
	ChangeUp3Percent    ChangeType = "u3"
	ChangeUp4Percent    ChangeType = "u4"
	ChangeUp5Percent    ChangeType = "u5"
	ChangeUp6Percent    ChangeType = "u6"
	ChangeUp7Percent    ChangeType = "u7"
	ChangeUp8Percent    ChangeType = "u8"
	ChangeUp9Percent    ChangeType = "u9"
	ChangeUp10Percent   ChangeType = "u10"
	ChangeUp15Percent   ChangeType = "u15"
	ChangeUp20Percent   ChangeType = "u20"
	ChangeDown          ChangeType = "d"
	ChangeDown1Percent  ChangeType = "d1"
	ChangeDown2Percent  ChangeType = "d2"
	ChangeDown3Percent  ChangeType = "d3"
	ChangeDown4Percent  ChangeType = "d4"
	ChangeDown5Percent  ChangeType = "d5"
	ChangeDown6Percent  ChangeType = "d6"
	ChangeDown7Percent  ChangeType = "d7"
	ChangeDown8Percent  ChangeType = "d8"
	ChangeDown9Percent  ChangeType = "d9"
	ChangeDown10Percent ChangeType = "d10"
	ChangeDown15Percent ChangeType = "d15"
	ChangeDown20Percent ChangeType = "d20"
)

// ChangeFilter is a filter type
type ChangeFilter struct {
	Value ChangeType
}

// GetURLKey returns the filter's url key
func (c ChangeFilter) GetURLKey() string {
	return fmt.Sprintf("ta_change_%v", c.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (c ChangeFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	c.Value = ChangeType(result)
	return c
}

// ChangeFromOpenFilter is a filter type
type ChangeFromOpenFilter struct {
	Value ChangeType
}

// GetURLKey returns the filter's url key
func (c ChangeFromOpenFilter) GetURLKey() string {
	return fmt.Sprintf("ta_changeopen_%v", c.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (c ChangeFromOpenFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	c.Value = ChangeType(result)
	return c
}

// HighLow20DayType represents the types of 20-day, 50-day high/low filters
type HighLow20DayType string

// Default filter constants
const (
	DHLNewHigh                   HighLow20DayType = "nh"
	DHLNewLow                    HighLow20DayType = "nl"
	DHLAtLeast5PercentBelowHigh  HighLow20DayType = "b5h"
	DHLAtLeast10PercentBelowHigh HighLow20DayType = "b10h"
	DHLAtLeast15PercentBelowHigh HighLow20DayType = "b15h"
	DHLAtLeast20PercentBelowHigh HighLow20DayType = "b20h"
	DHLAtLeast30PercentBelowHigh HighLow20DayType = "b30h"
	DHLAtLeast40PercentBelowHigh HighLow20DayType = "b40h"
	DHLAtLeast50PercentBelowHigh HighLow20DayType = "b50h"
	DHL0to3PercentBelowHigh      HighLow20DayType = "b0to3h"
	DHL0to5PercentBelowHigh      HighLow20DayType = "b0to5h"
	DHL0to10PercentBelowHigh     HighLow20DayType = "b0to10h"
	DHLAtLeast5PercentAboveLow   HighLow20DayType = "a5h"
	DHLAtLeast10PercentAboveLow  HighLow20DayType = "a10h"
	DHLAtLeast15PercentAboveLow  HighLow20DayType = "a15h"
	DHLAtLeast20PercentAboveLow  HighLow20DayType = "a20h"
	DHLAtLeast30PercentAboveLow  HighLow20DayType = "a30h"
	DHLAtLeast40PercentAboveLow  HighLow20DayType = "a40h"
	DHLAtLeast50PercentAboveLow  HighLow20DayType = "a50h"
	DHL0to3PercentAboveLow       HighLow20DayType = "a0to3h"
	DHL0to5PercentAboveLow       HighLow20DayType = "a0to5h"
	DHL0to10PercentAboveLow      HighLow20DayType = "a0to10h"
)

// HighLow20DayFilter is a filter type
type HighLow20DayFilter struct {
	Value HighLow20DayType
}

// GetURLKey returns the filter's url key
func (h HighLow20DayFilter) GetURLKey() string {
	return fmt.Sprintf("ta_highlow20d_%v", h.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (h HighLow20DayFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	h.Value = HighLow20DayType(result)
	return h
}

// HighLow50DayFilter is a filter type
type HighLow50DayFilter struct {
	Value HighLow20DayType
}

// GetURLKey returns the filter's url key
func (h HighLow50DayFilter) GetURLKey() string {
	return fmt.Sprintf("ta_highlow50d_%v", h.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (h HighLow50DayFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	h.Value = HighLow20DayType(result)
	return h
}

// HighLow52WeekType represents the types of 52-week high/low filters
type HighLow52WeekType string

// Default filter constants
const (
	WHLNewHigh                   HighLow52WeekType = "nh"
	WHLNewLow                    HighLow52WeekType = "nl"
	WHLAtLeast5PercentBelowHigh  HighLow52WeekType = "b5h"
	WHLAtLeast10PercentBelowHigh HighLow52WeekType = "b10h"
	WHLAtLeast15PercentBelowHigh HighLow52WeekType = "b15h"
	WHLAtLeast20PercentBelowHigh HighLow52WeekType = "b20h"
	WHLAtLeast30PercentBelowHigh HighLow52WeekType = "b30h"
	WHLAtLeast40PercentBelowHigh HighLow52WeekType = "b40h"
	WHLAtLeast50PercentBelowHigh HighLow52WeekType = "b50h"
	WHLAtLeast60PercentBelowHigh HighLow52WeekType = "b60h"
	WHLAtLeast70PercentBelowHigh HighLow52WeekType = "b70h"
	WHLAtLeast80PercentBelowHigh HighLow52WeekType = "b80h"
	WHLAtLeast90PercentBelowHigh HighLow52WeekType = "b90h"
	WHL0to3PercentBelowHigh      HighLow52WeekType = "b0to3h"
	WHL0to5PercentBelowHigh      HighLow52WeekType = "b0to5h"
	WHL0to10PercentBelowHigh     HighLow52WeekType = "b0to10h"
	WHLAtLeast5PercentAboveLow   HighLow52WeekType = "a5h"
	WHLAtLeast10PercentAboveLow  HighLow52WeekType = "a10h"
	WHLAtLeast15PercentAboveLow  HighLow52WeekType = "a15h"
	WHLAtLeast20PercentAboveLow  HighLow52WeekType = "a20h"
	WHLAtLeast30PercentAboveLow  HighLow52WeekType = "a30h"
	WHLAtLeast40PercentAboveLow  HighLow52WeekType = "a40h"
	WHLAtLeast50PercentAboveLow  HighLow52WeekType = "a50h"
	WHLAtLeast60PercentAboveLow  HighLow52WeekType = "a60h"
	WHLAtLeast70PercentAboveLow  HighLow52WeekType = "a70h"
	WHLAtLeast80PercentAboveLow  HighLow52WeekType = "a80h"
	WHLAtLeast90PercentAboveLow  HighLow52WeekType = "a90h"
	WHLAtLeast100PercentAboveLow HighLow52WeekType = "a100h"
	WHLAtLeast120PercentAboveLow HighLow52WeekType = "a120h"
	WHLAtLeast150PercentAboveLow HighLow52WeekType = "a150h"
	WHLAtLeast200PercentAboveLow HighLow52WeekType = "a200h"
	WHLAtLeast300PercentAboveLow HighLow52WeekType = "a300h"
	WHLAtLeast500PercentAboveLow HighLow52WeekType = "a500h"
	WHL0to3PercentAboveLow       HighLow52WeekType = "a0to3h"
	WHL0to5PercentAboveLow       HighLow52WeekType = "a0to5h"
	WHL0to10PercentAboveLow      HighLow52WeekType = "a0to10h"
)

// HighLow52WeekFilter is a filter type
type HighLow52WeekFilter struct {
	Value HighLow52WeekType
}

// GetURLKey returns the filter's url key
func (h HighLow52WeekFilter) GetURLKey() string {
	return fmt.Sprintf("ta_highlow52w_%v", h.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (h HighLow52WeekFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	h.Value = HighLow52WeekType(result)
	return h
}

// PatternType represents the types of technical pattern filters
type PatternType string

// Default filter constants
const (
	PatternHorizontalSR             PatternType = "horizontal"
	PatternHorizontalSRStrong       PatternType = "horizontal2"
	PatternTLResistance             PatternType = "tlresistance"
	PatternTLResistanceStrong       PatternType = "tlresistance2"
	PatternTLSupport                PatternType = "tlsupport"
	PatternTLSupportStrong          PatternType = "tlsupport2"
	PatternWedgeUp                  PatternType = "wedgeup"
	PatternWedgeUpStrong            PatternType = "wedgeup2"
	PatternWedgeDown                PatternType = "wedgedown"
	PatternWedgeDownStrong          PatternType = "wedgedown2"
	PatternTriangleAscending        PatternType = "wedgeresistance"
	PatternTriangleAscendingStrong  PatternType = "wedgeresistance2"
	PatternTriangleDescending       PatternType = "wedgesupport"
	PatternTriangleDescendingStrong PatternType = "wedgesupport2"
	PatternWedge                    PatternType = "wedge"
	PatternWedgeStrong              PatternType = "wedge2"
	PatternChannelUp                PatternType = "channelup"
	PatternChannelUpStrong          PatternType = "channelup2"
	PatternChannelDown              PatternType = "channeldown"
	PatternChannelDownStrong        PatternType = "channeldown2"
	PatternChannel                  PatternType = "channel"
	PatternChannelStrong            PatternType = "channel2"
	PatternDoubleTop                PatternType = "doubletop"
	PatternDoubleBottom             PatternType = "doublebottom"
	PatternMultipleTop              PatternType = "multipletop"
	PatternMultipleBottom           PatternType = "multiplebottom"
	PatternHeadAndShoulders         PatternType = "headandshoulders"
	PatternHeadAndShouldersInverse  PatternType = "headandshouldersinv"
)

// PatternFilter is a filter type
type PatternFilter struct {
	Value PatternType
}

// GetURLKey returns the filter's url key
func (p PatternFilter) GetURLKey() string {
	return fmt.Sprintf("ta_pattern_%v", p.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (p PatternFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	p.Value = PatternType(result)
	return p
}

// CandlestickType represents the types of candlestick filters
type CandlestickType string

// Default filter constants
const (
	CSLongLowerShadow  CandlestickType = "lls"
	CSLongUpperShadow  CandlestickType = "lus"
	CSHammer           CandlestickType = "h"
	CSInvertedHammer   CandlestickType = "ih"
	CSSpinningTopWhite CandlestickType = "stw"
	CSSpinningTopBlack CandlestickType = "stb"
	CSDoji             CandlestickType = "d"
	CSDragonflyDoji    CandlestickType = "dd"
	CSGravestoneDoji   CandlestickType = "gd"
	CSMarubozuWhite    CandlestickType = "mw"
	CSMarubozuBlack    CandlestickType = "mb"
)

// CandlestickFilter is a filter type
type CandlestickFilter struct {
	Value CandlestickType
}

// GetURLKey returns the filter's url key
func (c CandlestickFilter) GetURLKey() string {
	return fmt.Sprintf("ta_candlestick_%v", c.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (c CandlestickFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	c.Value = CandlestickType(result)
	return c
}

// BetaType represents the types of beta filters
type BetaType string

// Default filter constants
const (
	BetaUnder0           BetaType = "u0"
	BetaUnder0Point5     BetaType = "u0.5"
	BetaUnder1           BetaType = "u1"
	BetaUnder1Point5     BetaType = "u1.5"
	BetaUnder2           BetaType = "u2"
	BetaOver0            BetaType = "o0"
	BetaOver0Point5      BetaType = "o0.5"
	BetaOver1            BetaType = "o1"
	BetaOver1Point5      BetaType = "o1.5"
	BetaOver2            BetaType = "o2"
	BetaOver2Point5      BetaType = "o2.5"
	BetaOver3            BetaType = "o3"
	BetaOver4            BetaType = "o4"
	Beta0to0Point5       BetaType = "0to0.5"
	Beta0to1             BetaType = "0to1"
	Beta0Point5to1       BetaType = "0.5to1"
	Beta0Point5to1Point5 BetaType = "0.5to1.5"
	Beta1to1Point5       BetaType = "1to1.5"
	Beta1to2             BetaType = "1to2"
)

// BetaFilter is a filter type
type BetaFilter struct {
	Value BetaType
}

// GetURLKey returns the filter's url key
func (b BetaFilter) GetURLKey() string {
	return fmt.Sprintf("ta_beta_%v", b.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (b BetaFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	b.Value = BetaType(result)
	return b
}

// AverageTrueRangeType represents the types of average true range (stock volatility) filters
type AverageTrueRangeType string

// Default filter constants
const (
	ATROver0Point25  AverageTrueRangeType = "o0.25"
	ATROver0Point5   AverageTrueRangeType = "o0.5"
	ATROver0Point75  AverageTrueRangeType = "o0.75"
	ATROver1         AverageTrueRangeType = "o1"
	ATROver1Point5   AverageTrueRangeType = "o1.5"
	ATROver2         AverageTrueRangeType = "o2"
	ATROver2Point5   AverageTrueRangeType = "o2.5"
	ATROver3         AverageTrueRangeType = "o3"
	ATROver3Point5   AverageTrueRangeType = "o3.5"
	ATROver4         AverageTrueRangeType = "o4"
	ATROver4Point5   AverageTrueRangeType = "o4.5"
	ATROver5         AverageTrueRangeType = "o5"
	ATRUnder0Point25 AverageTrueRangeType = "u0.25"
	ATRUnder0Point5  AverageTrueRangeType = "u0.5"
	ATRUnder0Point75 AverageTrueRangeType = "u0.75"
	ATRUnder1        AverageTrueRangeType = "u1"
	ATRUnder1Point5  AverageTrueRangeType = "u1.5"
	ATRUnder2        AverageTrueRangeType = "u2"
	ATRUnder2Point5  AverageTrueRangeType = "u2.5"
	ATRUnder3        AverageTrueRangeType = "u3"
	ATRUnder3Point5  AverageTrueRangeType = "u3.5"
	ATRUnder4        AverageTrueRangeType = "u4"
	ATRUnder4Point5  AverageTrueRangeType = "u4.5"
	ATRUnder5        AverageTrueRangeType = "u5"
)

// AverageTrueRangeFilter is a filter type
type AverageTrueRangeFilter struct {
	Value AverageTrueRangeType
}

// GetURLKey returns the filter's url key
func (a AverageTrueRangeFilter) GetURLKey() string {
	return fmt.Sprintf("ta_averagetruerange_%v", a.Value)
}

// SetMultipleValues can be used to append several of the same filter types together using the pipe operator
func (a AverageTrueRangeFilter) SetMultipleValues(options ...interface{}) FilterInterface {
	optionCount := len(options)
	if optionCount == 0 {
		panic("No parameters given.")
	}
	result := fmt.Sprintf("%v", options[0])

	for i := 0; i < optionCount; i++ {
		result = fmt.Sprintf("%v|%v", result, options[i])
	}
	a.Value = AverageTrueRangeType(result)
	return a
}
