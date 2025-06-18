package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var result string
	var y int
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			y += 1
		}
	}
	if y > 0 {
		result = "YES"
	} else {
		result = "NO"
	}
	fmt.Print(result)
}
