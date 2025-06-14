package main

import (
	"fmt"
)

func main() {
	var weight int
	fmt.Scan(&weight)
	if weight < 60 {
		fmt.Print("Легкий вес")
	} else {
		if weight < 64 {
			fmt.Print("Первый полусредний вес")
		} else {
			fmt.Print("Полусредний вес")
		}
	}
}
