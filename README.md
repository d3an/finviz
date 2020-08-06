# Unofficial Finviz API

## Installation

Run `go get github.com/d3an/finviz` to install the package.

## Documentation


#### Descriptive Filters

##### Exchange
Filter: `ExchangeFilter`

Values:

- `AMEX`
- `NASDAQ`
- `NYSE`

Example: `ExchangeFilter{AMEX}`

##### Index
Filter: `IndexFilter`

Values:

- `SP500`
- `DJIA`

Example: `IndexFilter{SP500}`

##### Sector
Filter: `SectorFilter`

Values:

- `BasicMaterials`
- `CommunicationServices`
- `ConsumerCyclical`
- `ConsumerDefensive`
- `Energy`
- `Financial`
- `Healthcare`
- `Industrials`
- `RealEstate`
- `Technology`
- `Utilities`

Example: `SectorFilter{Technology}`

##### Industry
Filter: `IndustryFilter`

Values:

- `StocksOnly`
- `ETF`
- `AdvertisingAgencies`
- `AerospaceAndDefense`
- `AgriculturalInputs`
- `Airlines`
- `AirportsAirServices`
- `Aluminum`
- `ApparelManufacturing`
- `ApparelRetail`
- `AssetManagement`
- `AutoManufacturers`
- `AutoParts`
- `AutoTruckDealerships`
- `BanksDiversified`
- `BanksRegional`
- `BeveragesBrewers`
- `BeveragesNonAlcoholic`
- `BeveragesWineriesDistilleries`
- `Biotechnology`
- `Broadcasting`
- `BuildingMaterials`
- `BuildingProductsEquipment`
- `BusinessEquipmentSupplies`
- `CapitalMarkets`
- `Chemicals`
- `ClosedEndFundDebt`
- `ClosedEndFundEquity`
- `ClosedEndFundForeign`
- `CokingCoal`
- `CommunicationsEquipment`
- `ComputerHardware`
- `Confectioners`
- `Conglomerates`
- `ConsultingServices`
- `ConsumerElectronics`
- `Copper`
- `CreditServices`
- `DepartmentStores`
- `DiagnosticsResearch`
- `DiscountStores`
- `DrugManufacturersGeneral`
- `DrugManufacturersSpecialtyGeneric`
- `EducationTrainingServices`
- `ElectricalEquipmentParts`
- `ElectronicComponents`
- `ElectronicGamingMultimedia`
- `ElectronicsComputerDistribution`
- `EngineeringConstruction`
- `Entertainment`
- `FarmHeavyConstructionMachinery`
- `FarmProducts`
- `FinancialConglomerates`
- `FinancialDataStockExchanges`
- `FoodDistribution`
- `FootwearAccessories`
- `FurnishingsFixturesAppliances`
- `Gambling`
- `Gold`
- `GroceryStores`
- `HealthcarePlans`
- `HealthInformationServices`
- `HomeImprovementRetail`
- `HouseholdPersonalProducts`
- `IndustrialDistribution`
- `InformationTechnologyServices`
- `InfrastructureOperations`
- `InsuranceBrokers`
- `InsuranceDiversified`
- `InsuranceLife`
- `InsurancePropertyCasualty`
- `InsuranceReinsurance`
- `InsuranceSpecialty`
- `IntegratedFreightLogistics`
- `InternetContentInformation`
- `InternetRetail`
- `Leisure`
- `Lodging`
- `LumberWoodProduction`
- `LuxuryGoods`
- `MarineShipping`
- `MedicalCareFacilities`
- `MedicalDevices`
- `MedicalDistribution`
- `MedicalInstrumentsSupplies`
- `MetalFabrication`
- `MortgageFinance`
- `OilGasDrilling`
- `OilGasEP`
- `OilGasEquipmentServices`
- `OilGasIntegrated`
- `OilGasMidstream`
- `OilGasRefiningMarketing`
- `OtherIndustrialMetalsMining`
- `OtherPreciousMetalsMining`
- `PackagedFoods`
- `PackagingContainers`
- `PaperPaperProducts`
- `PersonalServices`
- `PharmaceuticalRetailers`
- `PollutionTreatmentControls`
- `Publishing`
- `Railroads`
- `RealEstateDevelopment`
- `RealEstateDiversified`
- `RealEstateServices`
- `RecreationalVehicles`
- `REITDiversified`
- `REITHealthcareFacilities`
- `REITHotelMotel`
- `REITIndustrial`
- `REITMortgage`
- `REITOffice`
- `REITResidential`
- `REITRetail`
- `REITSpecialty`
- `RentalLeasingServices`
- `ResidentialConstruction`
- `ResortsCasinos`
- `Restaurants`
- `ScientificTechnicalInstruments`
- `SecurityProtectionServices`
- `SemiconductorEquipmentMaterials`
- `Semiconductors`
- `ShellCompanies`
- `Silver`
- `SoftwareApplication`
- `SoftwareInfrastructure`
- `Solar`
- `SpecialtyBusinessServices`
- `SpecialtyChemicals`
- `SpecialtyIndustrialMachinery`
- `SpecialtyRetail`
- `StaffingEmploymentServices`
- `Steel`
- `TelecomServices`
- `TextileManufacturing`
- `ThermalCoal`
- `Tobacco`
- `ToolsAccessories`
- `TravelServices`
- `Trucking`
- `Uranium`
- `UtilitiesDiversified`
- `UtilitiesIndependentPowerProducers`
- `UtilitiesRegulatedElectric`
- `UtilitiesRegulatedGas`
- `UtilitiesRegulatedWater`
- `UtilitiesRenewable`
- `WasteManagement`

Example: `IndustryFilter{UtilitiesRegulatedGas}`

##### Country
Filter: `CountryFilter`

Values:

- `USA`
- `NotUSA`
- `Asia`
- `Europe`
- `LatinAmerica`
- `BRIC`
- `Argentina`
- `Australia`
- `Bahamas`
- `Belgium`
- `BeNeLux`
- `Bermuda`
- `Brazil`
- `Canada`
- `CaymanIslands`
- `Chile`
- `China`
- `ChinaHongKong`
- `Colombia`
- `Cyprus`
- `Denmark`
- `Finland`
- `France`
- `Germany`
- `Greece`
- `HongKong`
- `Hungary`
- `Iceland`
- `India`
- `Indonesia`
- `Ireland`
- `Israel`
- `Italy`
- `Japan`
- `Kazakhstan`
- `Luxembourg`
- `Malaysia`
- `Malta`
- `Mexico`
- `Monaco`
- `Netherlands`
- `NewZealand`
- `Norway`
- `Panama`
- `Peru`
- `Philippines`
- `Portugal`
- `Russia`
- `Singapore`
- `SouthAfrica`
- `SouthKorea`
- `Spain`
- `Sweden`
- `Switzerland`
- `Taiwan`
- `Turkey`
- `UAE`
- `UK`
- `Uruguay`

Example: `CountryFilter{Canada}`

##### Market Cap.
Filter: `MarketCapFilter`

Values:

- `MegaOver200B`
- `Large10to200B`
- `Mid2to10B`
- `Small300Mto2B`
- `Micro50to300M`
- `NanoUnder50M`
- `LargeOver10B`
- `MidOver2B`
- `SmallOver300M`
- `MicroOver50M`
- `LargeUnder200B`
- `MidUnder10B`
- `SmallUnder2B`
- `MicroUnder300M`

Example: `MarketCapFilter{MegaOver200B}`

##### Dividend Yield
Filter: `DividendYieldFilter`

Values:

