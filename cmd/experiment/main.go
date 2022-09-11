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
	b := `a
b
c
f
f
f`

	fmt.Println(multilinediff.MultilineDiff(a, b, "\n"))
}
