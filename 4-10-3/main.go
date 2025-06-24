package main

import "fmt"

func main() {
	//TODO
	var x int
	fmt.Scan(&x)
	var min, max, num int
	fmt.Scan(&min)
	fmt.Scan(&max)
	if min > max {
		t := max
		max = min
		min = t
	}
	for i := 2; i < x; i++ {
		fmt.Scan(&num)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Print(max - min)
}
