package main

import "fmt"

func main() {
	sum, counter := 0, 0
	for i := 1; i < 1000000; i++ {
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum = sum + j
			}
		}
		if sum == i {
			counter += 1
			fmt.Print(i, ",")
			if counter == 3 {
				break
			}
		}
		sum = 0
	}
}
