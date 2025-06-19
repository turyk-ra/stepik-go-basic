package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)
	var i = 0.0
	for math.Pow(2.0, i) <= n {
		fmt.Println(math.Pow(2.0, i))
		i++
	}
}
