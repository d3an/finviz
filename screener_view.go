// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package finviz

// ViewType represents the general view in which results are displayed
type ViewType string

// ColumnType represents the various column sets available for screens
type ColumnType string

// TabType represents which tab (filters, settings, or none) is visible on screen
type TabType string

// View Types
const (
	OverviewView    ViewType = "1"
	ChartsView      ViewType = "2"
	BasicView       ViewType = "3"
	TickersView     ViewType = "4"
	BulkTickersView ViewType = "5"
)

// Column Types
const (
	OverviewColumns    ColumnType = "1"
	ValuationColumns   ColumnType = "2"
	OwnershipColumns   ColumnType = "3"
	PerformanceColumns ColumnType = "4"
	CustomColumns      ColumnType = "5"
	FinancialColumns   ColumnType = "6"
	TechnicalColumns   ColumnType = "7"
)

// Extra Tabs
const (
	NoTab       TabType = "0"
	FilterTab   TabType = "1"
	SettingsTab TabType = "2"
)
