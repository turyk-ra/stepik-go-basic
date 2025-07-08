package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	list := make([][]int, x)
	for i := 0; i < x; i++ {
		list[i] = make([]int, x)
		for j := x; j > 0; j-- {
			if j == len(list[i])-1 {
				list[i][j] = 1
			}
		}
	}
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			fmt.Print(list[i][j], " ")
		}
		fmt.Println()
	}
}
