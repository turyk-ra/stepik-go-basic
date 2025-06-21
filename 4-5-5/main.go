package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	counter := 0
	for x > 0 {
		x1 := x % 2
		if x1 == 1 {
			counter += 1
		}
		x = x / 2
	}
	fmt.Print(counter)
}
