package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var prev float64
	for {
		z, prev = z - (z*z - x) / (2*z), z
		if math.Abs(prev - z) < 1e-8 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(196))
	fmt.Println(math.Sqrt(196))
}