- `DYNone`
- `DYPositive`
- `DYHigh`
- `DYVeryHigh`
- `DYOver1`
- `DYOver2`
- `DYOver3`
- `DYOver4`
- `DYOver5`
- `DYOver6`
- `DYOver7`
- `DYOver8`
- `DYOver9`
- `DYOver10`

Example: `DividendYieldFilter{DYNone}`

##### Float Short
Filter: `FloatShortFilter`

Values:

- `FSLow`
- `FSHigh`
- `FSUnder5`
- `FSUnder10`
- `FSUnder15`
- `FSUnder20`
- `FSUnder25`
- `FSUnder30`
- `FSOver5`
- `FSOver10`
- `FSOver15`
- `FSOver20`
- `FSOver25`
- `FSOver30`

Example: `FloatShortFilter{FSOver5}`

##### Analyst Recommendation
Filter: `RecommendationFilter`

Values:

- `StrongBuy`
- `BuyBetter`
- `Buy`
- `HoldBetter`
- `Hold`
- `HoldWorse`
- `Sell`
- `SellWorse`
- `StrongSell`

Example: `RecommendationFilter{StrongBuy}`

##### Optionable/Shortable
Filter: `OptionShortFilter`

Values:

- `Option`
- `Short`
- `OptionShort`

Example: `OptionShortFilter{Option}`

##### Earnings Date
Filter: `EarningsDateFilter`

Values:

- `EDToday`
- `EDTodayBeforeMarketOpen`
- `EDTodayAfterMarketClose`
- `EDTomorrow`
- `EDTomorrowBeforeMarketOpen`
- `EDTomorrowAfterMarketClose`
- `EDYesterday`
- `EDYesterdayBeforeMarketOpen`
- `EDYesterdayAfterMarketClose`
- `EDNext5Days`
- `EDPrevious5Days`
- `EDThisWeek`
- `EDNextWeek`
- `EDPrevWeek`
- `EDThisMonth`

Example: `EarningsDateFilter{EDNext5Days}`

##### Average Volume
Filter: `AverageVolumeFilter`

Values:

- `AvgVolUnder50K`
- `AvgVolUnder100K`
- `AvgVolUnder500K`
- `AvgVolUnder750K`
- `AvgVolUnder1M`
- `AvgVolOver50K`
- `AvgVolOver100K`
- `AvgVolOver200K`
- `AvgVolOver300K`
- `AvgVolOver400K`
- `AvgVolOver500K`
- `AvgVolOver750K`
- `AvgVolOver1M`
- `AvgVolOver2M`
- `AvgVolBetween100to500K`
- `AvgVolBetween100Kto1M`
- `AvgVolBetween500Kto1M`
- `AvgVolBetween500Kto10M`

Example: `AverageVolumeFilter{AvgVolOver1M}`

##### Relative Volume
Filter: `RelativeVolumeFilter`

Values:

- `RVOver10`
- `RVOver5`
- `RVOver3`
- `RVOver2`
- `RVOver1point5`
- `RVOver1`
- `RVOver0point75`
- `RVOver0point5`
- `RVOver0point25`
- `RVUnder2`
- `RVUnder1point5`
- `RVUnder1`
- `RVUnder0point75`
- `RVUnder0point5`
- `RVUnder0point25`
- `RVUnder0point1`

Example: `RelativeVolumeFilter{RVOver5}`

##### Current Volume
Filter: `CurrentVolumeFilter`

Values:

- `CurVolUnder50K`
- `CurVolUnder100K`
- `CurVolUnder500K`
- `CurVolUnder750K`
- `CurVolUnder1M`
- `CurVolOver0`
- `CurVolOver50K`
- `CurVolOver100K`
- `CurVolOver200K`
- `CurVolOver300K`
- `CurVolOver400K`
- `CurVolOver500K`
- `CurVolOver750K`
- `CurVolOver1M`
- `CurVolOver2M`
- `CurVolOver5M`
- `CurVolOver10M`
- `CurVolOver20M`

Example: `CurrentVolumeFilter{CurVolOver1M}`

##### Price
Filter: `PriceFilter`

Values:

- `PriceUnder1`
- `PriceUnder2`
- `PriceUnder3`
- `PriceUnder4`
- `PriceUnder5`
- `PriceUnder7`
- `PriceUnder10`
- `PriceUnder15`
- `PriceUnder20`
- `PriceUnder30`
- `PriceUnder40`
- `PriceUnder50`
- `PriceOver1`
- `PriceOver2`
- `PriceOver3`
- `PriceOver4`
- `PriceOver5`
- `PriceOver7`
- `PriceOver10`
- `PriceOver15`
- `PriceOver20`
- `PriceOver30`
- `PriceOver40`
- `PriceOver50`
- `PriceOver60`
- `PriceOver70`
- `PriceOver80`
- `PriceOver90`
- `PriceOver100`
- `Price1to5`
- `Price1to10`
- `Price1to20`
- `Price5to10`
- `Price5to20`
- `Price5to50`
- `Price10to20`
- `Price10to50`
- `Price20to50`
- `Price50to100`

Example: `PriceFilter{PriceOver5}`

##### Target Price
Filter: `TargetPriceFilter`

Values:

- `TargetAbovePriceBy50Percent`
- `TargetAbovePriceBy40Percent`
- `TargetAbovePriceBy30Percent`
- `TargetAbovePriceBy20Percent`
- `TargetAbovePriceBy10Percent`
- `TargetAbovePriceBy5Percent`
- `TargetAbovePrice`
- `TargetBelowPrice`
- `TargetBelowPriceBy50Percent`
- `TargetBelowPriceBy40Percent`
- `TargetBelowPriceBy30Percent`
- `TargetBelowPriceBy20Percent`
- `TargetBelowPriceBy10Percent`
- `TargetBelowPriceBy5Percent`

Example: `TargetPriceFilter{TargetBelowPrice}`

##### IPO Date
Filter: `IPODateFilter`

Values:

- `IDToday`
- `IDYesterday`
- `IDPreviousWeek`
- `IDPreviousMonth`
- `IDPreviousQuarter`
- `IDPreviousYear`
- `IDPrevious2Years`
- `IDPrevious3Years`
- `IDPrevious5Years`
- `IDMoreThan1YearAgo`
- `IDMoreThan5YearsAgo`
- `IDMoreThan10YearsAgo`
- `IDMoreThan15YearsAgo`
- `IDMoreThan20YearsAgo`
- `IDMoreThan25YearsAgo`

Example: `IPODateFilter{IDPrevious5Years}`

##### Shares Outstanding
Filter: `SharesOutstandingFilter`

Values:

- `SOUnder1M`
- `SOUnder5M`
- `SOUnder10M`
- `SOUnder20M`
- `SOUnder50M`
- `SOUnder100M`
- `SOOver1M`
- `SOOver2M`
- `SOOver5M`
- `SOOver10M`
- `SOOver20M`
- `SOOver50M`
- `SOOver100M`
- `SOOver200M`
- `SOOver500M`
- `SOOver1000M`

Example: `SharesOutstandingFilter{SOOver1M}`

##### Float Shares
Filter: `FloatFilter`

Values:

- `SOUnder1M`
- `SOUnder5M`
- `SOUnder10M`
- `SOUnder20M`
- `SOUnder50M`
- `SOUnder100M`
- `SOOver1M`
- `SOOver2M`
- `SOOver5M`
- `SOOver10M`
- `SOOver20M`
- `SOOver50M`
- `SOOver100M`
- `SOOver200M`
- `SOOver500M`
- `SOOver1000M`

