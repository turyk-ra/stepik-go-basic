package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	i := 1
	c := 0
	for i < n {
		i = i * 2
		c += 1
	}
	fmt.Println(c)
}
