package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d, e float64
	counter := 0
	fmt.Scan(&a, &b, &c, &d, &e)
	for x := 0.0; x < 1001; x++ {
		if (((a * math.Pow(x, 3)) + (b * math.Pow(x, 2)) + (c * x) + d) / (x - e)) == 0 {
			counter += 1
		}
	}
	fmt.Println(counter)
}
