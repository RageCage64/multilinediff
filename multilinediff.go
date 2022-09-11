package multilinediff

import (
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func MultilineDiff(a, b, lineSep string) string {
	reporter := Reporter{
		LineSep: lineSep,
	}
	cmp.Diff(
		a, b,
		cmpopts.AcyclicTransformer("multiline", func(s string) []string {
			return strings.Split(s, lineSep)
		}),
		cmp.Reporter(&reporter),
	)
	return reporter.String()
}
