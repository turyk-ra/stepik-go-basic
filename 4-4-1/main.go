package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	counter := 0
	for a >= 3 {
		if a%3 == 0 {
			a /= 3
			counter += 1
		} else {
			a = 0
		}
	}
	fmt.Print(counter)
}
