package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	list := make([]int, x)
	for i := range list {
		fmt.Scan(&list[i])
	}
	counter := 1
	for i := 0; i < len(list); i++ {
		if list[i] == list[x-counter] {
			counter += 1
		}
	}
	if counter == x+1 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
