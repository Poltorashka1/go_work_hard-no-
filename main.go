package main

import (
	"fmt"
)

type AmericanVelocity float64
type EuropeanVelocity float64

func main() {
	f := 120.4
	s := float64(130)
	x := AmericanVelocity(f) * 3.6
	y := EuropeanVelocity(s) * 2.23694
	fmt.Printf("%.2f \n", x)
	fmt.Printf("%.2f", y)

}
