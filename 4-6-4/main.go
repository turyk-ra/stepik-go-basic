package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	counter := 0
	y := 0
	for x != 0 {
		y = x
		fmt.Scan(&x)
		if x > y {
			counter += 1
		}
	}
	fmt.Print(counter)
}
