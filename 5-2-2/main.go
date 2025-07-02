package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	list := make([]int, x)
	for i := range list {
		fmt.Scan(&list[i])
	}
	minNum := list[0]
	for i := 1; i < len(list); i++ {
		if minNum > list[i] {
			minNum = list[i]
		}
	}
	for i := range list {
		list[i] -= minNum
		fmt.Print(list[i], " ")
	}
}