Example: `FloatFilter{SOOver2M}`

##### P/E
Filter: `PEFilter`

Values:

- `PELow`
- `PEProfitable`
- `PEHigh`
- `PEUnder5`
- `PEUnder10`
- `PEUnder15`
- `PEUnder20`
- `PEUnder25`
- `PEUnder30`
- `PEUnder35`
- `PEUnder40`
- `PEUnder45`
- `PEUnder50`
- `PEOver5`
- `PEOver10`
- `PEOver15`
- `PEOver20`
- `PEOver25`
- `PEOver30`
- `PEOver35`
- `PEOver40`
- `PEOver45`
- `PEOver50`

Example: `PEFilter{PEOver5}`

##### Forward P/E
Filter: `ForwardPEFilter`

Values:

- `PELow`
- `PEProfitable`
- `PEHigh`
- `PEUnder5`
- `PEUnder10`
- `PEUnder15`
- `PEUnder20`
- `PEUnder25`
- `PEUnder30`
- `PEUnder35`
- `PEUnder40`
- `PEUnder45`
- `PEUnder50`
- `PEOver5`
- `PEOver10`
- `PEOver15`
- `PEOver20`
- `PEOver25`
- `PEOver30`
- `PEOver35`
- `PEOver40`
- `PEOver45`
- `PEOver50`

Example: `ForwardPEFilter{PEProfitable}`

##### PEG
Filter: `PEGFilter`

Values:

- `PEGLow`
- `PEGHigh`
- `PEGUnder1`
- `PEGUnder2`
- `PEGUnder3`
- `PEGOver1`
- `PEGOver2`
- `PEGOver3`

Example: `PEGFilter{PEGLow}`

##### Price/Sales
Filter: `PriceSalesFilter`

Values:

- `PSLow`
- `PSHigh`
- `PSUnder1`
- `PSUnder2`
- `PSUnder3`
- `PSUnder4`
- `PSUnder5`
- `PSUnder6`
- `PSUnder7`
- `PSUnder8`
- `PSUnder9`
- `PSUnder10`
- `PSOver1`
- `PSOver2`
- `PSOver3`
- `PSOver4`
- `PSOver5`
- `PSOver6`
- `PSOver7`
- `PSOver8`
- `PSOver9`
- `PSOver10`

Example: `PriceSalesFilter{PSOver1}`

##### Price/Book
Filter: `PriceBookFilter`

Values:

- `PSLow`
- `PSHigh`
- `PSUnder1`
- `PSUnder2`
- `PSUnder3`
- `PSUnder4`
- `PSUnder5`
- `PSUnder6`
- `PSUnder7`
- `PSUnder8`
- `PSUnder9`
- `PSUnder10`
- `PSOver1`
- `PSOver2`
- `PSOver3`
- `PSOver4`
- `PSOver5`
- `PSOver6`
- `PSOver7`
- `PSOver8`
- `PSOver9`
- `PSOver10`

Example: `PriceBookFilter{PSOver1}`

##### Price/Cash
Filter: `PriceCashFilter`

Values:

- `PCLow`
- `PCHigh`
- `PCUnder1`
- `PCUnder2`
- `PCUnder3`
- `PCUnder4`
- `PCUnder5`
- `PCUnder6`
- `PCUnder7`
- `PCUnder8`
- `PCUnder9`
- `PCUnder10`
- `PCOver1`
- `PCOver2`
- `PCOver3`
- `PCOver4`
- `PCOver5`
- `PCOver6`
- `PCOver7`
- `PCOver8`
- `PCOver9`
- `PCOver10`
- `PCOver20`
- `PCOver30`
- `PCOver40`
- `PCOver50`

Example: `PriceCashFilter{PCOver1}`

##### Price/Free Cash Flow
Filter: `PriceFCFFilter`

Values:

- `PFCFLow`
- `PFCFHigh`
- `PFCFUnder5`
- `PFCFUnder10`
- `PFCFUnder15`
- `PFCFUnder20`
- `PFCFUnder25`
- `PFCFUnder30`
- `PFCFUnder35`
- `PFCFUnder40`
- `PFCFUnder45`
- `PFCFUnder50`
- `PFCFUnder60`
- `PFCFUnder70`
- `PFCFUnder80`
- `PFCFUnder90`
- `PFCFUnder100`
- `PFCFOver5`
- `PFCFOver10`
- `PFCFOver15`
- `PFCFOver20`
- `PFCFOver25`
- `PFCFOver30`
- `PFCFOver35`
- `PFCFOver40`
- `PFCFOver45`
- `PFCFOver50`
- `PFCFOver60`
- `PFCFOver70`
- `PFCFOver80`
- `PFCFOver90`
- `PFCFOver100`

Example: `PriceFCFFilter{PFCFOver5}`

##### EPS Growth This Year
Filter: `EPSGrowthThisYearFilter`

Values:

- `GrowthNegative`
- `GrowthPositive`
- `GrowthPositiveLow`
- `GrowthHigh`
- `GrowthUnder5`
- `GrowthUnder10`
- `GrowthUnder15`
- `GrowthUnder20`
- `GrowthUnder25`
- `GrowthUnder30`
- `GrowthOver5`
- `GrowthOver10`
- `GrowthOver15`
- `GrowthOver20`
- `GrowthOver25`
- `GrowthOver30`

Example: `EPSGrowthThisYearFilter{GrowthHigh}`

##### EPS Growth Next Year
Filter: `EPSGrowthNextYearFilter`

Values:

- `GrowthNegative`
- `GrowthPositive`
- `GrowthPositiveLow`
- `GrowthHigh`
- `GrowthUnder5`
- `GrowthUnder10`
- `GrowthUnder15`
- `GrowthUnder20`
- `GrowthUnder25`
- `GrowthUnder30`
- `GrowthOver5`
- `GrowthOver10`
- `GrowthOver15`
- `GrowthOver20`
- `GrowthOver25`
- `GrowthOver30`

Example: `EPSGrowthNextYearFilter{GrowthLow}`

##### EPS Growth Past 5 Years
Filter: `EPSGrowthPast5YearsFilter`

Values:

- `GrowthNegative`
- `GrowthPositive`
- `GrowthPositiveLow`
- `GrowthHigh`
- `GrowthUnder5`
- `GrowthUnder10`
- `GrowthUnder15`
- `GrowthUnder20`
- `GrowthUnder25`
- `GrowthUnder30`
- `GrowthOver5`
- `GrowthOver10`
- `GrowthOver15`
- `GrowthOver20`
- `GrowthOver25`
- `GrowthOver30`

Example: `EPSGrowthPast5YearsFilter{GrowthOver5}`

##### EPS Growth Next 5 Years
Filter: `EPSGrowthNext5YearsFilter`

Values:

- `GrowthNegative`
- `GrowthPositive`
- `GrowthPositiveLow`
- `GrowthHigh`
- `GrowthUnder5`
- `GrowthUnder10`
- `GrowthUnder15`
- `GrowthUnder20`
- `GrowthUnder25`
- `GrowthUnder30`
- `GrowthOver5`
- `GrowthOver10`
- `GrowthOver15`
- `GrowthOver20`
- `GrowthOver25`
- `GrowthOver30`

Example: `EPSGrowthNext5YearsFilter{GrowthUnder5}`

##### Sales Growth Past 5 Years
Filter: `SalesGrowthPast5YearsFilter`

Values:

