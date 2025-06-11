package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	a := num / 100000
	b := num / 10000 % 10
	c := num / 1000 % 10
	d := num / 100 % 10
	e := num / 10 % 10
	f := num % 10
	if (a + b + c) == (d + e + f) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
