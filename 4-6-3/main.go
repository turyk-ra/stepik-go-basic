package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	sum, counter := 0, 0
	for x != 0 {
		counter += 1
		sum += x
		fmt.Scan(&x)
	}
	fmt.Print(float64(sum) / float64(counter))
}
