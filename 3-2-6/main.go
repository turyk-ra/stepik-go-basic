package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	if (num%4 == 0 && num%100 != 0) || (num%400 == 0) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
