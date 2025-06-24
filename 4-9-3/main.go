package main

import "fmt"

func main() {
	sum, counter := 0, 0
	for i := 100; i < 1000; i++ {
		j := i
		for j > 0 {
			last := j % 10
			sum += last
			j = j / 10
		}
		if sum == 8 {
			counter += 1
		}
		sum = 0
	}
	fmt.Print(counter)
}
