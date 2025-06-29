package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	max, counter := 0, 0
	for i := a; i <= b; i++ {
		if i%7 == 0 {
			max = i
			counter++
		}

	}
	if counter == 0 {
		fmt.Print("NO")
	} else {
		fmt.Print(max)
	}
}
