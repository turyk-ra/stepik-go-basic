package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Print((num % 10) + (num / 10 % 10) + (num / 100 % 10))
}
