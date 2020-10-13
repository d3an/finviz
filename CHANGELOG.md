# Changelog
All notable changes to this project will be documented in this file.

This format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## Types of changes

 * `Added` for new features.
 * `Improvements` for improvements to existing features.
 * `Changed` for changes in existing functionality, or for refactors.
 * `Deprecated` for soon-to-be removed features.
 * `Removed` for now removed features.
 * `Fixed` for any bug fixes.
 * `Security` in case of vulnerabilities.
 * `Reverts` for any reverted features.
 * `[v.x.y.z][YYYY.MM.DD]` for version changes.


## [Unreleased]

## [v.1.0.4][2020.10.13]
### Added
- `SD-31` Added `CleanScreenerDataFrame` to `GetScreenerData`. Added unit test.

## [v.1.0.3][2020.10.13]
### Added
- `SD-50` Added coverage badge
- `SD-45` Added support for News page and CLI subcommand for news

### Changed
- `SD-54` Restructured views and CLI to support multi-app structure
- `SD-53` Fixed README.md typo

### Improvements
- `SD-55` Updated Go version from `1.14.6` to `1.15.2`

## [v.1.0.2][2020.10.08]
### Added
- `SD-17` Added export functions for csv and json
- added lookup tables for most types
- `SD-23` added support for most free views (110, 120, 130, 140, 150, 160, 170, 310, 320, 330, 340, 350, 410, 510, 520)
- `SD-24` added support for charts view (210)
- `SD-38` added support for custom chart timeframes
- `SD-39` added support for custom chart types
- `SD-22` added unit testing for all `Scrape` functions with `go-vcr`
- added unit testing for `ChartViewInterface.SetTimeFrame`
- added unit testing for `ChartViewInterface.SetChartType`
- `SD-27` added `PrintDataFrame` function to print full DataFrames to console
- `SD-14` added beta Finviz CLI
- `SD-28` added properties to filters

### Changed
- `SD-15` Refactored `filters.go`
- refactored `screener_view.go`
- `SD-37` Created `.golangci.yaml` and added `gocognit` support

### Fixed
- `SD-25` Fixed ordering for bulk views (500 series) with breadth-first iteration
- `SD-26` Fixed double question mark in URL generation
- `SD-16` Added random user-agent to requests to bypass 403 HTTP errors

### Removed
- `SD-37` Removed `gocyclo` from `.pre-commit-config.yaml`

## [v.1.0.1][2020.08.06]
### Added
- `SD-5` Added `SetMultipleValues` to `FilterInterface` to allow multiple of the same filters, with `|` operator.

## [v.1.0.0][2020.07.27]
### Added
- `SD-1` Added exhaustive types for filters, sorting, views, and signals. Also, added basic screening and scraping functions.
- LICENSE.

### Fixed
- `SD-2` Updated module name to common package syntax.
- `SD-3` Exported struct fields to fix implicit assignment errors.

### Improvements
- `SD-1` Updated the README with more documentation coverage.
- `SD-4` Corrected the README code sample.
