package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Print(num % 1000 / 100)
}
