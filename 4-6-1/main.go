package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	sum := 0
	for x != 0 {
		if x%2 == 0 && x%3 != 0 {
			sum += x
		}
		fmt.Scan(&x)
	}
	fmt.Print(sum)
}
