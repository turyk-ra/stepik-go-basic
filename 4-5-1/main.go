package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	counter := 0
	for x > 0 {
		step := x % 10
		if step%4 == 0 && step != 0 && step/4 == 1 {
			counter++
		}
		x = x / 10
	}
	fmt.Print(counter)
}
