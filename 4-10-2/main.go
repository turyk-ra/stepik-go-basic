package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	var max int
	var num int
	fmt.Scan(&max)
	for i := 1; i < x; i++ {
		fmt.Scan(&num)
		if num > max {
			max = num
		}
	}
	fmt.Print(max)
}
