package main

import "fmt"

func main() {
	var w int
	fmt.Scan(&w)

	if w%2 == 0 && w != 2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
