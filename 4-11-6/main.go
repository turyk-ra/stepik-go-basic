package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	sum, sum2, sum3 := 0, 0, 0
	for x > 0 {
		step := x % 10
		sum += step
		x /= 10
	}
	for sum > 0 {
		step := sum % 10
		sum2 += step
		sum /= 10
	}
	for sum2 > 0 {
		step := sum2 % 10
		sum3 += step
		sum2 /= 10
	}
	fmt.Print(sum3)
}
