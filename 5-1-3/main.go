package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	slice := make([]int, x)
	var num int
	counter := 0
	fmt.Scan(&num)
	slice = append(slice, num)
	for i := 1; i < len(slice)-1; i++ {
		fmt.Scan(&slice[i])
		if slice[i] > num {
			counter++
		}
		num = slice[i]
	}
	fmt.Print(counter)
}