- `GrowthNegative`
- `GrowthPositive`
- `GrowthPositiveLow`
- `GrowthHigh`
- `GrowthUnder5`
- `GrowthUnder10`
- `GrowthUnder15`
- `GrowthUnder20`
- `GrowthUnder25`
- `GrowthUnder30`
- `GrowthOver5`
- `GrowthOver10`
- `GrowthOver15`
- `GrowthOver20`
- `GrowthOver25`
- `GrowthOver30`

Example: `SalesGrowthPast5YearsFilter{GrowthPositiveLow}`

##### EPS Growth Quarter Over Quarter
Filter: `EPSGrowthQtrOverQtrFilter`

Values:

- `GrowthNegative`
- `GrowthPositive`
- `GrowthPositiveLow`
- `GrowthHigh`
- `GrowthUnder5`
- `GrowthUnder10`
- `GrowthUnder15`
- `GrowthUnder20`
- `GrowthUnder25`
- `GrowthUnder30`
- `GrowthOver5`
- `GrowthOver10`
- `GrowthOver15`
- `GrowthOver20`
- `GrowthOver25`
- `GrowthOver30`

Example: `EPSGrowthQtrOverQtrFilter{GrowthUnder25}`

##### Sales Growth Quarter Over Quarter
Filter: `SalesGrowthQtrOverQtrFilter`

Values:

- `GrowthNegative`
- `GrowthPositive`
- `GrowthPositiveLow`
- `GrowthHigh`
- `GrowthUnder5`
- `GrowthUnder10`
- `GrowthUnder15`
- `GrowthUnder20`
- `GrowthUnder25`
- `GrowthUnder30`
- `GrowthOver5`
- `GrowthOver10`
- `GrowthOver15`
- `GrowthOver20`
- `GrowthOver25`
- `GrowthOver30`

Example: `SalesGrowthQtrOverQtrFilter{GrowthNegative}`

##### Return on Assets
Filter: `ROAFilter`

Values:

- `ReturnPositive`
- `ReturnNegative`
- `ReturnVeryPositive`
- `ReturnVeryNegative`
- `ReturnUnderNeg50`
- `ReturnUnderNeg45`
- `ReturnUnderNeg40`
- `ReturnUnderNeg35`
- `ReturnUnderNeg30`
- `ReturnUnderNeg25`
- `ReturnUnderNeg20`
- `ReturnUnderNeg15`
- `ReturnUnderNeg10`
- `ReturnUnderNeg5`
- `ReturnOver50`
- `ReturnOver45`
- `ReturnOver40`
- `ReturnOver35`
- `ReturnOver30`
- `ReturnOver25`
- `ReturnOver20`
- `ReturnOver15`
- `ReturnOver10`
- `ReturnOver5`

Example: `ROAFilter{ReturnOver5}`

##### Return on Equity
Filter: `ROEFilter`

Values:

- `ReturnPositive`
- `ReturnNegative`
- `ReturnVeryPositive`
- `ReturnVeryNegative`
- `ReturnUnderNeg50`
- `ReturnUnderNeg45`
- `ReturnUnderNeg40`
- `ReturnUnderNeg35`
- `ReturnUnderNeg30`
- `ReturnUnderNeg25`
- `ReturnUnderNeg20`
- `ReturnUnderNeg15`
- `ReturnUnderNeg10`
- `ReturnUnderNeg5`
- `ReturnOver50`
- `ReturnOver45`
- `ReturnOver40`
- `ReturnOver35`
- `ReturnOver30`
- `ReturnOver25`
- `ReturnOver20`
- `ReturnOver15`
- `ReturnOver10`
- `ReturnOver5`

Example: `ROEFilter{ReturnUnderNeg5}`

##### Return on Investment
Filter: `ROIFilter`

Values:

- `ReturnPositive`
- `ReturnNegative`
- `ReturnVeryPositive`
- `ReturnVeryNegative`
- `ReturnUnderNeg50`
- `ReturnUnderNeg45`
- `ReturnUnderNeg40`
- `ReturnUnderNeg35`
- `ReturnUnderNeg30`
- `ReturnUnderNeg25`
- `ReturnUnderNeg20`
- `ReturnUnderNeg15`
- `ReturnUnderNeg10`
- `ReturnUnderNeg5`
- `ReturnOver50`
- `ReturnOver45`
- `ReturnOver40`
- `ReturnOver35`
- `ReturnOver30`
- `ReturnOver25`
- `ReturnOver20`
- `ReturnOver15`
- `ReturnOver10`
- `ReturnOver5`

Example: `ROIFilter{ReturnPositive}`

##### Current Ratio
Filter: `CurrentRatioFilter`

Values:

- `RatioHigh`
- `RatioLow`
- `RatioUnder1`
- `RatioUnder0point5`
- `RatioOver0point5`
- `RatioOver1`
- `RatioOver1point5`
- `RatioOver2`
- `RatioOver3`
- `RatioOver4`
- `RatioOver5`
- `RatioOver10`

Example: `CurrentRatioFilter{RatioHigh}`

##### Quick Ratio
Filter: `QuickRatioFilter`

Values:

- `RatioHigh`
- `RatioLow`
- `RatioUnder1`
- `RatioUnder0point5`
- `RatioOver0point5`
- `RatioOver1`
- `RatioOver1point5`
- `RatioOver2`
- `RatioOver3`
- `RatioOver4`
- `RatioOver5`
- `RatioOver10`

Example: `QuickRatioFilter{RatioUnder0point5}`

##### Long-Term Debt/Equity
Filter: `LTDebtEquityFilter`

Values:

- `DEHigh`
- `DELow`
- `DEUnder1`
- `DEUnder0point9`
- `DEUnder0point8`
- `DEUnder0point7`
- `DEUnder0point6`
- `DEUnder0point5`
- `DEUnder0point4`
- `DEUnder0point3`
- `DEUnder0point2`
- `DEUnder0point1`
- `DEOver1`
- `DEOver0point9`
- `DEOver0point8`
- `DEOver0point7`
- `DEOver0point6`
- `DEOver0point5`
- `DEOver0point4`
- `DEOver0point3`
- `DEOver0point2`
- `DEOver0point1`

Example: `LTDebtEquityFilter{DEOver1}`

##### Debt/Equity
Filter: `DebtEquityFilter`

Values:

- `DEHigh`
- `DELow`
- `DEUnder1`
- `DEUnder0point9`
- `DEUnder0point8`
- `DEUnder0point7`
- `DEUnder0point6`
- `DEUnder0point5`
- `DEUnder0point4`
- `DEUnder0point3`
- `DEUnder0point2`
- `DEUnder0point1`
- `DEOver1`
- `DEOver0point9`
- `DEOver0point8`
- `DEOver0point7`
- `DEOver0point6`
- `DEOver0point5`
- `DEOver0point4`
- `DEOver0point3`
- `DEOver0point2`
- `DEOver0point1`

Example: `DebtEquityFilter{DELow}`

##### Gross Margin
Filter: `GrossMarginFilter`

Values:

- `GMPositive`
- `GMNegative`
- `GMHigh`
- `GMUnder90`
- `GMUnder80`
- `GMUnder70`
- `GMUnder60`
- `GMUnder50`
- `GMUnder45`
- `GMUnder40`
- `GMUnder35`
- `GMUnder30`
- `GMUnder25`
- `GMUnder20`
- `GMUnder15`
- `GMUnder10`
- `GMUnder5`
- `GMUnder0`
- `GMUnderNeg10`
- `GMUnderNeg20`
- `GMUnderNeg30`
- `GMUnderNeg50`
- `GMUnderNeg70`
- `GMUnderNeg100`
- `GMOver0`
- `GMOver5`
- `GMOver10`
- `GMOver15`
- `GMOver20`
- `GMOver25`
- `GMOver30`
- `GMOver35`
- `GMOver40`
- `GMOver45`
- `GMOver50`
- `GMOver60`
- `GMOver70`
- `GMOver80`
- `GMOver90`

