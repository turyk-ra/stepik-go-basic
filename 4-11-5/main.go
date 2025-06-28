package main

import "fmt"

func main() {
	var x, min, num int
	fmt.Scan(&x)
	fmt.Scan(&min)
	counter := 1
	for i := 1; i < x; i++ {
		fmt.Scan(&num)
		if num < min {
			min = num
			counter = 1
		} else if num == min {
			counter += 1
		}
	}
	fmt.Print(counter)
}
