package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	list := make([]int, x)
	for i := 0; i < len(list); i++ {
		fmt.Scan(&list[i])
	}
	indexMin := 0
	minNum := list[indexMin]
	for i := 1; i < len(list); i++ {
		if minNum > list[i] {
			minNum = list[i]
			indexMin = i
		}
	}
	fmt.Print(indexMin)
}
