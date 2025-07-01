package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	list := make([]int, x)
	pair := false
	for i := 0; i < len(list); i++ {
		fmt.Scan(&list[i])
	}
	for i := 1; i < len(list); i++ {
		if (list[i-1] > 0 && list[i] > 0) || (list[i-1] < 0 && list[i] < 0) {
			pair = true
			break
		} else {
			pair = false
		}
	}
	if pair {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
