package main

import (
	"fmt"
	"math"
)

func main() {
	var Lenght float64 = 35
	var r = Lenght / (2 * math.Pi)
	fmt.Println(r)
	var R *float64 = &r

	S := math.Pi * math.Pow(*R, 2)
	fmt.Println(math.Round(S*1000) / 1000)
}
