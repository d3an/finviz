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
### Added
- `SD-5` Added `SetMultipleValues` to FilterInterface to allow multiple of the same filters, with | operator.

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
