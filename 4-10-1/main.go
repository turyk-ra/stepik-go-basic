package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	var min int
	var num int
	fmt.Scan(&min)
	for i := 1; i < x; i++ {
		fmt.Scan(&num)
		if num < min {
			min = num
		}
	}
	fmt.Print(min)
}
