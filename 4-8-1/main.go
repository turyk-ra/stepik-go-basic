package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	max := 0
	for i := 1; i < x; i++ {
		if x%i == 0 {
			max = i
		}
	}
	fmt.Print(max)
}
