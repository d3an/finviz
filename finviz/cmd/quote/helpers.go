// Copyright (c) 2020 James Bury. All rights reserved.
// Project site: https://github.com/d3an/finviz
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package quote

import (
	"fmt"
	"os"
)

func er(msg interface{}) {
	fmt.Println("Error: ", msg)
	os.Exit(1)
}
