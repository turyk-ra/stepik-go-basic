package main

import (
	"fmt"
)

func main() {
	var x1, x2, x3 int
	fmt.Scan(&x1, &x2, &x3)

	if x1 == x2 {
		if x1 == x3 {
			fmt.Print(3)
		} else {
			fmt.Print(2)
		}
	} else {
		if x1 == x3 {
			fmt.Print(2)
		} else {
			if x2 == x3 {
				fmt.Print(2)
			} else {
				fmt.Print(0)
			}
		}
	}
}
