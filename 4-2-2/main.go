package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Scan(&count)
	sum := 0
	for i := 0; i < count; i++ {
		var c int
		fmt.Scan(&c)
		sum += c
	}
	fmt.Print(sum)
}
