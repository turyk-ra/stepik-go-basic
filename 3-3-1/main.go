package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if (x >= -3 && x <= 1) || (x >= 5 && x <= 9) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
