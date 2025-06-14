package main

import (
	"fmt"
	"math"
)

func main() {
	var k, m float64
	fmt.Scan(&k, &m)

	if k <= m {
		fmt.Print(1)
	} else {
		fmt.Print(math.Ceil(k / m))
	}
}
