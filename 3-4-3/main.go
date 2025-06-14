package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	c := num / 1000
	d := num / 100 % 10
	e := num / 10 % 10
	f := num % 10

	if c == f && e == d {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
