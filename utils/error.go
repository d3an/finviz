// Copyright (c) 2022 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package utils

import (
	"fmt"
	"os"
)

// Err is used for error exits for the cmd util
func Err(msg interface{}) {
	fmt.Printf("[ERROR] %v\n", msg)
	os.Exit(1)
}

// ErrorSignalNotFound is the error thrown if a query string is not found in the SignalLookup dict
type ErrorSignalNotFound string

func (err ErrorSignalNotFound) Error() string {
	return string(err)
}

// SpecificOrderNotFoundError is the error thrown if a query string is not found in the SpecificOrderLookup dict
type SpecificOrderNotFoundError string

func (err SpecificOrderNotFoundError) Error() string {
	return string(err)
}

// MultipleValuesError is the error thrown if an unsupported filter is initialized with multiple values
type MultipleValuesError string

func (err MultipleValuesError) Error() string {
	return string(err)
}

// NoValuesError is the error thrown if a filter is initialized with no values
type NoValuesError string

func (err NoValuesError) Error() string {
	return string(err)
}

// DuplicateFilterError is the error thrown if the same filter is declared more than once
type DuplicateFilterError string

func (err DuplicateFilterError) Error() string {
	return string(err)
}

// FilterNotFoundError is the error thrown if a query string is not found in the FilterLookup dict
type FilterNotFoundError string

func (err FilterNotFoundError) Error() string {
	return string(err)
}

// IncompatibleChartTypeTimeFrameError is the error thrown if a newTimeFrame is not one of valid for the specified chart type
type IncompatibleChartTypeTimeFrameError string

func (err IncompatibleChartTypeTimeFrameError) Error() string {
	return string(err)
}

// InvalidChartTypeError is the error thrown if a newChartType is not one of "technical", "line", or "candle"
type InvalidChartTypeError string

func (err InvalidChartTypeError) Error() string {
	return string(err)
}

// InvalidTimeFrameError is the error thrown if a newTimeFrame is not one of valid for the specified chart type
type InvalidTimeFrameError string

func (err InvalidTimeFrameError) Error() string {
	return string(err)
}

// InvalidViewError is the error thrown if a query string is not associated with a labelled ViewType
type InvalidViewError string

func (err InvalidViewError) Error() string {
	return string(err)
}

// StatusCodeError is the error given if a request's status code is not 200
type StatusCodeError string

func (err StatusCodeError) Error() string {
	return string(err)
}

// NoStocksMatchedQueryError is the error given if a screen returns no results
type NoStocksMatchedQueryError string

func (err NoStocksMatchedQueryError) Error() string {
	return string(err)
}

// MethodNotImplementedError is the error given if a method has not been written yet
type MethodNotImplementedError string

func (err MethodNotImplementedError) Error() string {
	return string(err)
}
