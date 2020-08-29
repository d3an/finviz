// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

// FilterInterface is defines methods to interact with filters
type FilterInterface interface {
	GetName() string
	GetProperties() map[string]bool
	GetProperty(string) bool
	GetURLKey() string
	GetValue() string
	SetValues(values []string) *Filter
}

// Filter represents a filter to be used for a screen
type Filter struct {
	FilterInterface

	// Add Description field from Finviz (figure out printing with md support)
	// Add field (or property) for specifying unit type
	Name       string
	Properties map[string]bool
	Value      string
	URLPrefix  string
}

/***************************************************************************
*****	DESCRIPTIVE   ******************************************************
***************************************************************************/

// ExchangeType represents the types of exchange filters
type ExchangeType = string

// Default exchange filter constants
const (
	AMEX   ExchangeType = "amex"
	NASDAQ ExchangeType = "nasd"
	NYSE   ExchangeType = "nyse"
)

// IndexType represents the types of index filters
type IndexType = string

// Default filter constants
const (
	SP500 IndexType = "sp500"
	DJIA  IndexType = "dji"
)

// SectorType represents the type of sector filters
type SectorType = string

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

// IndustryType represents the type of industry filters
type IndustryType = string

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
	CommunicationEquipment             IndustryType = "communicationequipment"
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

// CountryType represents the type of country filters
type CountryType = string

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

// MarketCapType represents the types of market cap filters
type MarketCapType = string

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

// DividendYieldType represents the type of dividend yield filters
type DividendYieldType = string

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

// ShortSellingType represents the type of short selling filters
type ShortSellingType = string

// Default filter constants
const (
	FSLow     ShortSellingType = "low"
	FSHigh    ShortSellingType = "high"
	FSUnder5  ShortSellingType = "u5"
	FSUnder10 ShortSellingType = "u10"
	FSUnder15 ShortSellingType = "u15"
	FSUnder20 ShortSellingType = "u20"
	FSUnder25 ShortSellingType = "u25"
	FSUnder30 ShortSellingType = "u30"
	FSOver5   ShortSellingType = "o5"
	FSOver10  ShortSellingType = "o10"
	FSOver15  ShortSellingType = "o15"
	FSOver20  ShortSellingType = "o20"
	FSOver25  ShortSellingType = "o25"
	FSOver30  ShortSellingType = "o30"
)

// RecommendationType represents the type of analyst recommendation filters
type RecommendationType = string

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

// OptionShortType represents the type of optionable/shortable filters
type OptionShortType = string

// Default filter constants
const (
	Option      OptionShortType = "option"
	Short       OptionShortType = "short"
	OptionShort OptionShortType = "optionshort"
)

// EarningsDateType represents the type of earnings date filters
type EarningsDateType = string

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
	EDPreviousWeek              EarningsDateType = "prevweek"
	EDThisMonth                 EarningsDateType = "thismonth"
)

// AverageVolumeType represents the type of average volume filters
type AverageVolumeType = string

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

// RelativeVolumeType represents the type of relative volume filters
type RelativeVolumeType = string

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

// CurrentVolumeType represents the type of current volume filters
type CurrentVolumeType = string

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

// PriceType represents the type of price filters
type PriceType = string

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

// TargetPriceType represents the type of target price filters
type TargetPriceType = string

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

// IPODateType represents the type of IPO Date filters
type IPODateType = string

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

// SharesOutstandingType represents the type of shares outstanding filters
type SharesOutstandingType = string

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

// FloatType represents the type of float filters
type FloatType = SharesOutstandingType

/***************************************************************************
*****	FUNDAMENTALS   *****************************************************
***************************************************************************/

// PEType represents the types of PE filters
type PEType = string

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

// ForwardPEType represents the types of ForwardPE filters
// Same categories as PEType, so those will be reused
type ForwardPEType = PEType

// PEGType represents the types of PEG filters
type PEGType = string

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

// PriceSalesType represents the types of price/sales filters
type PriceSalesType = string

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

// PriceBookType represents the types of price/book filters
type PriceBookType = PriceSalesType

// PriceCashType represents the types of price/cash filters
type PriceCashType = string

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

// PriceFCFType represents the types of price/FCF (free cash flow) filters
type PriceFCFType = string

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

// GrowthType represents the types of growth (%) filters
type GrowthType = string

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

// ReturnType represents the types of % return on assets, equity, and investment filters
type ReturnType = string

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

// AssetRatioType represents the types of asset ratio filters
type AssetRatioType = string

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

// DebtEquityType represents the types of debt/equity filters
type DebtEquityType = string

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

// GrossMarginType represents the types of gross margin filters
type GrossMarginType = string

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

// OperatingMarginType represents the types of operating margin filters
type OperatingMarginType = string

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

// NetProfitMarginType represents the types of payout ratio filters
type NetProfitMarginType = OperatingMarginType

// PayoutRatioType represents the types of payout ratio filters
type PayoutRatioType = string

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

// InsiderOwnershipType represents the types of insider ownership filters
type InsiderOwnershipType = string

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

// InsiderTransactionsType represents the types of insider transactions filters
type InsiderTransactionsType = string

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

// InstitutionalOwnershipType represents the types of institutional ownership filters
type InstitutionalOwnershipType = string

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

// InstitutionalTransactionsType represents the types of institutional transactions filters
type InstitutionalTransactionsType = string

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

/***************************************************************************
*****	TECHNICAL   ********************************************************
***************************************************************************/

// PerformanceType represents the types of performance filters
type PerformanceType = string

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

// VolatilityType represents the types of volatility filters
type VolatilityType = string

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

// RSIType represents the types of RSI 14-day filters
type RSIType = string

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

// GapType represents the types of gap filters
type GapType = string

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

// SMA20Type represents the types of 20-Day Simple Moving Average filters
type SMA20Type = string

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

// SMA50Type represents the types of 50-Day Simple Moving Average filters
type SMA50Type = string

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

// SMA200Type represents the types of 200-Day Simple Moving Average filters
type SMA200Type = string

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

// ChangeType represents the types of change (%) filters
type ChangeType = string

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

// HighLowDayType represents the types of 20-day, 50-day high/low filters
type HighLowDayType = string

// Default filter constants
const (
	DHLNewHigh                   HighLowDayType = "nh"
	DHLNewLow                    HighLowDayType = "nl"
	DHLAtLeast5PercentBelowHigh  HighLowDayType = "b5h"
	DHLAtLeast10PercentBelowHigh HighLowDayType = "b10h"
	DHLAtLeast15PercentBelowHigh HighLowDayType = "b15h"
	DHLAtLeast20PercentBelowHigh HighLowDayType = "b20h"
	DHLAtLeast30PercentBelowHigh HighLowDayType = "b30h"
	DHLAtLeast40PercentBelowHigh HighLowDayType = "b40h"
	DHLAtLeast50PercentBelowHigh HighLowDayType = "b50h"
	DHL0to3PercentBelowHigh      HighLowDayType = "b0to3h"
	DHL0to5PercentBelowHigh      HighLowDayType = "b0to5h"
	DHL0to10PercentBelowHigh     HighLowDayType = "b0to10h"
	DHLAtLeast5PercentAboveLow   HighLowDayType = "a5h"
	DHLAtLeast10PercentAboveLow  HighLowDayType = "a10h"
	DHLAtLeast15PercentAboveLow  HighLowDayType = "a15h"
	DHLAtLeast20PercentAboveLow  HighLowDayType = "a20h"
	DHLAtLeast30PercentAboveLow  HighLowDayType = "a30h"
	DHLAtLeast40PercentAboveLow  HighLowDayType = "a40h"
	DHLAtLeast50PercentAboveLow  HighLowDayType = "a50h"
	DHL0to3PercentAboveLow       HighLowDayType = "a0to3h"
	DHL0to5PercentAboveLow       HighLowDayType = "a0to5h"
	DHL0to10PercentAboveLow      HighLowDayType = "a0to10h"
)

// HighLow52WeekType represents the types of 52-week high/low filters
type HighLow52WeekType = string

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

// PatternType represents the types of technical pattern filters
type PatternType = string

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

// CandlestickType represents the types of candlestick filters
type CandlestickType = string

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

// BetaType represents the types of beta filters
type BetaType = string

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

// AverageTrueRangeType represents the types of average true range (stock volatility) filters
type AverageTrueRangeType = string

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

