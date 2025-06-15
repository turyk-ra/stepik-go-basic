package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	if (x > -30 && x <= -2) || (x > 7 && x <= 25) {
		fmt.Print("Принадлежит")
	} else {
		fmt.Print("Не принадлежит")
	}
}
