package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	a := num % 10
	b := num % 100 / 10
	c := num % 1000 / 100
	if a != b && a != c && b != c {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
