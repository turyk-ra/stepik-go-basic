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
	for i := 1; i < x; i++ {
		if numMin > list[i] {
			numMin = list[i]
			indexMin = i
			break
		}
	}

}
