package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	if ((x1 == x2 && y1 != y2) || (y1 == y2 && x1 != x2)) || (math.Abs(x2-x1) == math.Abs(y2-y1)) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
