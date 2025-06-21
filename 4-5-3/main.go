package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	for x > 0 {
		step := x % 10
		i := step
		x = x / 10
		fmt.Print(i)
	}
}
