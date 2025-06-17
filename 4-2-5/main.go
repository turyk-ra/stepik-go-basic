package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var result string
	for i := 0; i < n; i++ {
		var x float64
		fmt.Scan(&x)
		if x == 0 {
			result = "YES"
		} else {
			result = "NO"
		}
	}
	// ошибка
	fmt.Print(result)
}