Example: `GrossMarginFilter{GMPositive}`

##### Operating Margin
Filter: `OperatingMarginFilter`

Values:

- `OMPositive`
- `OMNegative`
- `OMVeryNegative`
- `OMHigh`
- `OMUnder90`
- `OMUnder80`
- `OMUnder70`
- `OMUnder60`
- `OMUnder50`
- `OMUnder45`
- `OMUnder40`
- `OMUnder35`
- `OMUnder30`
- `OMUnder25`
- `OMUnder20`
- `OMUnder15`
- `OMUnder10`
- `OMUnder5`
- `OMUnder0`
- `OMUnderNeg10`
- `OMUnderNeg20`
- `OMUnderNeg30`
- `OMUnderNeg50`
- `OMUnderNeg70`
- `OMUnderNeg100`
- `OMOver0`
- `OMOver5`
- `OMOver10`
- `OMOver15`
- `OMOver20`
- `OMOver25`
- `OMOver30`
- `OMOver35`
- `OMOver40`
- `OMOver45`
- `OMOver50`
- `OMOver60`
- `OMOver70`
- `OMOver80`
- `OMOver90`

Example: `OperatingMarginFilter{OMUnder15}`

##### Net Profit Margin
Filter: `NetProfitMarginFilter`

Values:

- `OMPositive`
- `OMNegative`
- `OMVeryNegative`
- `OMHigh`
- `OMUnder90`
- `OMUnder80`
- `OMUnder70`
- `OMUnder60`
- `OMUnder50`
- `OMUnder45`
- `OMUnder40`
- `OMUnder35`
- `OMUnder30`
- `OMUnder25`
- `OMUnder20`
- `OMUnder15`
- `OMUnder10`
- `OMUnder5`
- `OMUnder0`
- `OMUnderNeg10`
- `OMUnderNeg20`
- `OMUnderNeg30`
- `OMUnderNeg50`
- `OMUnderNeg70`
- `OMUnderNeg100`
- `OMOver0`
- `OMOver5`
- `OMOver10`
- `OMOver15`
- `OMOver20`
- `OMOver25`
- `OMOver30`
- `OMOver35`
- `OMOver40`
- `OMOver45`
- `OMOver50`
- `OMOver60`
- `OMOver70`
- `OMOver80`
- `OMOver90`

Example: `NetProfitMarginFilter{OMOver80}`

##### Payout Ratio
Filter: `PayoutRatioFilter`

Values:

- `PRNone`
- `PRPositive`
- `PRLow`
- `PRHigh`
- `PROver1`
- `PROver2`
- `PROver3`
- `PROver4`
- `PROver5`
- `PROver6`
- `PROver7`
- `PROver8`
- `PROver9`
- `PROver10`
- `PRUnder1`
- `PRUnder2`
- `PRUnder3`
- `PRUnder4`
- `PRUnder5`
- `PRUnder6`
- `PRUnder7`
- `PRUnder8`
- `PRUnder9`
- `PRUnder10`

Example: `PayoutRatioFilter{PRNone}`

##### Insider Ownership
Filter: `InsiderOwnershipFilter`

Values:

- `InsdrOwnLow`
- `InsdrOwnHigh`
- `InsdrOwnVeryHigh`
- `InsdrOwnOver10`
- `InsdrOwnOver20`
- `InsdrOwnOver30`
- `InsdrOwnOver40`
- `InsdrOwnOver50`
- `InsdrOwnOver60`
- `InsdrOwnOver70`
- `InsdrOwnOver80`
- `InsdrOwnOver90`

Example: `InsiderOwnershipFilter{InsdrOwnLow}`

##### Insider Transactions
Filter: `InsiderTransactionsFilter`

Values:

- `InsdrTransVeryNeg`
- `InsdrTransNeg`
- `InsdrTransPos`
- `InsdrTransVeryPos`
- `InsdrTransOver5`
- `InsdrTransOver10`
- `InsdrTransOver15`
- `InsdrTransOver20`
- `InsdrTransOver25`
- `InsdrTransOver30`
- `InsdrTransOver35`
- `InsdrTransOver40`
- `InsdrTransOver45`
- `InsdrTransOver50`
- `InsdrTransOver60`
- `InsdrTransOver70`
- `InsdrTransOver80`
- `InsdrTransOver90`
- `InsdrTransUnderNeg5`
- `InsdrTransUnderNeg10`
- `InsdrTransUnderNeg15`
- `InsdrTransUnderNeg20`
- `InsdrTransUnderNeg25`
- `InsdrTransUnderNeg30`
- `InsdrTransUnderNeg35`
- `InsdrTransUnderNeg40`
- `InsdrTransUnderNeg45`
- `InsdrTransUnderNeg50`
- `InsdrTransUnderNeg60`
- `InsdrTransUnderNeg70`
- `InsdrTransUnderNeg80`
- `InsdrTransUnderNeg90`

Example: `InsiderTransactionsFilter{InsdrTransUnderNeg5}`

##### Institutional Ownership
Filter: `InstitutionalOwnershipFilter`

Values:

- `InstOwnLow`
- `InstOwnHigh`
- `InstOwnOver10`
- `InstOwnOver20`
- `InstOwnOver30`
- `InstOwnOver40`
- `InstOwnOver50`
- `InstOwnOver60`
- `InstOwnOver70`
- `InstOwnOver80`
- `InstOwnOver90`
- `InstOwnUnder10`
- `InstOwnUnder20`
- `InstOwnUnder30`
- `InstOwnUnder40`
- `InstOwnUnder50`
- `InstOwnUnder60`
- `InstOwnUnder70`
- `InstOwnUnder80`
- `InstOwnUnder90`

Example: `InstitutionalOwnershipFilter{InstOwnUnder50}`

##### Institutional Transactions
Filter: `InstitutionalTransactionsFilter`

Values:

- `InstTransVeryNeg`
- `InstTransNeg`
- `InstTransPos`
- `InstTransVeryPos`
- `InstTransOver5`
- `InstTransOver10`
- `InstTransOver15`
- `InstTransOver20`
- `InstTransOver25`
- `InstTransOver30`
- `InstTransOver35`
- `InstTransOver40`
- `InstTransOver45`
- `InstTransOver50`
- `InstTransUnderNeg5`
- `InstTransUnderNeg10`
- `InstTransUnderNeg15`
- `InstTransUnderNeg20`
- `InstTransUnderNeg25`
- `InstTransUnderNeg30`
- `InstTransUnderNeg35`
- `InstTransUnderNeg40`
- `InstTransUnderNeg45`
- `InstTransUnderNeg50`

Example: `InstitutionalTransactionsFilter{InstTransPos}`

##### Performance
Filter: `PerformanceFilter`

Values:

