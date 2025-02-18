package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	var z float64 = 1.0
	for i := 0; i < 10; i++ {
		if z == math.Sqrt(2) {
			return z
		}
		z -= (z*z - x) / (2 * z)
		fmt.Println("pass: ", i, "value: ", z)
	}

	return z
}

func main() {
	fmt.Println("Final ans: ", sqrt(2))
}
