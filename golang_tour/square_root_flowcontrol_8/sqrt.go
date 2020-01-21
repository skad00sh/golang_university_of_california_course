package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.
	t := 0.
	const Delta = 1e-3
	for i := 0; i < 999999999; i++ {

		z -= (z*z - x) / (2 * z)

		fmt.Printf("Iteration = %d,t = %f, z = %f\n", i, t, z)

		if t == z || math.Abs(z-t) < Delta {
			break
		}
		t = z

	}
	return z
}

func main() {
	var ip float64
	fmt.Printf("\n Enter the number =  ")
	fmt.Scan(&ip)
	fmt.Println(Sqrt(ip))
}
