package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	max := 0
	num := 0
	str := " NO"
	if x == 1 {
		fmt.Scan(&max)
		if max < 30 {
			str = " YES"
			fmt.Print(max, str)
		} else {
			fmt.Print(max, str)
		}
	} else {
		fmt.Scan(&max)
		if max < 30 {
			str = " YES"
		}
		for i := 1; i < x; i++ {
			fmt.Scan(&num)
			if num > max {
				max = num
			}
			if num < 30 {
				str = " YES"
			}
		}
		fmt.Print(max, str)
	}
}
