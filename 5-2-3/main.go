package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	list := make([]int, x)
	for i := range list {
		fmt.Scan(&list[i])
	}
	indexMin := 0
	numMin := list[indexMin]
	for i := 1; i < len(list); i++ {
		if numMin > list[i] {
			numMin = list[i]
			indexMin = i
		}
	}
	indexMax := 0
	numMax := list[indexMax]
	for i := 1; i < len(list); i++ {
		if numMax < list[i] {
			numMax = list[i]
			indexMax = i
		}
	}
	fmt.Print(indexMax - indexMin)
}
