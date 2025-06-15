package main

import (
	"fmt"
	"math"
)

func main() {
	var d1, d2, d3 int
	fmt.Scan(&d1, &d2, &d3)

	path1 := d1 + d2 + d3
	path2 := d1 + d1 + d2 + d2
	path3 := d2 + d3 + d3 + d2
	path4 := d1 + d3 + d1 + d3

	pathRight := math.Min(math.Min(float64(path1), float64(path2)), math.Min(float64(path3), float64(path4)))
	fmt.Print(int(pathRight))
}
