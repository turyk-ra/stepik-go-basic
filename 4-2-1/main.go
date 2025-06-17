package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	sum := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			sum++
		}
	}
	fmt.Print(sum)
}
