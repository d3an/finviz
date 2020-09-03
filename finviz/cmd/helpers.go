// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"fmt"
	"os"
	"strings"
)

func er(msg interface{}) {
	fmt.Println("Error: ", msg)
	os.Exit(1)
}

func extractFilterInput(filterArg string) (filterName string, filterValues []string, err error) {
	firstSplit := strings.Split(filterArg, ":")
	if splitCount := len(firstSplit); splitCount != 2 {
		return "", nil, MalformedFilterError(fmt.Sprintf("Filter argument: \"%v\" is not well-formed. Proper form: \"Filter:Value1,Value2\" OR \"filter:value\"", filterArg))
	}
	return firstSplit[0], strings.Split(firstSplit[1], ","), nil
}
