package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-gota/gota/dataframe"
)

type Enum struct {
	Allowed []string
	Value   string
}

// NewEnum give a list of allowed flag parameters, where the second argument is the default
func NewEnum(allowed []string, d string) *Enum {
	return &Enum{
		Allowed: allowed,
		Value:   d,
	}
}

func (a Enum) String() string {
	return a.Value
}

func (a *Enum) Set(p string) error {
	isIncluded := func(opts []string, val string) bool {
		for _, opt := range opts {
			if val == opt {
				return true
			}
		}
		return false
	}
	if !isIncluded(a.Allowed, p) {
		return fmt.Errorf("%s is not included in %s", p, strings.Join(a.Allowed, ","))
	}
	a.Value = p
	return nil
}

func (a *Enum) Type() string {
	return "string"
}

func ExportData(df *dataframe.DataFrame, outFile string) error {
	if outFile == "" {
		PrintFullDataFrame(df)
		return nil
	}

	switch filepath.Ext(outFile) {
	default:
		fallthrough
	case ".csv":
		return ExportCSV(df, outFile)
	case ".json":
		return ExportJSON(df, outFile)
	}
}