// FilterValueLookup is a map of default filters to common default value queries and corresponding value
var FilterValueLookup = map[string]map[string]string{
	"exchange": {
		"amex":   AMEX,
		"nasdaq": NASDAQ,
		"nasd":   NASDAQ,
		"nyse":   NYSE,
	},
	"index": {
		"s&p500": SP500,
		"sp500":  SP500,
		"djia":   DJIA,
		"dji":    DJIA,
		"dow":    DJIA,
	},
	"sector": {
		"basic materials":        BasicMaterials,
		"communication services": CommunicationServices,
		"consumer cyclical":      ConsumerCyclical,
		"consumer defensive":     ConsumerDefensive,
		"energy":                 Energy,
		"financial":              Financial,
		"healthcare":             Healthcare,
		"industrials":            Industrials,
		"real estate":            RealEstate,
		"technology":             Technology,
		"utilities":              Utilities,
	},
	"industry": {
		"stocks only":                              StocksOnly,
		"etf":                                      ETF,
		"advertising agencies":                     AdvertisingAgencies,
		"aerospace & defense":                      AerospaceAndDefense,
		"agricultural inputs":                      AgriculturalInputs,
		"airlines":                                 Airlines,
		"airports air services":                    AirportsAirServices,
		"aluminum":                                 Aluminum,
		"apparel manufacturing":                    ApparelManufacturing,
		"apparel retail":                           ApparelRetail,
		"asset management":                         AssetManagement,
		"auto manufacturers":                       AutoManufacturers,
		"auto parts":                               AutoParts,
		"auto truck dealerships":                   AutoTruckDealerships,
		"banks - diversified":                      BanksDiversified,
		"banks - regional":                         BanksRegional,
		"beverages - brewers":                      BeveragesBrewers,
		"beverages - non-alcoholic":                BeveragesNonAlcoholic,
		"beverages - wineries & distilleries":      BeveragesWineriesDistilleries,
		"biotechnology":                            Biotechnology,
		"broadcasting":                             Broadcasting,
		"building materials":                       BuildingMaterials,
		"building products & equipment":            BuildingProductsEquipment,
		"business equipment & supplies":            BusinessEquipmentSupplies,
		"capital markets":                          CapitalMarkets,
		"chemicals":                                Chemicals,
		"closed-end fund - debt":                   ClosedEndFundDebt,
		"closed-end fund - equity":                 ClosedEndFundEquity,
		"closed-end fund - foreign":                ClosedEndFundForeign,
		"coking coal":                              CokingCoal,
		"communication equipment":                  CommunicationEquipment,
		"computer hardware":                        ComputerHardware,
		"confectioners":                            Confectioners,
		"conglomerates":                            Conglomerates,
		"consulting services":                      ConsultingServices,
		"consumer electronics":                     ConsumerElectronics,
		"copper":                                   Copper,
		"credit services":                          CreditServices,
		"department stores":                        DepartmentStores,
		"diagnostics & research":                   DiagnosticsResearch,
		"discount stores":                          DiscountStores,
		"drug manufacturers - general":             DrugManufacturersGeneral,
		"drug manufacturers - specialty & generic": DrugManufacturersSpecialtyGeneric,
		"education & training services":            EducationTrainingServices,
		"electrical equipment & parts":             ElectricalEquipmentParts,
		"electronic components":                    ElectronicComponents,
		"electronic gaming multimedia":             ElectronicGamingMultimedia,
		"electronics & computer distribution":      ElectronicsComputerDistribution,
		"engineering & construction":               EngineeringConstruction,
		"entertainment":                            Entertainment,
		"farm & heavy construction machinery":      FarmHeavyConstructionMachinery,
		"farm products":                            FarmProducts,
		"financial conglomerates":                  FinancialConglomerates,
		"financial data & stock exchanges":         FinancialDataStockExchanges,
		"food distribution":                        FoodDistribution,
		"footwear & accessories":                   FootwearAccessories,
		"furnishings, fixtures & appliances":       FurnishingsFixturesAppliances,
		"gambling":                                 Gambling,
		"gold":                                     Gold,
		"grocery stores":                           GroceryStores,
		"healthcare plans":                         HealthcarePlans,
		"health information services":              HealthInformationServices,
		"home improvement retail":                  HomeImprovementRetail,
		"household & personal products":            HouseholdPersonalProducts,
		"industrial distribution":                  IndustrialDistribution,
		"information technology services":          InformationTechnologyServices,
		"infrastructure operations":                InfrastructureOperations,
		"insurance brokers":                        InsuranceBrokers,
		"insurance - diversified":                  InsuranceDiversified,
		"insurance - life":                         InsuranceLife,
		"insurance - property casualty":            InsurancePropertyCasualty,
		"insurance - reinsurance":                  InsuranceReinsurance,
		"insurance - specialty":                    InsuranceSpecialty,
		"integrated freight & logistics":           IntegratedFreightLogistics,
		"internet content & information":           InternetContentInformation,
		"internet retail":                          InternetRetail,
		"leisure":                                  Leisure,
		"lodging":                                  Lodging,
		"lumber & wood production":                 LumberWoodProduction,
		"luxury goods":                             LuxuryGoods,
		"marine shipping":                          MarineShipping,
		"medical care facilities":                  MedicalCareFacilities,
		"medical devices":                          MedicalDevices,
		"medical distribution":                     MedicalDistribution,
		"medical instruments & supplies":           MedicalInstrumentsSupplies,
		"metal fabrication":                        MetalFabrication,
		"mortgage finance":                         MortgageFinance,
		"oil & gas drilling":                       OilGasDrilling,
		"oil & gas e&p":                            OilGasEP,
		"oil & gas equipment & services":           OilGasEquipmentServices,
		"oil & gas integrated":                     OilGasIntegrated,
		"oil & gas midstream":                      OilGasMidstream,
		"oil & gas refining & marketing":           OilGasRefiningMarketing,
		"other industrial metals & mining":         OtherIndustrialMetalsMining,
		"other precious metals & mining":           OtherPreciousMetalsMining,
		"packaged foods":                           PackagedFoods,
		"packaging & containers":                   PackagingContainers,
		"paper & paper products":                   PaperPaperProducts,
		"personal services":                        PersonalServices,
		"pharmaceutical retailers":                 PharmaceuticalRetailers,
		"pollution & treatment controls":           PollutionTreatmentControls,
		"publishing":                               Publishing,
		"railroads":                                Railroads,
		"real estate - development":                RealEstateDevelopment,
		"real estate - diversified":                RealEstateDiversified,
		"real estate services":                     RealEstateServices,
		"recreational vehicles":                    RecreationalVehicles,
		"reit - diversified":                       REITDiversified,
		"reit - healthcare facilities":             REITHealthcareFacilities,
		"reit - hotel & motel":                     REITHotelMotel,
		"reit - industrial":                        REITIndustrial,
		"reit - mortgage":                          REITMortgage,
		"reit - office":                            REITOffice,
		"reit - residential":                       REITResidential,
		"reit - retail":                            REITRetail,
		"reit - specialty":                         REITSpecialty,
		"rental & leasing services":                RentalLeasingServices,
		"residential construction":                 ResidentialConstruction,
		"resorts & casinos":                        ResortsCasinos,
		"restaurants":                              Restaurants,
		"scientific & technical instruments":       ScientificTechnicalInstruments,
		"security & protection services":           SecurityProtectionServices,
		"semiconductor equipment & materials":      SemiconductorEquipmentMaterials,
		"semiconductors":                           Semiconductors,
		"shell companies":                          ShellCompanies,
		"silver":                                   Silver,
		"software - application":                   SoftwareApplication,
		"software - infrastructure":                SoftwareInfrastructure,
		"solar":                                    Solar,
		"specialty business services":              SpecialtyBusinessServices,
		"specialty chemicals":                      SpecialtyChemicals,
		"specialty industrial machinery":           SpecialtyIndustrialMachinery,
		"specialty retail":                         SpecialtyRetail,
		"staffing & employment services":           StaffingEmploymentServices,
		"steel":                                    Steel,
		"telecom services":                         TelecomServices,
		"textile manufacturing":                    TextileManufacturing,
		"thermal coal":                             ThermalCoal,
		"tobacco":                                  Tobacco,
		"tools & accessories":                      ToolsAccessories,
		"travel services":                          TravelServices,
		"trucking":                                 Trucking,
		"uranium":                                  Uranium,
		"utilities - diversified":                  UtilitiesDiversified,
		"utilities - independent power producers":  UtilitiesIndependentPowerProducers,
		"utilities - regulated electric":           UtilitiesRegulatedElectric,
		"utilities - regulated gas":                UtilitiesRegulatedGas,
		"utilities - regulated water":              UtilitiesRegulatedWater,
		"utilities - renewable":                    UtilitiesRenewable,
		"waste management":                         WasteManagement,
	},
	"country": {
		"usa":               USA,
		"foreign (ex-usa)":  NotUSA,
		"asia":              Asia,
		"europe":            Europe,
		"latin america":     LatinAmerica,
		"bric":              BRIC,
		"argentina":         Argentina,
		"australia":         Australia,
		"bahamas":           Bahamas,
		"belgium":           Belgium,
		"benelux":           BeNeLux,
		"bermuda":           Bermuda,
		"brazil":            Brazil,
		"canada":            Canada,
		"cayman islands":    CaymanIslands,
		"chile":             Chile,
		"china":             China,
		"china & hong kong": ChinaHongKong,
		"colombia":          Colombia,
		"cyprus":            Cyprus,
		"denmark":           Denmark,
		"finland":           Finland,
		"france":            France,
		"germany":           Germany,
		"greece":            Greece,
		"hong kong":         HongKong,
		"hungary":           Hungary,
		"iceland":           Iceland,
		"india":             India,
		"indonesia":         Indonesia,
		"ireland":           Ireland,
		"israel":            Israel,
		"italy":             Italy,
		"japan":             Japan,
		"kazakhstan":        Kazakhstan,
		"luxembourg":        Luxembourg,
		"malaysia":          Malaysia,
		"malta":             Malta,
		"mexico":            Mexico,
		"monaco":            Monaco,
		"netherlands":       Netherlands,
		"new zealand":       NewZealand,
		"norway":            Norway,
		"panama":            Panama,
		"peru":              Peru,
		"philippines":       Philippines,
		"portugal":          Portugal,
		"russia":            Russia,
		"singapore":         Singapore,
		"south africa":      SouthAfrica,
		"south korea":       SouthKorea,
		"spain":             Spain,
		"sweden":            Sweden,
		"switzerland":       Switzerland,
		"taiwan":            Taiwan,
		"turkey":            Turkey,
		"uae":               UAE,
		"uk":                UK,
		"uruguay":           Uruguay,
	},
	"market cap": {
		"megaover200b":   MegaOver200B,
		"large10to200b":  Large10to200B,
		"mid2to10b":      Mid2to10B,
		"small300mto2b":  Small300Mto2B,
		"micro50to300m":  Micro50to300M,
		"nanounder50m":   NanoUnder50M,
		"largeover10b":   LargeOver10B,
		"midover2b":      MidOver2B,
		"smallover300m":  SmallOver300M,
		"microover50m":   MicroOver50M,
		"largeunder200b": LargeUnder200B,
		"midunder10b":    MidUnder10B,
		"smallunder2b":   SmallUnder2B,
		"microunder300m": MicroUnder300M,
	},
	"dividend yield": {
		"none":     DYNone,
		"positive": DYPositive,
		"high":     DYHigh,
		"veryhigh": DYVeryHigh,
		"over1":    DYOver1,
		"over2":    DYOver2,
		"over3":    DYOver3,
		"over4":    DYOver4,
		"over5":    DYOver5,
		"over6":    DYOver6,
		"over7":    DYOver7,
		"over8":    DYOver8,
		"over9":    DYOver9,
		"over10":   DYOver10,
	},
	"short selling": {
		"low":     FSLow,
		"high":    FSHigh,
		"under5":  FSUnder5,
		"under10": FSUnder10,
		"under15": FSUnder15,
		"under20": FSUnder20,
		"under25": FSUnder25,
		"under30": FSUnder30,
		"over5":   FSOver5,
		"over10":  FSOver10,
		"over15":  FSOver15,
		"over20":  FSOver20,
		"over25":  FSOver25,
		"over30":  FSOver30,
	},
	"analyst recommendation": {
		"strong buy":  StrongBuy,
		"buy better":  BuyBetter,
		"buy":         Buy,
		"hold better": HoldBetter,
		"hold":        Hold,
		"hold worse":  HoldWorse,
		"sell":        Sell,
		"sell worse":  SellWorse,
		"strong sell": StrongSell,
	},
	"option short": {
		"optionable":               Option,
		"shortable":                Short,
		"optionable and shortable": OptionShort,
	},
	"earnings date": {
		"today":                        EDToday,
		"today before market open":     EDTodayBeforeMarketOpen,
		"today after market close":     EDTodayAfterMarketClose,
		"tomorrow":                     EDTomorrow,
		"tomorrow before market open":  EDTomorrowBeforeMarketOpen,
		"tomorrow after market close":  EDTomorrowAfterMarketClose,
		"yesterday":                    EDYesterday,
		"yesterday before market open": EDYesterdayBeforeMarketOpen,
		"yesterday after market close": EDYesterdayAfterMarketClose,
		"next 5 days":                  EDNext5Days,
		"previous 5 days":              EDPrevious5Days,
		"this week":                    EDThisWeek,
		"next week":                    EDNextWeek,
		"previous week":                EDPreviousWeek,
		"this month":                   EDThisMonth,
	},
	"average volume": {
		"under 50K":           AvgVolUnder50K,
		"under 100K":          AvgVolUnder100K,
		"under 500K":          AvgVolUnder500K,
		"under 750K":          AvgVolUnder750K,
		"under 1M":            AvgVolUnder1M,
		"over 50K":            AvgVolOver50K,
		"over 100K":           AvgVolOver100K,
		"over 200K":           AvgVolOver200K,
		"over 300K":           AvgVolOver300K,
		"over 400K":           AvgVolOver400K,
		"over 500K":           AvgVolOver500K,
		"over 750K":           AvgVolOver750K,
		"over 1M":             AvgVolOver1M,
		"over 2M":             AvgVolOver2M,
		"between 100 to 500K": AvgVolBetween100to500K,
		"between 100K to 1M":  AvgVolBetween100Kto1M,
		"between 500K to 1M":  AvgVolBetween500Kto1M,
		"between 500K to 10M": AvgVolBetween500Kto10M,
	},
	"relative volume": {
		"over 10":    RVOver10,
		"over 5":     RVOver5,
		"over 3":     RVOver3,
		"over 2":     RVOver2,
		"over 1.5":   RVOver1point5,
		"over 1":     RVOver1,
		"over 0.75":  RVOver0point75,
		"over 0.5":   RVOver0point5,
		"over 0.25":  RVOver0point25,
		"under 2":    RVUnder2,
		"under 1.5":  RVUnder1point5,
		"under 1":    RVUnder1,
		"under 0.75": RVUnder0point75,
		"under 0.5":  RVUnder0point5,
		"under 0.25": RVUnder0point25,
		"under 0.1":  RVUnder0point1,
	},
	"current volume": {
		"under 50K":  CurVolUnder50K,
		"under 100K": CurVolUnder100K,
		"under 500K": CurVolUnder500K,
		"under 750K": CurVolUnder750K,
		"under 1M":   CurVolUnder1M,
		"over 0":     CurVolOver0,
		"over 50K":   CurVolOver50K,
		"over 100K":  CurVolOver100K,
		"over 200K":  CurVolOver200K,
		"over 300K":  CurVolOver300K,
		"over 400K":  CurVolOver400K,
		"over 500K":  CurVolOver500K,
		"over 750K":  CurVolOver750K,
		"over 1M":    CurVolOver1M,
		"over 2M":    CurVolOver2M,
		"over 5M":    CurVolOver5M,
		"over 10M":   CurVolOver10M,
		"over 20M":   CurVolOver20M,
	},
	"price": {
		"under 1":   PriceUnder1,
		"under 2":   PriceUnder2,
		"under 3":   PriceUnder3,
		"under 4":   PriceUnder4,
		"under 5":   PriceUnder5,
		"under 7":   PriceUnder7,
		"under 10":  PriceUnder10,
		"under 15":  PriceUnder15,
		"under 20":  PriceUnder20,
		"under 30":  PriceUnder30,
		"under 40":  PriceUnder40,
		"under 50":  PriceUnder50,
		"over 1":    PriceOver1,
		"over 2":    PriceOver2,
		"over 3":    PriceOver3,
		"over 4":    PriceOver4,
		"over 5":    PriceOver5,
		"over 7":    PriceOver7,
		"over 10":   PriceOver10,
		"over 15":   PriceOver15,
		"over 20":   PriceOver20,
		"over 30":   PriceOver30,
		"over 40":   PriceOver40,
		"over 50":   PriceOver50,
		"over 60":   PriceOver60,
		"over 70":   PriceOver70,
		"over 80":   PriceOver80,
		"over 90":   PriceOver90,
		"over 100":  PriceOver100,
		"1 to 5":    Price1to5,
		"1 to 10":   Price1to10,
		"1 to 20":   Price1to20,
		"5 to 10":   Price5to10,
		"5 to 20":   Price5to20,
		"5 to 50":   Price5to50,
		"10 to 20":  Price10to20,
		"10 to 50":  Price10to50,
		"20 to 50":  Price20to50,
		"50 to 100": Price50to100,
	},
	"target price": {
		"50% above price": TargetAbovePriceBy50Percent,
		"40% above price": TargetAbovePriceBy40Percent,
		"30% above price": TargetAbovePriceBy30Percent,
		"20% above price": TargetAbovePriceBy20Percent,
		"10% above price": TargetAbovePriceBy10Percent,
		"5% above price":  TargetAbovePriceBy5Percent,
		"above price":     TargetAbovePrice,
		"below price":     TargetBelowPrice,
		"50% below price": TargetBelowPriceBy50Percent,
		"40% below price": TargetBelowPriceBy40Percent,
		"30% below price": TargetBelowPriceBy30Percent,
		"20% below price": TargetBelowPriceBy20Percent,
		"10% below price": TargetBelowPriceBy10Percent,
		"5% below price":  TargetBelowPriceBy5Percent,
	},
	"ipo date": {
		"Today":                  IDToday,
		"Yesterday":              IDYesterday,
		"Previous Week":          IDPreviousWeek,
		"Previous Month":         IDPreviousMonth,
		"Previous Quarter":       IDPreviousQuarter,
		"Previous Year":          IDPreviousYear,
		"Previous 2 Years":       IDPrevious2Years,
		"Previous 3 Years":       IDPrevious3Years,
		"Previous 5 Years":       IDPrevious5Years,
		"More Than 1 Year Ago":   IDMoreThan1YearAgo,
		"More Than 5 Years Ago":  IDMoreThan5YearsAgo,
		"More Than 10 Years Ago": IDMoreThan10YearsAgo,
		"More Than 15 Years Ago": IDMoreThan15YearsAgo,
		"More Than 20 Years Ago": IDMoreThan20YearsAgo,
		"More Than 25 Years Ago": IDMoreThan25YearsAgo,
	},
	"shares outstanding": {
		"under 1M":   SOUnder1M,
		"under 5M":   SOUnder5M,
		"under 10M":  SOUnder10M,
		"under 20M":  SOUnder20M,
		"under 50M":  SOUnder50M,
		"under 100M": SOUnder100M,
		"over 1M":    SOOver1M,
		"over 2M":    SOOver2M,
		"over 5M":    SOOver5M,
		"over 10M":   SOOver10M,
		"over 20M":   SOOver20M,
		"over 50M":   SOOver50M,
		"over 100M":  SOOver100M,
		"over 200M":  SOOver200M,
		"over 500M":  SOOver500M,
		"over 1000M": SOOver1000M,
	},
	"float": {
		"under 1M":   SOUnder1M,
		"under 5M":   SOUnder5M,
		"under 10M":  SOUnder10M,
		"under 20M":  SOUnder20M,
		"under 50M":  SOUnder50M,
		"under 100M": SOUnder100M,
		"over 1M":    SOOver1M,
		"over 2M":    SOOver2M,
		"over 5M":    SOOver5M,
		"over 10M":   SOOver10M,
		"over 20M":   SOOver20M,
		"over 50M":   SOOver50M,
		"over 100M":  SOOver100M,
		"over 200M":  SOOver200M,
		"over 500M":  SOOver500M,
		"over 1000M": SOOver1000M,
	},
	"p/e": {
		"low":        PELow,
		"profitable": PEProfitable,
		"high":       PEHigh,
		"under 5":    PEUnder5,
		"under 10":   PEUnder10,
		"under 15":   PEUnder15,
		"under 20":   PEUnder20,
		"under 25":   PEUnder25,
		"under 30":   PEUnder30,
		"under 35":   PEUnder35,
		"under 40":   PEUnder40,
		"under 45":   PEUnder45,
		"under 50":   PEUnder50,
		"over 5":     PEOver5,
		"over 10":    PEOver10,
		"over 15":    PEOver15,
		"over 20":    PEOver20,
		"over 25":    PEOver25,
		"over 30":    PEOver30,
		"over 35":    PEOver35,
		"over 40":    PEOver40,
		"over 45":    PEOver45,
		"over 50":    PEOver50,
	},
	"forward p/e": {
		"low":        PELow,
		"profitable": PEProfitable,
		"high":       PEHigh,
		"under 5":    PEUnder5,
		"under 10":   PEUnder10,
		"under 15":   PEUnder15,
		"under 20":   PEUnder20,
		"under 25":   PEUnder25,
		"under 30":   PEUnder30,
		"under 35":   PEUnder35,
		"under 40":   PEUnder40,
		"under 45":   PEUnder45,
		"under 50":   PEUnder50,
		"over 5":     PEOver5,
		"over 10":    PEOver10,
		"over 15":    PEOver15,
		"over 20":    PEOver20,
		"over 25":    PEOver25,
		"over 30":    PEOver30,
		"over 35":    PEOver35,
		"over 40":    PEOver40,
		"over 45":    PEOver45,
		"over 50":    PEOver50,
	},
	"peg": {
		"low":     PEGLow,
		"high":    PEGHigh,
		"under 1": PEGUnder1,
		"under 2": PEGUnder2,
		"under 3": PEGUnder3,
		"over 1":  PEGOver1,
		"over 2":  PEGOver2,
		"over 3":  PEGOver3,
	},
	"p/s": {
		"low":      PSLow,
		"high":     PSHigh,
		"under 1":  PSUnder1,
		"under 2":  PSUnder2,
		"under 3":  PSUnder3,
		"under 4":  PSUnder4,
		"under 5":  PSUnder5,
		"under 6":  PSUnder6,
		"under 7":  PSUnder7,
		"under 8":  PSUnder8,
		"under 9":  PSUnder9,
		"under 10": PSUnder10,
		"over 1":   PSOver1,
		"over 2":   PSOver2,
		"over 3":   PSOver3,
		"over 4":   PSOver4,
		"over 5":   PSOver5,
		"over 6":   PSOver6,
		"over 7":   PSOver7,
		"over 8":   PSOver8,
		"over 9":   PSOver9,
		"over 10":  PSOver10,
	},
	"p/b": {
		"low":      PSLow,
		"high":     PSHigh,
		"under 1":  PSUnder1,
		"under 2":  PSUnder2,
		"under 3":  PSUnder3,
		"under 4":  PSUnder4,
		"under 5":  PSUnder5,
		"under 6":  PSUnder6,
		"under 7":  PSUnder7,
		"under 8":  PSUnder8,
		"under 9":  PSUnder9,
		"under 10": PSUnder10,
		"over 1":   PSOver1,
		"over 2":   PSOver2,
		"over 3":   PSOver3,
		"over 4":   PSOver4,
		"over 5":   PSOver5,
		"over 6":   PSOver6,
		"over 7":   PSOver7,
		"over 8":   PSOver8,
		"over 9":   PSOver9,
		"over 10":  PSOver10,
	},
	"price/cash": {
		"low":      PCLow,
		"high":     PCHigh,
		"under 1":  PCUnder1,
		"under 2":  PCUnder2,
		"under 3":  PCUnder3,
		"under 4":  PCUnder4,
		"under 5":  PCUnder5,
		"under 6":  PCUnder6,
		"under 7":  PCUnder7,
		"under 8":  PCUnder8,
		"under 9":  PCUnder9,
		"under 10": PCUnder10,
		"over 1":   PCOver1,
		"over 2":   PCOver2,
		"over 3":   PCOver3,
		"over 4":   PCOver4,
		"over 5":   PCOver5,
		"over 6":   PCOver6,
		"over 7":   PCOver7,
		"over 8":   PCOver8,
		"over 9":   PCOver9,
		"over 10":  PCOver10,
		"over 20":  PCOver20,
		"over 30":  PCOver30,
		"over 40":  PCOver40,
		"over 50":  PCOver50,
	},
	"price/free cash flow": {
		"low":       PFCFLow,
		"high":      PFCFHigh,
		"under 5":   PFCFUnder5,
		"under 10":  PFCFUnder10,
		"under 15":  PFCFUnder15,
		"under 20":  PFCFUnder20,
		"under 25":  PFCFUnder25,
		"under 30":  PFCFUnder30,
		"under 35":  PFCFUnder35,
		"under 40":  PFCFUnder40,
		"under 45":  PFCFUnder45,
		"under 50":  PFCFUnder50,
		"under 60":  PFCFUnder60,
		"under 70":  PFCFUnder70,
		"under 80":  PFCFUnder80,
		"under 90":  PFCFUnder90,
		"under 100": PFCFUnder100,
		"over 5":    PFCFOver5,
		"over 10":   PFCFOver10,
		"over 15":   PFCFOver15,
		"over 20":   PFCFOver20,
		"over 25":   PFCFOver25,
		"over 30":   PFCFOver30,
		"over 35":   PFCFOver35,
		"over 40":   PFCFOver40,
		"over 45":   PFCFOver45,
		"over 50":   PFCFOver50,
		"over 60":   PFCFOver60,
		"over 70":   PFCFOver70,
		"over 80":   PFCFOver80,
		"over 90":   PFCFOver90,
		"over 100":  PFCFOver100,
	},
	"eps growth this year": {
		"negative":     GrowthNegative,
		"positive":     GrowthPositive,
		"positive low": GrowthPositiveLow,
		"high":         GrowthHigh,
		"under 5":      GrowthUnder5,
		"under 10":     GrowthUnder10,
		"under 15":     GrowthUnder15,
		"under 20":     GrowthUnder20,
		"under 25":     GrowthUnder25,
		"under 30":     GrowthUnder30,
		"over 5":       GrowthOver5,
		"over 10":      GrowthOver10,
		"over 15":      GrowthOver15,
		"over 20":      GrowthOver20,
		"over 25":      GrowthOver25,
		"over 30":      GrowthOver30,
	},
	"eps growth next year": {
		"negative":     GrowthNegative,
		"positive":     GrowthPositive,
		"positive low": GrowthPositiveLow,
		"high":         GrowthHigh,
		"under 5":      GrowthUnder5,
		"under 10":     GrowthUnder10,
		"under 15":     GrowthUnder15,
		"under 20":     GrowthUnder20,
		"under 25":     GrowthUnder25,
		"under 30":     GrowthUnder30,
		"over 5":       GrowthOver5,
		"over 10":      GrowthOver10,
		"over 15":      GrowthOver15,
		"over 20":      GrowthOver20,
		"over 25":      GrowthOver25,
		"over 30":      GrowthOver30,
	},
	"eps growth past 5 years": {
		"negative":     GrowthNegative,
		"positive":     GrowthPositive,
		"positive low": GrowthPositiveLow,
		"high":         GrowthHigh,
		"under 5":      GrowthUnder5,
		"under 10":     GrowthUnder10,
		"under 15":     GrowthUnder15,
		"under 20":     GrowthUnder20,
		"under 25":     GrowthUnder25,
		"under 30":     GrowthUnder30,
		"over 5":       GrowthOver5,
		"over 10":      GrowthOver10,
		"over 15":      GrowthOver15,
		"over 20":      GrowthOver20,
		"over 25":      GrowthOver25,
		"over 30":      GrowthOver30,
	},
	"eps growth next 5 years": {
		"negative":     GrowthNegative,
		"positive":     GrowthPositive,
		"positive low": GrowthPositiveLow,
		"high":         GrowthHigh,
		"under 5":      GrowthUnder5,
		"under 10":     GrowthUnder10,
		"under 15":     GrowthUnder15,
		"under 20":     GrowthUnder20,
		"under 25":     GrowthUnder25,
		"under 30":     GrowthUnder30,
		"over 5":       GrowthOver5,
		"over 10":      GrowthOver10,
		"over 15":      GrowthOver15,
		"over 20":      GrowthOver20,
		"over 25":      GrowthOver25,
		"over 30":      GrowthOver30,
	},
	"sales growth past 5 years": {
		"negative":     GrowthNegative,
		"positive":     GrowthPositive,
		"positive low": GrowthPositiveLow,
		"high":         GrowthHigh,
		"under 5":      GrowthUnder5,
		"under 10":     GrowthUnder10,
		"under 15":     GrowthUnder15,
		"under 20":     GrowthUnder20,
		"under 25":     GrowthUnder25,
		"under 30":     GrowthUnder30,
		"over 5":       GrowthOver5,
		"over 10":      GrowthOver10,
		"over 15":      GrowthOver15,
		"over 20":      GrowthOver20,
		"over 25":      GrowthOver25,
		"over 30":      GrowthOver30,
	},
	"eps growth quarter over quarter": {
		"negative":     GrowthNegative,
		"positive":     GrowthPositive,
		"positive low": GrowthPositiveLow,
		"high":         GrowthHigh,
		"under 5":      GrowthUnder5,
		"under 10":     GrowthUnder10,
		"under 15":     GrowthUnder15,
		"under 20":     GrowthUnder20,
		"under 25":     GrowthUnder25,
		"under 30":     GrowthUnder30,
		"over 5":       GrowthOver5,
		"over 10":      GrowthOver10,
		"over 15":      GrowthOver15,
		"over 20":      GrowthOver20,
		"over 25":      GrowthOver25,
		"over 30":      GrowthOver30,
	},
	"sales growth quarter over quarter": {
		"negative":     GrowthNegative,
		"positive":     GrowthPositive,
		"positive low": GrowthPositiveLow,
		"high":         GrowthHigh,
		"under 5":      GrowthUnder5,
		"under 10":     GrowthUnder10,
		"under 15":     GrowthUnder15,
		"under 20":     GrowthUnder20,
		"under 25":     GrowthUnder25,
		"under 30":     GrowthUnder30,
		"over 5":       GrowthOver5,
		"over 10":      GrowthOver10,
		"over 15":      GrowthOver15,
		"over 20":      GrowthOver20,
		"over 25":      GrowthOver25,
		"over 30":      GrowthOver30,
	},
	"roa": {
		"positive":      ReturnPositive,
		"negative":      ReturnNegative,
		"very positive": ReturnVeryPositive,
		"very negative": ReturnVeryNegative,
		"under -50":     ReturnUnderNeg50,
		"under -45":     ReturnUnderNeg45,
		"under -40":     ReturnUnderNeg40,
		"under -35":     ReturnUnderNeg35,
		"under -30":     ReturnUnderNeg30,
		"under -25":     ReturnUnderNeg25,
		"under -20":     ReturnUnderNeg20,
		"under -15":     ReturnUnderNeg15,
		"under -10":     ReturnUnderNeg10,
		"under -5":      ReturnUnderNeg5,
		"Over 50":       ReturnOver50,
		"Over 45":       ReturnOver45,
		"Over 40":       ReturnOver40,
		"Over 35":       ReturnOver35,
		"Over 30":       ReturnOver30,
		"Over 25":       ReturnOver25,
		"Over 20":       ReturnOver20,
		"Over 15":       ReturnOver15,
		"Over 10":       ReturnOver10,
		"Over 5":        ReturnOver5,
	},
	"roe": {
		"positive":      ReturnPositive,
		"negative":      ReturnNegative,
		"very positive": ReturnVeryPositive,
		"very negative": ReturnVeryNegative,
		"under -50":     ReturnUnderNeg50,
		"under -45":     ReturnUnderNeg45,
		"under -40":     ReturnUnderNeg40,
		"under -35":     ReturnUnderNeg35,
		"under -30":     ReturnUnderNeg30,
		"under -25":     ReturnUnderNeg25,
		"under -20":     ReturnUnderNeg20,
		"under -15":     ReturnUnderNeg15,
		"under -10":     ReturnUnderNeg10,
		"under -5":      ReturnUnderNeg5,
		"Over 50":       ReturnOver50,
		"Over 45":       ReturnOver45,
		"Over 40":       ReturnOver40,
		"Over 35":       ReturnOver35,
		"Over 30":       ReturnOver30,
		"Over 25":       ReturnOver25,
		"Over 20":       ReturnOver20,
		"Over 15":       ReturnOver15,
		"Over 10":       ReturnOver10,
		"Over 5":        ReturnOver5,
	},
	"roi": {
		"positive":      ReturnPositive,
		"negative":      ReturnNegative,
		"very positive": ReturnVeryPositive,
		"very negative": ReturnVeryNegative,
		"under -50":     ReturnUnderNeg50,
		"under -45":     ReturnUnderNeg45,
		"under -40":     ReturnUnderNeg40,
		"under -35":     ReturnUnderNeg35,
		"under -30":     ReturnUnderNeg30,
		"under -25":     ReturnUnderNeg25,
		"under -20":     ReturnUnderNeg20,
		"under -15":     ReturnUnderNeg15,
		"under -10":     ReturnUnderNeg10,
		"under -5":      ReturnUnderNeg5,
		"Over 50":       ReturnOver50,
		"Over 45":       ReturnOver45,
		"Over 40":       ReturnOver40,
		"Over 35":       ReturnOver35,
		"Over 30":       ReturnOver30,
		"Over 25":       ReturnOver25,
		"Over 20":       ReturnOver20,
		"Over 15":       ReturnOver15,
		"Over 10":       ReturnOver10,
		"Over 5":        ReturnOver5,
	},
	"current ratio": {
		"high":      RatioHigh,
		"low":       RatioLow,
		"under 1":   RatioUnder1,
		"under 0.5": RatioUnder0point5,
		"over 0.5":  RatioOver0point5,
		"over 1":    RatioOver1,
		"over 1.5":  RatioOver1point5,
		"over 2":    RatioOver2,
		"over 3":    RatioOver3,
		"over 4":    RatioOver4,
		"over 5":    RatioOver5,
		"over 10":   RatioOver10,
	},
	"quick ratio": {
		"high":      RatioHigh,
		"low":       RatioLow,
		"under 1":   RatioUnder1,
		"under 0.5": RatioUnder0point5,
		"over 0.5":  RatioOver0point5,
		"over 1":    RatioOver1,
		"over 1.5":  RatioOver1point5,
		"over 2":    RatioOver2,
		"over 3":    RatioOver3,
		"over 4":    RatioOver4,
		"over 5":    RatioOver5,
		"over 10":   RatioOver10,
	},
	"long-term debt/equity": {
		"high":      DEHigh,
		"low":       DELow,
		"under 1":   DEUnder1,
		"under 0.9": DEUnder0point9,
		"under 0.8": DEUnder0point8,
		"under 0.7": DEUnder0point7,
		"under 0.6": DEUnder0point6,
		"under 0.5": DEUnder0point5,
		"under 0.4": DEUnder0point4,
		"under 0.3": DEUnder0point3,
		"under 0.2": DEUnder0point2,
		"under 0.1": DEUnder0point1,
		"over 1":    DEOver1,
		"over 0.9":  DEOver0point9,
		"over 0.8":  DEOver0point8,
		"over 0.7":  DEOver0point7,
		"over 0.6":  DEOver0point6,
		"over 0.5":  DEOver0point5,
		"over 0.4":  DEOver0point4,
		"over 0.3":  DEOver0point3,
		"over 0.2":  DEOver0point2,
		"over 0.1":  DEOver0point1,
	},
	"debt/equity": {
		"high":      DEHigh,
		"low":       DELow,
		"under 1":   DEUnder1,
		"under 0.9": DEUnder0point9,
		"under 0.8": DEUnder0point8,
		"under 0.7": DEUnder0point7,
		"under 0.6": DEUnder0point6,
		"under 0.5": DEUnder0point5,
		"under 0.4": DEUnder0point4,
		"under 0.3": DEUnder0point3,
		"under 0.2": DEUnder0point2,
		"under 0.1": DEUnder0point1,
		"over 1":    DEOver1,
		"over 0.9":  DEOver0point9,
		"over 0.8":  DEOver0point8,
		"over 0.7":  DEOver0point7,
		"over 0.6":  DEOver0point6,
		"over 0.5":  DEOver0point5,
		"over 0.4":  DEOver0point4,
		"over 0.3":  DEOver0point3,
		"over 0.2":  DEOver0point2,
		"over 0.1":  DEOver0point1,
	},
	"gross margin": {
		"positive":   GMPositive,
		"negative":   GMNegative,
		"high":       GMHigh,
		"under 90":   GMUnder90,
		"under 80":   GMUnder80,
		"under 70":   GMUnder70,
		"under 60":   GMUnder60,
		"under 50":   GMUnder50,
		"under 45":   GMUnder45,
		"under 40":   GMUnder40,
		"under 35":   GMUnder35,
		"under 30":   GMUnder30,
		"under 25":   GMUnder25,
		"under 20":   GMUnder20,
		"under 15":   GMUnder15,
		"under 10":   GMUnder10,
		"under 5":    GMUnder5,
		"under 0":    GMUnder0,
		"under -10":  GMUnderNeg10,
		"under -20":  GMUnderNeg20,
		"under -30":  GMUnderNeg30,
		"under -50":  GMUnderNeg50,
		"under -70":  GMUnderNeg70,
		"under -100": GMUnderNeg100,
		"over 0":     GMOver0,
		"over 5":     GMOver5,
		"over 10":    GMOver10,
		"over 15":    GMOver15,
		"over 20":    GMOver20,
		"over 25":    GMOver25,
		"over 30":    GMOver30,
		"over 35":    GMOver35,
		"over 40":    GMOver40,
		"over 45":    GMOver45,
		"over 50":    GMOver50,
		"over 60":    GMOver60,
		"over 70":    GMOver70,
		"over 80":    GMOver80,
		"over 90":    GMOver90,
	},
	"operating margin": {
		"positive":      OMPositive,
		"negative":      OMNegative,
		"very negative": OMVeryNegative,
		"high":          OMHigh,
		"under 90":      OMUnder90,
		"under 80":      OMUnder80,
		"under 70":      OMUnder70,
		"under 60":      OMUnder60,
		"under 50":      OMUnder50,
		"under 45":      OMUnder45,
		"under 40":      OMUnder40,
		"under 35":      OMUnder35,
		"under 30":      OMUnder30,
		"under 25":      OMUnder25,
		"under 20":      OMUnder20,
		"under 15":      OMUnder15,
		"under 10":      OMUnder10,
		"under 5":       OMUnder5,
		"under 0":       OMUnder0,
		"under -10":     OMUnderNeg10,
		"under -20":     OMUnderNeg20,
		"under -30":     OMUnderNeg30,
		"under -50":     OMUnderNeg50,
		"under -70":     OMUnderNeg70,
		"under -100":    OMUnderNeg100,
		"over 0":        OMOver0,
		"over 5":        OMOver5,
		"over 10":       OMOver10,
		"over 15":       OMOver15,
		"over 20":       OMOver20,
		"over 25":       OMOver25,
		"over 30":       OMOver30,
		"over 35":       OMOver35,
		"over 40":       OMOver40,
		"over 45":       OMOver45,
		"over 50":       OMOver50,
		"over 60":       OMOver60,
		"over 70":       OMOver70,
		"over 80":       OMOver80,
		"over 90":       OMOver90,
	},
	"net profit margin": {
		"positive":      OMPositive,
		"negative":      OMNegative,
		"very negative": OMVeryNegative,
		"high":          OMHigh,
		"under 90":      OMUnder90,
		"under 80":      OMUnder80,
		"under 70":      OMUnder70,
		"under 60":      OMUnder60,
		"under 50":      OMUnder50,
		"under 45":      OMUnder45,
		"under 40":      OMUnder40,
		"under 35":      OMUnder35,
		"under 30":      OMUnder30,
		"under 25":      OMUnder25,
		"under 20":      OMUnder20,
		"under 15":      OMUnder15,
		"under 10":      OMUnder10,
		"under 5":       OMUnder5,
		"under 0":       OMUnder0,
		"under -10":     OMUnderNeg10,
		"under -20":     OMUnderNeg20,
		"under -30":     OMUnderNeg30,
		"under -50":     OMUnderNeg50,
		"under -70":     OMUnderNeg70,
		"under -100":    OMUnderNeg100,
		"over 0":        OMOver0,
		"over 5":        OMOver5,
		"over 10":       OMOver10,
		"over 15":       OMOver15,
		"over 20":       OMOver20,
		"over 25":       OMOver25,
		"over 30":       OMOver30,
		"over 35":       OMOver35,
		"over 40":       OMOver40,
		"over 45":       OMOver45,
		"over 50":       OMOver50,
		"over 60":       OMOver60,
		"over 70":       OMOver70,
		"over 80":       OMOver80,
		"over 90":       OMOver90,
	},
	"payout ratio": {
		"none":      PRNone,
		"positive":  PRPositive,
		"low":       PRLow,
		"high":      PRHigh,
		"over 1%":   PROver1,
		"over 2%":   PROver2,
		"over 3%":   PROver3,
		"over 4%":   PROver4,
		"over 5%":   PROver5,
		"over 6%":   PROver6,
		"over 7%":   PROver7,
		"over 8%":   PROver8,
		"over 9%":   PROver9,
		"over 10%":  PROver10,
		"under 1%":  PRUnder1,
		"under 2%":  PRUnder2,
		"under 3%":  PRUnder3,
		"under 4%":  PRUnder4,
		"under 5%":  PRUnder5,
		"under 6%":  PRUnder6,
		"under 7%":  PRUnder7,
		"under 8%":  PRUnder8,
		"under 9%":  PRUnder9,
		"under 10%": PRUnder10,
	},
	"insider ownership": {
		"low":       InsdrOwnLow,
		"high":      InsdrOwnHigh,
		"very high": InsdrOwnVeryHigh,
		"over 10%":  InsdrOwnOver10,
		"over 20%":  InsdrOwnOver20,
		"over 30%":  InsdrOwnOver30,
		"over 40%":  InsdrOwnOver40,
		"over 50%":  InsdrOwnOver50,
		"over 60%":  InsdrOwnOver60,
		"over 70%":  InsdrOwnOver70,
		"over 80%":  InsdrOwnOver80,
		"over 90%":  InsdrOwnOver90,
	},
	"insider transactions": {
		"very negative": InsdrTransVeryNeg,
		"negative":      InsdrTransNeg,
		"positive":      InsdrTransPos,
		"very positive": InsdrTransVeryPos,
		"over +5%":      InsdrTransOver5,
		"over +10%":     InsdrTransOver10,
		"over +15%":     InsdrTransOver15,
		"over +20%":     InsdrTransOver20,
		"over +25%":     InsdrTransOver25,
		"over +30%":     InsdrTransOver30,
		"over +35%":     InsdrTransOver35,
		"over +40%":     InsdrTransOver40,
		"over +45%":     InsdrTransOver45,
		"over +50%":     InsdrTransOver50,
		"over +60%":     InsdrTransOver60,
		"over +70%":     InsdrTransOver70,
		"over +80%":     InsdrTransOver80,
		"over +90%":     InsdrTransOver90,
		"under -5%":     InsdrTransUnderNeg5,
		"under -10%":    InsdrTransUnderNeg10,
		"under -15%":    InsdrTransUnderNeg15,
		"under -20%":    InsdrTransUnderNeg20,
		"under -25%":    InsdrTransUnderNeg25,
		"under -30%":    InsdrTransUnderNeg30,
		"under -35%":    InsdrTransUnderNeg35,
		"under -40%":    InsdrTransUnderNeg40,
		"under -45%":    InsdrTransUnderNeg45,
		"under -50%":    InsdrTransUnderNeg50,
		"under -60%":    InsdrTransUnderNeg60,
		"under -70%":    InsdrTransUnderNeg70,
		"under -80%":    InsdrTransUnderNeg80,
		"under -90%":    InsdrTransUnderNeg90,
	},
	"institutional ownership": {
		"low":       InstOwnLow,
		"high":      InstOwnHigh,
		"over 10%":  InstOwnOver10,
		"over 20%":  InstOwnOver20,
		"over 30%":  InstOwnOver30,
		"over 40%":  InstOwnOver40,
		"over 50%":  InstOwnOver50,
		"over 60%":  InstOwnOver60,
		"over 70%":  InstOwnOver70,
		"over 80%":  InstOwnOver80,
		"over 90%":  InstOwnOver90,
		"under 10%": InstOwnUnder10,
		"under 20%": InstOwnUnder20,
		"under 30%": InstOwnUnder30,
		"under 40%": InstOwnUnder40,
		"under 50%": InstOwnUnder50,
		"under 60%": InstOwnUnder60,
		"under 70%": InstOwnUnder70,
		"under 80%": InstOwnUnder80,
		"under 90%": InstOwnUnder90,
	},
	"institutional transactions": {
		"very negative": InstTransVeryNeg,
		"negative":      InstTransNeg,
		"positive":      InstTransPos,
		"very positive": InstTransVeryPos,
		"over +5%":      InstTransOver5,
		"over +10%":     InstTransOver10,
		"over +15%":     InstTransOver15,
		"over +20%":     InstTransOver20,
		"over +25%":     InstTransOver25,
		"over +30%":     InstTransOver30,
		"over +35%":     InstTransOver35,
		"over +40%":     InstTransOver40,
		"over +45%":     InstTransOver45,
		"over +50%":     InstTransOver50,
		"under -5%":     InstTransUnderNeg5,
		"under -10%":    InstTransUnderNeg10,
		"under -15%":    InstTransUnderNeg15,
		"under -20%":    InstTransUnderNeg20,
		"under -25%":    InstTransUnderNeg25,
		"under -30%":    InstTransUnderNeg30,
		"under -35%":    InstTransUnderNeg35,
		"under -40%":    InstTransUnderNeg40,
		"under -45%":    InstTransUnderNeg45,
		"under -50%":    InstTransUnderNeg50,
	},

	"performance": {
		"today up":     PerfTodayUp,
		"today down":   PerfTodayDown,
		"today -15%":   PerfTodayDown15Percent,
		"today -10%":   PerfTodayDown10Percent,
		"today -5%":    PerfTodayDown5Percent,
		"today +5%":    PerfTodayUp5Percent,
		"today +10%":   PerfTodayUp10Percent,
		"today +15%":   PerfTodayUp15Percent,
		"week -30%":    PerfWeekDown30Percent,
		"week -20%":    PerfWeekDown20Percent,
		"week -10%":    PerfWeekDown10Percent,
		"week down":    PerfWeekDown,
		"week up":      PerfWeekUp,
		"week +10%":    PerfWeekUp10Percent,
		"week +20%":    PerfWeekUp20Percent,
		"week +30%":    PerfWeekUp30Percent,
		"month -50%":   PerfMonthDown50Percent,
		"month -30%":   PerfMonthDown30Percent,
		"month -20%":   PerfMonthDown20Percent,
		"month -10%":   PerfMonthDown10Percent,
		"month down":   PerfMonthDown,
		"month up":     PerfMonthUp,
		"month +10%":   PerfMonthUp10Percent,
		"month +20%":   PerfMonthUp20Percent,
		"month +30%":   PerfMonthUp30Percent,
		"month +50%":   PerfMonthUp50Percent,
		"quarter -50%": PerfQuarterDown50Percent,
		"quarter -30%": PerfQuarterDown30Percent,
		"quarter -20%": PerfQuarterDown20Percent,
		"quarter -10%": PerfQuarterDown10Percent,
		"quarter Down": PerfQuarterDown,
		"quarter Up":   PerfQuarterUp,
		"quarter +10%": PerfQuarterUp10Percent,
		"quarter +20%": PerfQuarterUp20Percent,
		"quarter +30%": PerfQuarterUp30Percent,
		"quarter +50%": PerfQuarterUp50Percent,
		"half -75%":    PerfSixMonthsDown75Percent,
		"half -50%":    PerfSixMonthsDown50Percent,
		"half -30%":    PerfSixMonthsDown30Percent,
		"half -20%":    PerfSixMonthsDown20Percent,
		"half -10%":    PerfSixMonthsDown10Percent,
		"half down":    PerfSixMonthsDown,
		"half up":      PerfSixMonthsUp,
		"half +10%":    PerfSixMonthsUp10Percent,
		"half +20%":    PerfSixMonthsUp20Percent,
		"half +30%":    PerfSixMonthsUp30Percent,
		"half +50%":    PerfSixMonthsUp50Percent,
		"half +100%":   PerfSixMonthsUp100Percent,
		"year -75%":    PerfYearDown75Percent,
		"year -50%":    PerfYearDown50Percent,
		"year -30%":    PerfYearDown30Percent,
		"year -20%":    PerfYearDown20Percent,
		"year -10%":    PerfYearDown10Percent,
		"year down":    PerfYearDown,
		"year up":      PerfYearUp,
		"year +10%":    PerfYearUp10Percent,
		"year +20%":    PerfYearUp20Percent,
		"year +30%":    PerfYearUp30Percent,
		"year +50%":    PerfYearUp50Percent,
		"year +100%":   PerfYearUp100Percent,
		"year +200%":   PerfYearUp200Percent,
		"year +300%":   PerfYearUp300Percent,
		"year +500%":   PerfYearUp500Percent,
		"ytd -75%":     PerfYTDDown75Percent,
		"ytd -50%":     PerfYTDDown50Percent,
		"ytd -30%":     PerfYTDDown30Percent,
		"ytd -20%":     PerfYTDDown20Percent,
		"ytd -10%":     PerfYTDDown10Percent,
		"ytd -5%":      PerfYTDDown5Percent,
		"ytd down":     PerfYTDDown,
		"ytd up":       PerfYTDUp,
		"ytd +5%":      PerfYTDUp5Percent,
		"ytd +10%":     PerfYTDUp10Percent,
		"ytd +20%":     PerfYTDUp20Percent,
		"ytd +30%":     PerfYTDUp30Percent,
		"ytd +50%":     PerfYTDUp50Percent,
		"ytd +100%":    PerfYTDUp100Percent,
	},
	"performance2": {
		"today up":     PerfTodayUp,
		"today down":   PerfTodayDown,
		"today -15%":   PerfTodayDown15Percent,
		"today -10%":   PerfTodayDown10Percent,
		"today -5%":    PerfTodayDown5Percent,
		"today +5%":    PerfTodayUp5Percent,
		"today +10%":   PerfTodayUp10Percent,
		"today +15%":   PerfTodayUp15Percent,
		"week -30%":    PerfWeekDown30Percent,
		"week -20%":    PerfWeekDown20Percent,
		"week -10%":    PerfWeekDown10Percent,
		"week down":    PerfWeekDown,
		"week up":      PerfWeekUp,
		"week +10%":    PerfWeekUp10Percent,
		"week +20%":    PerfWeekUp20Percent,
		"week +30%":    PerfWeekUp30Percent,
		"month -50%":   PerfMonthDown50Percent,
		"month -30%":   PerfMonthDown30Percent,
		"month -20%":   PerfMonthDown20Percent,
		"month -10%":   PerfMonthDown10Percent,
		"month down":   PerfMonthDown,
		"month up":     PerfMonthUp,
		"month +10%":   PerfMonthUp10Percent,
		"month +20%":   PerfMonthUp20Percent,
		"month +30%":   PerfMonthUp30Percent,
		"month +50%":   PerfMonthUp50Percent,
		"quarter -50%": PerfQuarterDown50Percent,
		"quarter -30%": PerfQuarterDown30Percent,
		"quarter -20%": PerfQuarterDown20Percent,
		"quarter -10%": PerfQuarterDown10Percent,
		"quarter Down": PerfQuarterDown,
		"quarter Up":   PerfQuarterUp,
		"quarter +10%": PerfQuarterUp10Percent,
		"quarter +20%": PerfQuarterUp20Percent,
		"quarter +30%": PerfQuarterUp30Percent,
		"quarter +50%": PerfQuarterUp50Percent,
		"half -75%":    PerfSixMonthsDown75Percent,
		"half -50%":    PerfSixMonthsDown50Percent,
		"half -30%":    PerfSixMonthsDown30Percent,
		"half -20%":    PerfSixMonthsDown20Percent,
		"half -10%":    PerfSixMonthsDown10Percent,
		"half down":    PerfSixMonthsDown,
		"half up":      PerfSixMonthsUp,
		"half +10%":    PerfSixMonthsUp10Percent,
		"half +20%":    PerfSixMonthsUp20Percent,
		"half +30%":    PerfSixMonthsUp30Percent,
		"half +50%":    PerfSixMonthsUp50Percent,
		"half +100%":   PerfSixMonthsUp100Percent,
		"year -75%":    PerfYearDown75Percent,
		"year -50%":    PerfYearDown50Percent,
		"year -30%":    PerfYearDown30Percent,
		"year -20%":    PerfYearDown20Percent,
		"year -10%":    PerfYearDown10Percent,
		"year down":    PerfYearDown,
		"year up":      PerfYearUp,
		"year +10%":    PerfYearUp10Percent,
		"year +20%":    PerfYearUp20Percent,
		"year +30%":    PerfYearUp30Percent,
		"year +50%":    PerfYearUp50Percent,
		"year +100%":   PerfYearUp100Percent,
		"year +200%":   PerfYearUp200Percent,
		"year +300%":   PerfYearUp300Percent,
		"year +500%":   PerfYearUp500Percent,
		"ytd -75%":     PerfYTDDown75Percent,
		"ytd -50%":     PerfYTDDown50Percent,
		"ytd -30%":     PerfYTDDown30Percent,
		"ytd -20%":     PerfYTDDown20Percent,
		"ytd -10%":     PerfYTDDown10Percent,
		"ytd -5%":      PerfYTDDown5Percent,
		"ytd down":     PerfYTDDown,
		"ytd up":       PerfYTDUp,
		"ytd +5%":      PerfYTDUp5Percent,
		"ytd +10%":     PerfYTDUp10Percent,
		"ytd +20%":     PerfYTDUp20Percent,
		"ytd +30%":     PerfYTDUp30Percent,
		"ytd +50%":     PerfYTDUp50Percent,
		"ytd +100%":    PerfYTDUp100Percent,
	},
	"volatility": {
		"week - over 3%":   VolWeekOver3,
		"week - over 4%":   VolWeekOver4,
		"week - over 5%":   VolWeekOver5,
		"week - over 6%":   VolWeekOver6,
		"week - over 7%":   VolWeekOver7,
		"week - over 8%":   VolWeekOver8,
		"week - over 9%":   VolWeekOver9,
		"week - over 10%":  VolWeekOver10,
		"week - over 12%":  VolWeekOver12,
		"week - over 15%":  VolWeekOver15,
		"month - over 2%":  VolMonthOver2,
		"month - over 3%":  VolMonthOver3,
		"month - over 4%":  VolMonthOver4,
		"month - over 5%":  VolMonthOver5,
		"month - over 6%":  VolMonthOver6,
		"month - over 7%":  VolMonthOver7,
		"month - over 8%":  VolMonthOver8,
		"month - over 9%":  VolMonthOver9,
		"month - over 10%": VolMonthOver10,
		"month - over 12%": VolMonthOver12,
		"month - over 15%": VolMonthOver15,
	},
	"rsi": {
		"overbought (90)":      Overbought90,
		"overbought (80)":      Overbought80,
		"overbought (70)":      Overbought70,
		"overbought (60)":      Overbought60,
		"oversold (40)":        Oversold40,
		"oversold (30)":        Oversold30,
		"oversold (20)":        Oversold20,
		"oversold (10)":        Oversold10,
		"not overbought (<60)": NotOverboughtUnder60,
		"not overbought (<50)": NotOverboughtUnder50,
		"not oversold (>50)":   NotOversoldOver50,
		"not oversold (>40)":   NotOversoldOver40,
	},
	"gap": {
		"up":       GapUp,
		"up 0%":    GapUp0Percent,
		"up 1%":    GapUp1Percent,
		"up 2%":    GapUp2Percent,
		"up 3%":    GapUp3Percent,
		"up 4%":    GapUp4Percent,
		"up 5%":    GapUp5Percent,
		"up 6%":    GapUp6Percent,
		"up 7%":    GapUp7Percent,
		"up 8%":    GapUp8Percent,
		"up 9%":    GapUp9Percent,
		"up 10%":   GapUp10Percent,
		"up 15%":   GapUp15Percent,
		"up 20%":   GapUp20Percent,
		"down":     GapDown,
		"down 0%":  GapDown0Percent,
		"down 1%":  GapDown1Percent,
		"down 2%":  GapDown2Percent,
		"down 3%":  GapDown3Percent,
		"down 4%":  GapDown4Percent,
		"down 5%":  GapDown5Percent,
		"down 6%":  GapDown6Percent,
		"down 7%":  GapDown7Percent,
		"down 8%":  GapDown8Percent,
		"down 9%":  GapDown9Percent,
		"down 10%": GapDown10Percent,
		"down 15%": GapDown15Percent,
		"down 20%": GapDown20Percent,
	},
	"sma 20": {
		"price below sma20":          Mavg20PriceBelowSMA20,
		"price 10% below sma20":      Mavg20Price10PercentBelowSMA20,
		"price 20% below sma20":      Mavg20Price20PercentBelowSMA20,
		"price 30% below sma20":      Mavg20Price30PercentBelowSMA20,
		"price 40% below sma20":      Mavg20Price40PercentBelowSMA20,
		"price 50% below sma20":      Mavg20Price50PercentBelowSMA20,
		"price above sma20":          Mavg20PriceAboveSMA20,
		"price 10% above sma20":      Mavg20Price10PercentAboveSMA20,
		"price 20% above sma20":      Mavg20Price20PercentAboveSMA20,
		"price 30% above sma20":      Mavg20Price30PercentAboveSMA20,
		"price 40% above sma20":      Mavg20Price40PercentAboveSMA20,
		"price 50% above sma20":      Mavg20Price50PercentAboveSMA20,
		"price crossed sma20":        Mavg20PriceCrossedSMA20,
		"price crossed sma20 above":  Mavg20PriceCrossedSMA20Above,
		"price crossed sma20 below":  Mavg20PriceCrossedSMA20Below,
		"sma20 crossed sma50":        Mavg20SMA20CrossedSMA50,
		"sma20 crossed sma50 above":  Mavg20SMA20CrossedSMA50Above,
		"sma20 crossed sma50 below":  Mavg20SMA20CrossedSMA50Below,
		"sma20 crossed sma200":       Mavg20SMA20CrossedSMA200,
		"sma20 crossed sma200 above": Mavg20SMA20CrossedSMA200Above,
		"sma20 crossed sma200 below": Mavg20SMA20CrossedSMA200Below,
		"sma20 above sma50":          Mavg20SMA20AboveSMA50,
		"sma20 below sma50":          Mavg20SMA20BelowSMA50,
		"sma20 above sma200":         Mavg20SMA20AboveSMA200,
		"sma20 below sma200":         Mavg20SMA20BelowSMA200,
	},
	"sma 50": {
		"price below sma50":          Mavg50PriceBelowSMA50,
		"price 10% below sma50":      Mavg50Price10PercentBelowSMA50,
		"price 20% below sma50":      Mavg50Price20PercentBelowSMA50,
		"price 30% below sma50":      Mavg50Price30PercentBelowSMA50,
		"price 40% below sma50":      Mavg50Price40PercentBelowSMA50,
		"price 50% below sma50":      Mavg50Price50PercentBelowSMA50,
		"price above sma50":          Mavg50PriceAboveSMA50,
		"price 10% above sma50":      Mavg50Price10PercentAboveSMA50,
		"price 20% above sma50":      Mavg50Price20PercentAboveSMA50,
		"price 30% above sma50":      Mavg50Price30PercentAboveSMA50,
		"price 40% above sma50":      Mavg50Price40PercentAboveSMA50,
		"price 50% above sma50":      Mavg50Price50PercentAboveSMA50,
		"price crossed sma50":        Mavg50PriceCrossedSMA50,
		"price crossed sma50 Above":  Mavg50PriceCrossedSMA50Above,
		"price crossed sma50 Below":  Mavg50PriceCrossedSMA50Below,
		"sma50 crossed sma20":        Mavg50SMA50CrossedSMA20,
		"sma50 crossed sma20 Above":  Mavg50SMA50CrossedSMA20Above,
		"sma50 crossed sma20 Below":  Mavg50SMA50CrossedSMA20Below,
		"sma50 crossed sma200":       Mavg50SMA50CrossedSMA200,
		"sma50 crossed sma200 Above": Mavg50SMA50CrossedSMA200Above,
		"sma50 crossed sma200 Below": Mavg50SMA50CrossedSMA200Below,
		"sma50 Above sma20":          Mavg50SMA50AboveSMA20,
		"sma50 below sma20":          Mavg50SMA50BelowSMA20,
		"sma50 above sma200":         Mavg50SMA50AboveSMA200,
		"sma50 below sma200":         Mavg50SMA50BelowSMA200,
	},
	"sma 200": {
		"price below sma200":         Mavg200PriceBelowSMA200,
		"price 10% below sma200":     Mavg200Price10PercentBelowSMA200,
		"price 20% below sma200":     Mavg200Price20PercentBelowSMA200,
		"price 30% below sma200":     Mavg200Price30PercentBelowSMA200,
		"price 40% below sma200":     Mavg200Price40PercentBelowSMA200,
		"price 50% below sma200":     Mavg200Price50PercentBelowSMA200,
		"price 60% below sma200":     Mavg200Price60PercentBelowSMA200,
		"price 70% below sma200":     Mavg200Price70PercentBelowSMA200,
		"price 80% below sma200":     Mavg200Price80PercentBelowSMA200,
		"price 90% below sma200":     Mavg200Price90PercentBelowSMA200,
		"price above sma200":         Mavg200PriceAboveSMA200,
		"price 10% above sma200":     Mavg200Price10PercentAboveSMA200,
		"price 20% above sma200":     Mavg200Price20PercentAboveSMA200,
		"price 30% above sma200":     Mavg200Price30PercentAboveSMA200,
		"price 40% above sma200":     Mavg200Price40PercentAboveSMA200,
		"price 50% above sma200":     Mavg200Price50PercentAboveSMA200,
		"price 60% above sma200":     Mavg200Price60PercentAboveSMA200,
		"price 70% above sma200":     Mavg200Price70PercentAboveSMA200,
		"price 80% above sma200":     Mavg200Price80PercentAboveSMA200,
		"price 90% above sma200":     Mavg200Price90PercentAboveSMA200,
		"price 100% above sma200":    Mavg200Price100PercentAboveSMA200,
		"price crossed sma200":       Mavg200PriceCrossedSMA200,
		"price crossed sma200 above": Mavg200PriceCrossedSMA200Above,
		"price crossed sma200 below": Mavg200PriceCrossedSMA200Below,
		"sma200 crossed sma20":       Mavg200SMA200CrossedSMA20,
		"sma200 crossed sma20 above": Mavg200SMA200CrossedSMA20Above,
		"sma200 crossed sma20 below": Mavg200SMA200CrossedSMA20Below,
		"sma200 crossed sma50":       Mavg200SMA200CrossedSMA50,
		"sma200 crossed sma50 above": Mavg200SMA200CrossedSMA50Above,
		"sma200 crossed sma50 below": Mavg200SMA200CrossedSMA50Below,
		"sma200 above sma20":         Mavg200SMA200AboveSMA20,
		"sma200 below sma20":         Mavg200SMA200BelowSMA20,
		"sma200 above sma50":         Mavg200SMA200AboveSMA50,
		"sma200 below sma50":         Mavg200SMA200BelowSMA50,
	},
	"change": {
		"up":       ChangeUp,
		"up 1%":    ChangeUp1Percent,
		"up 2%":    ChangeUp2Percent,
		"up 3%":    ChangeUp3Percent,
		"up 4%":    ChangeUp4Percent,
		"up 5%":    ChangeUp5Percent,
		"up 6%":    ChangeUp6Percent,
		"up 7%":    ChangeUp7Percent,
		"up 8%":    ChangeUp8Percent,
		"up 9%":    ChangeUp9Percent,
		"up 10%":   ChangeUp10Percent,
		"up 15%":   ChangeUp15Percent,
		"up 20%":   ChangeUp20Percent,
		"down":     ChangeDown,
		"down 1%":  ChangeDown1Percent,
		"down 2%":  ChangeDown2Percent,
		"down 3%":  ChangeDown3Percent,
		"down 4%":  ChangeDown4Percent,
		"down 5%":  ChangeDown5Percent,
		"down 6%":  ChangeDown6Percent,
		"down 7%":  ChangeDown7Percent,
		"down 8%":  ChangeDown8Percent,
		"down 9%":  ChangeDown9Percent,
		"down 10%": ChangeDown10Percent,
		"down 15%": ChangeDown15Percent,
		"down 20%": ChangeDown20Percent,
	},
	"change from open": {
		"up":       ChangeUp,
		"up 1%":    ChangeUp1Percent,
		"up 2%":    ChangeUp2Percent,
		"up 3%":    ChangeUp3Percent,
		"up 4%":    ChangeUp4Percent,
		"up 5%":    ChangeUp5Percent,
		"up 6%":    ChangeUp6Percent,
		"up 7%":    ChangeUp7Percent,
		"up 8%":    ChangeUp8Percent,
		"up 9%":    ChangeUp9Percent,
		"up 10%":   ChangeUp10Percent,
		"up 15%":   ChangeUp15Percent,
		"up 20%":   ChangeUp20Percent,
		"down":     ChangeDown,
		"down 1%":  ChangeDown1Percent,
		"down 2%":  ChangeDown2Percent,
		"down 3%":  ChangeDown3Percent,
		"down 4%":  ChangeDown4Percent,
		"down 5%":  ChangeDown5Percent,
		"down 6%":  ChangeDown6Percent,
		"down 7%":  ChangeDown7Percent,
		"down 8%":  ChangeDown8Percent,
		"down 9%":  ChangeDown9Percent,
		"down 10%": ChangeDown10Percent,
		"down 15%": ChangeDown15Percent,
		"down 20%": ChangeDown20Percent,
	},
	"20 day high/low": {
		"new high":               DHLNewHigh,
		"new low":                DHLNewLow,
		"5% or more below high":  DHLAtLeast5PercentBelowHigh,
		"10% or more below high": DHLAtLeast10PercentBelowHigh,
		"15% or more below high": DHLAtLeast15PercentBelowHigh,
		"20% or more below high": DHLAtLeast20PercentBelowHigh,
		"30% or more below high": DHLAtLeast30PercentBelowHigh,
		"40% or more below high": DHLAtLeast40PercentBelowHigh,
		"50% or more below high": DHLAtLeast50PercentBelowHigh,
		"0-3% below high":        DHL0to3PercentBelowHigh,
		"0-5% below high":        DHL0to5PercentBelowHigh,
		"0-10% below high":       DHL0to10PercentBelowHigh,
		"5% or more above low":   DHLAtLeast5PercentAboveLow,
		"10% or more above low":  DHLAtLeast10PercentAboveLow,
		"15% or more above low":  DHLAtLeast15PercentAboveLow,
		"20% or more above low":  DHLAtLeast20PercentAboveLow,
		"30% or more above low":  DHLAtLeast30PercentAboveLow,
		"40% or more above low":  DHLAtLeast40PercentAboveLow,
		"50% or more above low":  DHLAtLeast50PercentAboveLow,
		"0-3% above low":         DHL0to3PercentAboveLow,
		"0-5% above low":         DHL0to5PercentAboveLow,
		"0-10% above low":        DHL0to10PercentAboveLow,
	},
	"50 day high/low": {
		"new high":               DHLNewHigh,
		"new low":                DHLNewLow,
		"5% or more below high":  DHLAtLeast5PercentBelowHigh,
		"10% or more below high": DHLAtLeast10PercentBelowHigh,
		"15% or more below high": DHLAtLeast15PercentBelowHigh,
		"20% or more below high": DHLAtLeast20PercentBelowHigh,
		"30% or more below high": DHLAtLeast30PercentBelowHigh,
		"40% or more below high": DHLAtLeast40PercentBelowHigh,
		"50% or more below high": DHLAtLeast50PercentBelowHigh,
		"0-3% below high":        DHL0to3PercentBelowHigh,
		"0-5% below high":        DHL0to5PercentBelowHigh,
		"0-10% below high":       DHL0to10PercentBelowHigh,
		"5% or more above low":   DHLAtLeast5PercentAboveLow,
		"10% or more above low":  DHLAtLeast10PercentAboveLow,
		"15% or more above low":  DHLAtLeast15PercentAboveLow,
		"20% or more above low":  DHLAtLeast20PercentAboveLow,
		"30% or more above low":  DHLAtLeast30PercentAboveLow,
		"40% or more above low":  DHLAtLeast40PercentAboveLow,
		"50% or more above low":  DHLAtLeast50PercentAboveLow,
		"0-3% above low":         DHL0to3PercentAboveLow,
		"0-5% above low":         DHL0to5PercentAboveLow,
		"0-10% above low":        DHL0to10PercentAboveLow,
	},
	"52 week high/low": {
		"new high":               WHLNewHigh,
		"new low":                WHLNewLow,
		"5% or more below high":  WHLAtLeast5PercentBelowHigh,
		"10% or more below high": WHLAtLeast10PercentBelowHigh,
		"15% or more below high": WHLAtLeast15PercentBelowHigh,
		"20% or more below high": WHLAtLeast20PercentBelowHigh,
		"30% or more below high": WHLAtLeast30PercentBelowHigh,
		"40% or more below high": WHLAtLeast40PercentBelowHigh,
		"50% or more below high": WHLAtLeast50PercentBelowHigh,
		"60% or more below high": WHLAtLeast60PercentBelowHigh,
		"70% or more below high": WHLAtLeast70PercentBelowHigh,
		"80% or more below high": WHLAtLeast80PercentBelowHigh,
		"90% or more below high": WHLAtLeast90PercentBelowHigh,
		"0-3% below high":        WHL0to3PercentBelowHigh,
		"0-5% below high":        WHL0to5PercentBelowHigh,
		"0-10% below high":       WHL0to10PercentBelowHigh,
		"5% or more above low":   WHLAtLeast5PercentAboveLow,
		"10% or more above low":  WHLAtLeast10PercentAboveLow,
		"15% or more above low":  WHLAtLeast15PercentAboveLow,
		"20% or more above low":  WHLAtLeast20PercentAboveLow,
		"30% or more above low":  WHLAtLeast30PercentAboveLow,
		"40% or more above low":  WHLAtLeast40PercentAboveLow,
		"50% or more above low":  WHLAtLeast50PercentAboveLow,
		"60% or more above low":  WHLAtLeast60PercentAboveLow,
		"70% or more above low":  WHLAtLeast70PercentAboveLow,
		"80% or more above low":  WHLAtLeast80PercentAboveLow,
		"90% or more above low":  WHLAtLeast90PercentAboveLow,
		"100% or more above low": WHLAtLeast100PercentAboveLow,
		"120% or more above low": WHLAtLeast120PercentAboveLow,
		"150% or more above low": WHLAtLeast150PercentAboveLow,
		"200% or more above low": WHLAtLeast200PercentAboveLow,
		"300% or more above low": WHLAtLeast300PercentAboveLow,
		"500% or more above low": WHLAtLeast500PercentAboveLow,
		"0-3% above low":         WHL0to3PercentAboveLow,
		"0-5% above low":         WHL0to5PercentAboveLow,
		"0-10% above low":        WHL0to10PercentAboveLow,
	},
	"pattern": {
		"horizontal s/r":               PatternHorizontalSR,
		"horizontal s/r (strong)":      PatternHorizontalSRStrong,
		"tl resistance":                PatternTLResistance,
		"tl resistance (strong)":       PatternTLResistanceStrong,
		"tl support":                   PatternTLSupport,
		"tl support (strong)":          PatternTLSupportStrong,
		"wedge up":                     PatternWedgeUp,
		"wedge up (strong)":            PatternWedgeUpStrong,
		"wedge down":                   PatternWedgeDown,
		"wedge down (strong)":          PatternWedgeDownStrong,
		"triangle ascending":           PatternTriangleAscending,
		"triangle ascending (strong)":  PatternTriangleAscendingStrong,
		"triangle descending":          PatternTriangleDescending,
		"triangle descending (strong)": PatternTriangleDescendingStrong,
		"wedge":                        PatternWedge,
		"wedge (strong)":               PatternWedgeStrong,
		"channel up":                   PatternChannelUp,
		"channel up (strong)":          PatternChannelUpStrong,
		"channel down":                 PatternChannelDown,
		"channel down (strong)":        PatternChannelDownStrong,
		"channel":                      PatternChannel,
		"channel (strong)":             PatternChannelStrong,
		"double top":                   PatternDoubleTop,
		"double bottom":                PatternDoubleBottom,
		"multiple top":                 PatternMultipleTop,
		"multiple bottom":              PatternMultipleBottom,
		"head & shoulders":             PatternHeadAndShoulders,
		"head & shoulders inverse":     PatternHeadAndShouldersInverse,
	},
	"candlestick": {
		"long lower shadow":  CSLongLowerShadow,
		"long upper shadow":  CSLongUpperShadow,
		"hammer":             CSHammer,
		"inverted hammer":    CSInvertedHammer,
		"spinning top white": CSSpinningTopWhite,
		"spinning top black": CSSpinningTopBlack,
		"doji":               CSDoji,
		"dragonfly doji":     CSDragonflyDoji,
		"gravestone doji":    CSGravestoneDoji,
		"marubozu white":     CSMarubozuWhite,
		"marubozu black":     CSMarubozuBlack,
	},
	"beta": {
		"under 0":    BetaUnder0,
		"under 0.5":  BetaUnder0Point5,
		"under 1":    BetaUnder1,
		"under 1.5":  BetaUnder1Point5,
		"under 2":    BetaUnder2,
		"over 0":     BetaOver0,
		"over 0.5":   BetaOver0Point5,
		"over 1":     BetaOver1,
		"over 1.5":   BetaOver1Point5,
		"over 2":     BetaOver2,
		"over 2.5":   BetaOver2Point5,
		"over 3":     BetaOver3,
		"over 4":     BetaOver4,
		"0 to 0.5":   Beta0to0Point5,
		"0 to 1":     Beta0to1,
		"0.5 to 1":   Beta0Point5to1,
		"0.5 to 1.5": Beta0Point5to1Point5,
		"1 to 1.5":   Beta1to1Point5,
		"1 to 2":     Beta1to2,
	},
	"average true range": {
		"over 0.25":  ATROver0Point25,
		"over 0.5":   ATROver0Point5,
		"over 0.75":  ATROver0Point75,
		"over 1":     ATROver1,
		"over 1.5":   ATROver1Point5,
		"over 2":     ATROver2,
		"over 2.5":   ATROver2Point5,
		"over 3":     ATROver3,
		"over 3.5":   ATROver3Point5,
		"over 4":     ATROver4,
		"over 4.5":   ATROver4Point5,
		"over 5":     ATROver5,
		"under 0.25": ATRUnder0Point25,
		"under 0.5":  ATRUnder0Point5,
		"under 0.75": ATRUnder0Point75,
		"under 1":    ATRUnder1,
		"under 1.5":  ATRUnder1Point5,
		"under 2":    ATRUnder2,
		"under 2.5":  ATRUnder2Point5,
		"under 3":    ATRUnder3,
		"under 3.5":  ATRUnder3Point5,
		"under 4":    ATRUnder4,
		"under 4.5":  ATRUnder4Point5,
		"under 5":    ATRUnder5,
	},
}
