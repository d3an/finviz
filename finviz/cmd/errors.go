// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

// MalformedFilterError is the error thrown if a filter argument is not well-formed
// Proper form: `Filter:Value1,Value2,Value3` OR `Filter:SingleValue`
type MalformedFilterError string

func (err MalformedFilterError) Error() string {
	return string(err)
}
