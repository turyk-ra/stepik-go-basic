package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	for x := 0; x <= 1000; x++ {
		if a*(x*x*x)+b*(x*x)+c*x+d == 0 {
			fmt.Print(x, " ")
		}
	}
}
