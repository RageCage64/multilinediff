package prettyreporter

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

type Reporter struct {
	path      cmp.Path
	old       []string
	new       []string
	diffCount int
}

func (r *Reporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

func (r *Reporter) Report(rs cmp.Result) {
	vOld, vNew := r.path.Last().Values()
	if !rs.Equal() {
		r.diffCount++
		if vOld.IsValid() {
			r.old = append(r.old, fmt.Sprintf("%+v", vOld))
		}
		if vNew.IsValid() {
			r.new = append(r.new, fmt.Sprintf("%+v", vNew))
		}
	} else {
		r.old = append(r.old, "")
		r.new = append(r.new, fmt.Sprintf("%+v", vOld))
	}
}

func (r *Reporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

func (r *Reporter) String() string {

	return strings.Join(r.lines, "\n")
}

func maxLen(strs []string) int {
	max := 0
	for _, s := range strs {
		if len(s) > max {
			max = len(s)
		}
	}
	return max
}
