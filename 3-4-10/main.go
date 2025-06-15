package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	numOfDesksA := math.Round(a / 2.0)
	numOfDesksB := math.Round(b / 2.0)
	numOfDesksC := math.Round(c / 2.0)
	fmt.Print(numOfDesksA + numOfDesksB + numOfDesksC)
}
