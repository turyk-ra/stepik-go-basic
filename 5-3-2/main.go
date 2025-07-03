package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	list := make([]int, x)
	for i := range list {
		fmt.Scan(&list[i])
	}
	for k := 0; k < len(list); k++ {
		cnt := 0
		for _, num := range list {
			if num == list[k] {
				cnt += 1

			}
		}
		if cnt == 1 {
			fmt.Print(list[k], " ")
		}
	}

}