- `PerfTodayUp`
- `PerfTodayDown`
- `PerfTodayDown15Percent`
- `PerfTodayDown10Percent`
- `PerfTodayDown5Percent`
- `PerfTodayUp5Percent`
- `PerfTodayUp10Percent`
- `PerfTodayUp15Percent`
- `PerfWeekDown30Percent`
- `PerfWeekDown20Percent`
- `PerfWeekDown10Percent`
- `PerfWeekDown`
- `PerfWeekUp`
- `PerfWeekUp10Percent`
- `PerfWeekUp20Percent`
- `PerfWeekUp30Percent`
- `PerfMonthDown50Percent`
- `PerfMonthDown30Percent`
- `PerfMonthDown20Percent`
- `PerfMonthDown10Percent`
- `PerfMonthDown`
- `PerfMonthUp`
- `PerfMonthUp10Percent`
- `PerfMonthUp20Percent`
- `PerfMonthUp30Percent`
- `PerfMonthUp50Percent`
- `PerfQuarterDown50Percent`
- `PerfQuarterDown30Percent`
- `PerfQuarterDown20Percent`
- `PerfQuarterDown10Percent`
- `PerfQuarterDown`
- `PerfQuarterUp`
- `PerfQuarterUp10Percent`
- `PerfQuarterUp20Percent`
- `PerfQuarterUp30Percent`
- `PerfQuarterUp50Percent`
- `PerfSixMonthsDown75Percent`
- `PerfSixMonthsDown50Percent`
- `PerfSixMonthsDown30Percent`
- `PerfSixMonthsDown20Percent`
- `PerfSixMonthsDown10Percent`
- `PerfSixMonthsDown`
- `PerfSixMonthsUp`
- `PerfSixMonthsUp10Percent`
- `PerfSixMonthsUp20Percent`
- `PerfSixMonthsUp30Percent`
- `PerfSixMonthsUp50Percent`
- `PerfSixMonthsUp100Percent`
- `PerfYearDown75Percent`
- `PerfYearDown50Percent`
- `PerfYearDown30Percent`
- `PerfYearDown20Percent`
- `PerfYearDown10Percent`
- `PerfYearDown`
- `PerfYearUp`
- `PerfYearUp10Percent`
- `PerfYearUp20Percent`
- `PerfYearUp30Percent`
- `PerfYearUp50Percent`
- `PerfYearUp100Percent`
- `PerfYearUp200Percent`
- `PerfYearUp300Percent`
- `PerfYearUp500Percent`
- `PerfYTDDown75Percent`
- `PerfYTDDown50Percent`
- `PerfYTDDown30Percent`
- `PerfYTDDown20Percent`
- `PerfYTDDown10Percent`
- `PerfYTDDown5Percent`
- `PerfYTDDown`
- `PerfYTDUp`
- `PerfYTDUp5Percent`
- `PerfYTDUp10Percent`
- `PerfYTDUp20Percent`
- `PerfYTDUp30Percent`
- `PerfYTDUp50Percent`
- `PerfYTDUp100Percent`

Example: `PerformanceFilter{PerfYTDDown}`

##### Performance 2
Filter: `Performance2Filter`

Values:

- `PerfTodayUp`
- `PerfTodayDown`
- `PerfTodayDown15Percent`
- `PerfTodayDown10Percent`
- `PerfTodayDown5Percent`
- `PerfTodayUp5Percent`
- `PerfTodayUp10Percent`
- `PerfTodayUp15Percent`
- `PerfWeekDown30Percent`
- `PerfWeekDown20Percent`
- `PerfWeekDown10Percent`
- `PerfWeekDown`
- `PerfWeekUp`
- `PerfWeekUp10Percent`
- `PerfWeekUp20Percent`
- `PerfWeekUp30Percent`
- `PerfMonthDown50Percent`
- `PerfMonthDown30Percent`
- `PerfMonthDown20Percent`
- `PerfMonthDown10Percent`
- `PerfMonthDown`
- `PerfMonthUp`
- `PerfMonthUp10Percent`
- `PerfMonthUp20Percent`
- `PerfMonthUp30Percent`
- `PerfMonthUp50Percent`
- `PerfQuarterDown50Percent`
- `PerfQuarterDown30Percent`
- `PerfQuarterDown20Percent`
- `PerfQuarterDown10Percent`
- `PerfQuarterDown`
- `PerfQuarterUp`
- `PerfQuarterUp10Percent`
- `PerfQuarterUp20Percent`
- `PerfQuarterUp30Percent`
- `PerfQuarterUp50Percent`
- `PerfSixMonthsDown75Percent`
- `PerfSixMonthsDown50Percent`
- `PerfSixMonthsDown30Percent`
- `PerfSixMonthsDown20Percent`
- `PerfSixMonthsDown10Percent`
- `PerfSixMonthsDown`
- `PerfSixMonthsUp`
- `PerfSixMonthsUp10Percent`
- `PerfSixMonthsUp20Percent`
- `PerfSixMonthsUp30Percent`
- `PerfSixMonthsUp50Percent`
- `PerfSixMonthsUp100Percent`
- `PerfYearDown75Percent`
- `PerfYearDown50Percent`
- `PerfYearDown30Percent`
- `PerfYearDown20Percent`
- `PerfYearDown10Percent`
- `PerfYearDown`
- `PerfYearUp`
- `PerfYearUp10Percent`
- `PerfYearUp20Percent`
- `PerfYearUp30Percent`
- `PerfYearUp50Percent`
- `PerfYearUp100Percent`
- `PerfYearUp200Percent`
- `PerfYearUp300Percent`
- `PerfYearUp500Percent`
- `PerfYTDDown75Percent`
- `PerfYTDDown50Percent`
- `PerfYTDDown30Percent`
- `PerfYTDDown20Percent`
- `PerfYTDDown10Percent`
- `PerfYTDDown5Percent`
- `PerfYTDDown`
- `PerfYTDUp`
- `PerfYTDUp5Percent`
- `PerfYTDUp10Percent`
- `PerfYTDUp20Percent`
- `PerfYTDUp30Percent`
- `PerfYTDUp50Percent`
- `PerfYTDUp100Percent`

Example: `Performance2Filter{PerfWeekUp}`

##### Volatility
Filter: `VolatilityFilter`

Values:

- `VolWeekOver3`
- `VolWeekOver4`
- `VolWeekOver5`
- `VolWeekOver6`
- `VolWeekOver7`
- `VolWeekOver8`
- `VolWeekOver9`
- `VolWeekOver10`
- `VolWeekOver12`
- `VolWeekOver15`
- `VolMonthOver2`
- `VolMonthOver3`
- `VolMonthOver4`
- `VolMonthOver5`
- `VolMonthOver6`
- `VolMonthOver7`
- `VolMonthOver8`
- `VolMonthOver9`
- `VolMonthOver10`
- `VolMonthOver12`
- `VolMonthOver15`

Example: `VolatilityFilter{VolWeekOver7}`

##### RSI (14)
Filter: `RSIFilter`

Values:

- `Overbought90`
- `Overbought80`
- `Overbought70`
- `Overbought60`
- `Oversold40`
- `Oversold30`
- `Oversold20`
- `Oversold10`
- `NotOverboughtUnder60`
- `NotOverboughtUnder50`
- `NotOversoldOver50`
- `NotOversoldOver40`

Example: `RSIFilter{Oversold10}`

##### Gap
Filter: `GapFilter`

Values:

- `GapUp`
- `GapUp0Percent`
- `GapUp1Percent`
- `GapUp2Percent`
- `GapUp3Percent`
- `GapUp4Percent`
- `GapUp5Percent`
- `GapUp6Percent`
- `GapUp7Percent`
- `GapUp8Percent`
- `GapUp9Percent`
- `GapUp10Percent`
- `GapUp15Percent`
- `GapUp20Percent`
- `GapDown`
- `GapDown0Percent`
- `GapDown1Percent`
- `GapDown2Percent`
- `GapDown3Percent`
- `GapDown4Percent`
- `GapDown5Percent`
- `GapDown6Percent`
- `GapDown7Percent`
- `GapDown8Percent`
- `GapDown9Percent`
- `GapDown10Percent`
- `GapDown15Percent`
- `GapDown20Percent`

