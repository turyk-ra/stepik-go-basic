package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	y, counter := 0, 0
	for x != 0 {
		y = x
		fmt.Scan(&x)
		if (x > 0 && y < 0) || (x < 0 && y > 0) {
			counter += 1
		}
	}
	fmt.Print(counter)
}
