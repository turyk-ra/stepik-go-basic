package main

import (
	"fmt"
	"math"
)

func main() {
	var bit float64
	fmt.Scan(&bit)
	fmt.Print((bit / math.Pow(2, 3)) / math.Pow(2, 10))
}
