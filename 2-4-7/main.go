package main

import (
	"fmt"
)

func main() {
	var first int
	fmt.Scan(&first)
	fmt.Println(first)
	first++
	second := first
	fmt.Println(second)
	second++
	third := second
	fmt.Print(third)
}
