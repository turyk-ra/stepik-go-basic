package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)
	st1 := strconv.Itoa(num % 10)
	st2 := strconv.Itoa(num / 10 % 10)
	st3 := strconv.Itoa(num / 100 % 10)
	fmt.Print(st1, st2, st3)
}
