package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	list := make([]int, x)
	for i := 0; i < len(list); i++ {
		fmt.Scan(&list[i])
	}
	for i := x - 1; i > 0; {
		tmp := list[i-1]
		list[i-1] = list[i]
		list[i] = tmp
		i -= 1
	}
	for i := 0; i < len(list); i++ {
		fmt.Print(list[i], " ")
	}
}
