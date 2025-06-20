package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		for a != 0 && b != 0 {
			a = a % b
			a, b = b, a
		}
		fmt.Println(a)
	} else {
		for a != 0 && b != 0 {
			b = b % a
			a, b = b, a
		}
		fmt.Println(b)
	}
}
