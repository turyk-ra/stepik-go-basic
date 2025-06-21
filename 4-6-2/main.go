package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	sumPlus, sumMinus := 0, 0
	for x != 0 {
		if x > 0 {
			sumPlus += 1
		} else {
			sumMinus += 1
		}
		fmt.Scan(&x)
	}
	fmt.Print(sumPlus - sumMinus)
}
