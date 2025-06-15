package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num%2 != 0 {
		fmt.Print("YES")
	} else {
		if num >= 6 && num <= 20 {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	}
}
