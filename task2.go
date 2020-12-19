package main

import (
	"fmt"
	"math"
)

func main() {

	x := 4.1
	y := 2.2
	z := 1.3
	a := 2.0
	c := ((math.Sqrt(x) + math.Sqrt(y)) * z) / a
	fmt.Printf("((Sqrt(%.2f) + Sqrt(%.2f)) * %.2f) / %.2f = %.2f \n", x, y, z, a, c)

}
