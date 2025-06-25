package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	num, counter := 0, 0
	for i := 0; i < x; i++ {
		fmt.Scan(&num)
		if num == 0 {
			counter += 1
		}
	}
	fmt.Print(counter)
}
