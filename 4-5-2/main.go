package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	sum := 0
	full := x
	for x > 0 {
		step := x % 10
		sum += step
		x = x / 10
	}
	if full%sum == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
