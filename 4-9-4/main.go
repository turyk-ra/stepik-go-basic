package main

import "fmt"

func main() {
	add := 1
	for i := 10; i < 100; i++ {
		j := i
		for j > 0 {
			last := j % 10
			add *= last
			j /= 10
		}
		if add*2 == i {
			fmt.Print(i, ",")
		}
		add = 1
	}
}
