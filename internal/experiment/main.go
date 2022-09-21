package main

import (
	"fmt"

	"github.com/RageCage64/multilinediff"
)

func main() {
	a := `a
b
c
diff`
	b := `a b
c
f
f
f
diff`
	diff, count := multilinediff.Diff(a, b, "\n")
	fmt.Println(count)
	fmt.Println("--------")
	fmt.Println(diff)
}
