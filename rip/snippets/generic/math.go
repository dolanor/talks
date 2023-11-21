package generic

import (
	"fmt"
	"math"
)

func min() {
	// start min OMIT
	x := 1
	y := 2
	z := math.Min(float64(x), float64(y)) // HLUgly
	fmt.Println(z)
	// end min OMIT
}
