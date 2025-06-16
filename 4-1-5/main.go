package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		fmt.Println(math.Pow(float64(i), 2))
	}
}
