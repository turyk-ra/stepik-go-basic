package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	slice := make([]int, x)
	for i := 0; i < len(slice); i++ {
		fmt.Scan(&slice[i])
		if i%3 == 0 {
			fmt.Print(slice[i], " ")
		}
	}
}
