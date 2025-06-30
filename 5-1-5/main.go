package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	list := make([]int, x)
	//num := list[0]
	for i := 0; i < len(list); i++ {
		fmt.Scan(&list[i])
	}
}
