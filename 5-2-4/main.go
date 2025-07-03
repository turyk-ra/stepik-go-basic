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
	// fmt.Println(numMin, indexMin)
	indexMax := 0
	numMax := list[indexMin]
	for i := x - 1; i > 0; i-- {
		if numMax < list[i] {
			numMax = list[i]
			indexMax = i
		}
	}
	// fmt.Println(numMax, indexMax)
	tmp := list[indexMin]
	list[indexMin] = list[indexMax]
	list[indexMax] = tmp
	for i := range list {
		fmt.Print(list[i], " ")
	}
}
