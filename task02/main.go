package main

import (
	"fmt"
	"math"
)

func main() {
	type AmericanVelocity float64
	type EuropeanVelocity float64

	var speedAmerican AmericanVelocity
	var speedEuropean EuropeanVelocity

	speedAmerican = 120.4 * 3.6
	fmt.Println(speedAmerican)

	speedEuropean = EuropeanVelocity(130 * 3.6 / 1.609)
	speedEuropean = EuropeanVelocity(math.Round(float64(speedEuropean*100)) / 100)
	fmt.Println(speedEuropean)
}
