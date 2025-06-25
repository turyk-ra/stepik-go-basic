package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	var min, max, num int
	if x == 1 {
		fmt.Scan(&min)
		fmt.Print(0)
	} else {
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
}