Example: `GapFilter{GapDown}`

##### 20-Day Simple Moving Average
Filter: `SMA20Filter`

Values:

- `Mavg20PriceBelowSMA20`
- `Mavg20Price10PercentBelowSMA20`
- `Mavg20Price20PercentBelowSMA20`
- `Mavg20Price30PercentBelowSMA20`
- `Mavg20Price40PercentBelowSMA20`
- `Mavg20Price50PercentBelowSMA20`
- `Mavg20PriceAboveSMA20`
- `Mavg20Price10PercentAboveSMA20`
- `Mavg20Price20PercentAboveSMA20`
- `Mavg20Price30PercentAboveSMA20`
- `Mavg20Price40PercentAboveSMA20`
- `Mavg20Price50PercentAboveSMA20`
- `Mavg20PriceCrossedSMA20`
- `Mavg20PriceCrossedSMA20Above`
- `Mavg20PriceCrossedSMA20Below`
- `Mavg20SMA20CrossedSMA50`
- `Mavg20SMA20CrossedSMA50Above`
- `Mavg20SMA20CrossedSMA50Below`
- `Mavg20SMA20CrossedSMA200`
- `Mavg20SMA20CrossedSMA200Above`
- `Mavg20SMA20CrossedSMA200Below`
- `Mavg20SMA20AboveSMA50`
- `Mavg20SMA20BelowSMA50`
- `Mavg20SMA20AboveSMA200`
- `Mavg20SMA20BelowSMA200`

Example: `SMA20Filter{Mavg20PriceBelowSMA20}`

##### 50-Day Simple Moving Average
Filter: `SMA50Filter`

Values:

- `Mavg50PriceBelowSMA50`
- `Mavg50Price10PercentBelowSMA50`
- `Mavg50Price20PercentBelowSMA50`
- `Mavg50Price30PercentBelowSMA50`
- `Mavg50Price40PercentBelowSMA50`
- `Mavg50Price50PercentBelowSMA50`
- `Mavg50PriceAboveSMA50`
- `Mavg50Price10PercentAboveSMA50`
- `Mavg50Price20PercentAboveSMA50`
- `Mavg50Price30PercentAboveSMA50`
- `Mavg50Price40PercentAboveSMA50`
- `Mavg50Price50PercentAboveSMA50`
- `Mavg50PriceCrossedSMA50`
- `Mavg50PriceCrossedSMA50Above`
- `Mavg50PriceCrossedSMA50Below`
- `Mavg50SMA50CrossedSMA20`
- `Mavg50SMA50CrossedSMA20Above`
- `Mavg50SMA50CrossedSMA20Below`
- `Mavg50SMA50CrossedSMA200`
- `Mavg50SMA50CrossedSMA200Above`
- `Mavg50SMA50CrossedSMA200Below`
- `Mavg50SMA50AboveSMA20`
- `Mavg50SMA50BelowSMA20`
- `Mavg50SMA50AboveSMA200`
- `Mavg50SMA50BelowSMA200`

Example: `SMA50Filter{Mavg50SMA50AboveSMA20}`

##### 200-Day Simple Moving Average
Filter: `SMA200Filter`

Values:

- `Mavg200PriceBelowSMA200`
- `Mavg200Price10PercentBelowSMA200`
- `Mavg200Price20PercentBelowSMA200`
- `Mavg200Price30PercentBelowSMA200`
- `Mavg200Price40PercentBelowSMA200`
- `Mavg200Price50PercentBelowSMA200`
- `Mavg200Price60PercentBelowSMA200`
- `Mavg200Price70PercentBelowSMA200`
- `Mavg200Price80PercentBelowSMA200`
- `Mavg200Price90PercentBelowSMA200`
- `Mavg200PriceAboveSMA200`
- `Mavg200Price10PercentAboveSMA200`
- `Mavg200Price20PercentAboveSMA200`
- `Mavg200Price30PercentAboveSMA200`
- `Mavg200Price40PercentAboveSMA200`
- `Mavg200Price50PercentAboveSMA200`
- `Mavg200Price60PercentAboveSMA200`
- `Mavg200Price70PercentAboveSMA200`
- `Mavg200Price80PercentAboveSMA200`
- `Mavg200Price90PercentAboveSMA200`
- `Mavg200Price100PercentAboveSMA200`
- `Mavg200PriceCrossedSMA200`
- `Mavg200PriceCrossedSMA200Above`
- `Mavg200PriceCrossedSMA200Below`
- `Mavg200SMA200CrossedSMA20`
- `Mavg200SMA200CrossedSMA20Above`
- `Mavg200SMA200CrossedSMA20Below`
- `Mavg200SMA200CrossedSMA50`
- `Mavg200SMA200CrossedSMA50Above`
- `Mavg200SMA200CrossedSMA50Below`
- `Mavg200SMA200AboveSMA20`
- `Mavg200SMA200BelowSMA20`
- `Mavg200SMA200AboveSMA50`
- `Mavg200SMA200BelowSMA50`

Example: `SMA200Filter{Mavg200PriceCrossedSMA200}`

##### Change
Filter: `ChangeFilter`

Values:

- `ChangeUp`
- `ChangeUp1Percent`
- `ChangeUp2Percent`
- `ChangeUp3Percent`
- `ChangeUp4Percent`
- `ChangeUp5Percent`
- `ChangeUp6Percent`
- `ChangeUp7Percent`
- `ChangeUp8Percent`
- `ChangeUp9Percent`
- `ChangeUp10Percent`
- `ChangeUp15Percent`
- `ChangeUp20Percent`
- `ChangeDown`
- `ChangeDown1Percent`
- `ChangeDown2Percent`
- `ChangeDown3Percent`
- `ChangeDown4Percent`
- `ChangeDown5Percent`
- `ChangeDown6Percent`
- `ChangeDown7Percent`
- `ChangeDown8Percent`
- `ChangeDown9Percent`
- `ChangeDown10Percent`
- `ChangeDown15Percent`
- `ChangeDown20Percent`

Example: `ChangeFilter{ChangeUp}`

##### Change From Open
Filter: `ChangeFromOpenFilter`

Values:

- `ChangeUp`
- `ChangeUp1Percent`
- `ChangeUp2Percent`
- `ChangeUp3Percent`
- `ChangeUp4Percent`
- `ChangeUp5Percent`
- `ChangeUp6Percent`
- `ChangeUp7Percent`
- `ChangeUp8Percent`
- `ChangeUp9Percent`
- `ChangeUp10Percent`
- `ChangeUp15Percent`
- `ChangeUp20Percent`
- `ChangeDown`
- `ChangeDown1Percent`
- `ChangeDown2Percent`
- `ChangeDown3Percent`
- `ChangeDown4Percent`
- `ChangeDown5Percent`
- `ChangeDown6Percent`
- `ChangeDown7Percent`
- `ChangeDown8Percent`
- `ChangeDown9Percent`
- `ChangeDown10Percent`
- `ChangeDown15Percent`
- `ChangeDown20Percent`

Example: `ChangeFromOpenFilter{ChangeDown10Percent}`

##### 20-Day High/Low
Filter: `HighLow20DayFilter`

Values:

