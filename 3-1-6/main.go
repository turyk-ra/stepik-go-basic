package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	if num != 0 {
		if num < 0 {
			fmt.Print(-1)
		} else {
			fmt.Print(1)
		}
	} else {
		fmt.Print(0)
	}
}
