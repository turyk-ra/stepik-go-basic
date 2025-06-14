package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	num = num % 10
	if num%2 == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