- `DHLNewHigh`
- `DHLNewLow`
- `DHLAtLeast5PercentBelowHigh`
- `DHLAtLeast10PercentBelowHigh`
- `DHLAtLeast15PercentBelowHigh`
- `DHLAtLeast20PercentBelowHigh`
- `DHLAtLeast30PercentBelowHigh`
- `DHLAtLeast40PercentBelowHigh`
- `DHLAtLeast50PercentBelowHigh`
- `DHL0to3PercentBelowHigh`
- `DHL0to5PercentBelowHigh`
- `DHL0to10PercentBelowHigh`
- `DHLAtLeast5PercentAboveLow`
- `DHLAtLeast10PercentAboveLow`
- `DHLAtLeast15PercentAboveLow`
- `DHLAtLeast20PercentAboveLow`
- `DHLAtLeast30PercentAboveLow`
- `DHLAtLeast40PercentAboveLow`
- `DHLAtLeast50PercentAboveLow`
- `DHL0to3PercentAboveLow`
- `DHL0to5PercentAboveLow`
- `DHL0to10PercentAboveLow`

Example: `HighLow20DayFilter{DHL0to3PercentBelowHigh}`

##### 50-Day High/Low
Filter: `HighLow50DayFilter`

Values:

- `DHLNewHigh`
- `DHLNewLow`
- `DHLAtLeast5PercentBelowHigh`
- `DHLAtLeast10PercentBelowHigh`
- `DHLAtLeast15PercentBelowHigh`
- `DHLAtLeast20PercentBelowHigh`
- `DHLAtLeast30PercentBelowHigh`
- `DHLAtLeast40PercentBelowHigh`
- `DHLAtLeast50PercentBelowHigh`
- `DHL0to3PercentBelowHigh`
- `DHL0to5PercentBelowHigh`
- `DHL0to10PercentBelowHigh`
- `DHLAtLeast5PercentAboveLow`
- `DHLAtLeast10PercentAboveLow`
- `DHLAtLeast15PercentAboveLow`
- `DHLAtLeast20PercentAboveLow`
- `DHLAtLeast30PercentAboveLow`
- `DHLAtLeast40PercentAboveLow`
- `DHLAtLeast50PercentAboveLow`
- `DHL0to3PercentAboveLow`
- `DHL0to5PercentAboveLow`
- `DHL0to10PercentAboveLow`

Example: `HighLow50DayFilter{DHLNewHigh}`

##### 52-Week High/Low
Filter: `HighLow52WeekFilter`

Values:

- `WHLNewHigh`
- `WHLNewLow`
- `WHLAtLeast5PercentBelowHigh`
- `WHLAtLeast10PercentBelowHigh`
- `WHLAtLeast15PercentBelowHigh`
- `WHLAtLeast20PercentBelowHigh`
- `WHLAtLeast30PercentBelowHigh`
- `WHLAtLeast40PercentBelowHigh`
- `WHLAtLeast50PercentBelowHigh`
- `WHLAtLeast60PercentBelowHigh`
- `WHLAtLeast70PercentBelowHigh`
- `WHLAtLeast80PercentBelowHigh`
- `WHLAtLeast90PercentBelowHigh`
- `WHL0to3PercentBelowHigh`
- `WHL0to5PercentBelowHigh`
- `WHL0to10PercentBelowHigh`
- `WHLAtLeast5PercentAboveLow`
- `WHLAtLeast10PercentAboveLow`
- `WHLAtLeast15PercentAboveLow`
- `WHLAtLeast20PercentAboveLow`
- `WHLAtLeast30PercentAboveLow`
- `WHLAtLeast40PercentAboveLow`
- `WHLAtLeast50PercentAboveLow`
- `WHLAtLeast60PercentAboveLow`
- `WHLAtLeast70PercentAboveLow`
- `WHLAtLeast80PercentAboveLow`
- `WHLAtLeast90PercentAboveLow`
- `WHLAtLeast100PercentAboveLow`
- `WHLAtLeast120PercentAboveLow`
- `WHLAtLeast150PercentAboveLow`
- `WHLAtLeast200PercentAboveLow`
- `WHLAtLeast300PercentAboveLow`
- `WHLAtLeast500PercentAboveLow`
- `WHL0to3PercentAboveLow`
- `WHL0to5PercentAboveLow`
- `WHL0to10PercentAboveLow`

Example: `HighLow52WeekFilter{WHLNewLow}`

##### Pattern
Filter: `PatternFilter`

Values:

- `PatternHorizontalSR`
- `PatternHorizontalSRStrong`
- `PatternTLResistance`
- `PatternTLResistanceStrong`
- `PatternTLSupport`
- `PatternTLSupportStrong`
- `PatternWedgeUp`
- `PatternWedgeUpStrong`
- `PatternWedgeDown`
- `PatternWedgeDownStrong`
- `PatternTriangleAscending`
- `PatternTriangleAscendingStrong`
- `PatternTriangleDescending`
- `PatternTriangleDescendingStrong`
- `PatternWedge`
- `PatternWedgeStrong`
- `PatternChannelUp`
- `PatternChannelUpStrong`
- `PatternChannelDown`
- `PatternChannelDownStrong`
- `PatternChannel`
- `PatternChannelStrong`
- `PatternDoubleTop`
- `PatternDoubleBottom`
- `PatternMultipleTop`
- `PatternMultipleBottom`
- `PatternHeadAndShoulders`
- `PatternHeadAndShouldersInverse`

Example: `PatternFilter{PatternChannelDown}`

##### Candlestick
Filter: `CandlestickFilter`

Values:

- `CSLongLowerShadow`
- `CSLongUpperShadow`
- `CSHammer`
- `CSInvertedHammer`
- `CSSpinningTopWhite`
- `CSSpinningTopBlack`
- `CSDoji`
- `CSDragonflyDoji`
- `CSGravestoneDoji`
- `CSMarubozuWhite`
- `CSMarubozuBlack`

Example: `CandlestickFilter{CSMarubozuBlack}`

##### Beta
Filter: `BetaFilter`

Values:

- `BetaUnder0`
- `BetaUnder0Point5`
- `BetaUnder1`
- `BetaUnder1Point5`
- `BetaUnder2`
- `BetaOver0`
- `BetaOver0Point5`
- `BetaOver1`
- `BetaOver1Point5`
- `BetaOver2`
- `BetaOver2Point5`
- `BetaOver3`
- `BetaOver4`
- `Beta0to0Point5`
- `Beta0to1`
- `Beta0Point5to1`
- `Beta0Point5to1Point5`
- `Beta1to1Point5`
- `Beta1to2`

Example: `BetaFilter{BetaUnder0Point5}`

##### Average True Range
Filter: `AverageTrueRangeFilter`

Values:

- `ATROver0Point25`
- `ATROver0Point5`
- `ATROver0Point75`
- `ATROver1`
- `ATROver1Point5`
- `ATROver2`
- `ATROver2Point5`
- `ATROver3`
- `ATROver3Point5`
- `ATROver4`
- `ATROver4Point5`
- `ATROver5`
- `ATRUnder0Point25`
- `ATRUnder0Point5`
- `ATRUnder0Point75`
- `ATRUnder1`
- `ATRUnder1Point5`
- `ATRUnder2`
- `ATRUnder2Point5`
- `ATRUnder3`
- `ATRUnder3Point5`
- `ATRUnder4`
- `ATRUnder4Point5`
- `ATRUnder5`

Example: `AverageTrueRangeFilter{ATROver2}`


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
- [ ] Add support for non-default views (increase the number of tickers a screen can return)
- [ ] Review dataframe package and consider migration for greater support

## Contributing

You can contribute to this project by reporting bugs, suggesting enhancements, or directly by extending and writing features.

## Disclaimer

Using this library to acquire data from Finviz is against their Terms of Service and `robots.txt`.
Use it responsively and at your own risk. This library was built purely for educational purposes.
