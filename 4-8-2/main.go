package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x >= 10 && x <= 100 {
		fmt.Println(x)
	}
	for x <= 100 {
		fmt.Scan(&x)
		if x < 10 {
			continue
		} else if x > 100 {
			break
		} else {
			fmt.Println(x)
		}
	}
}
