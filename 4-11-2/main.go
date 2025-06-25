package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	num := 0
	for x > 0 {
		last := x % 10
		x = x / 10
	}
	//TODO
}
