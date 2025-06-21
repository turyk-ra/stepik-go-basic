package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	for x > 0 {
		fmt.Print(x % 2)
		x = x / 2
	}
}
