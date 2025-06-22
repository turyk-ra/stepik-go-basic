package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	counter := 0
	for i := x; i >= 1; i-- {
		j := i
		//fmt.Println("i-", i)
		for j > 0 {
			last := j % 10
			//fmt.Println("last-", last)
			if last == 7 {
				counter += 1
			}
			j = j / 10
			//fmt.Println("i in loop-", j)
		}
	}
	fmt.Println(counter)
}
