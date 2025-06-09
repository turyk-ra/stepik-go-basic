package main

import (
	"fmt"
)

func main() {
	var num float64
	fmt.Scan(&num)
	numInt := int(num)
	fmt.Print(num - float64(numInt))
}
