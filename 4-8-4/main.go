package main

import (
	"fmt"
)

func main() {
	var x string
	fmt.Scan(&x)
	counter := 0
	a := x
	for i := 1; i > 0; {
		fmt.Scan(&x)
		counter += 1
		if a == x {
			fmt.Print(counter)
			break
		}
	}
}
