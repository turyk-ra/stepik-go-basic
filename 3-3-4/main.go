package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()
	aStr := sc.Text()
	_ = sc.Scan()
	bStr := sc.Text()
	a, _ := strconv.ParseFloat(aStr, 64)
	b, _ := strconv.ParseFloat(bStr, 64)

	if a > 0 {
		if b > 0 {
			fmt.Print(1)
		} else {
			fmt.Print(4)
		}
	} else {
		if b > 0 {
			fmt.Print(2)
		} else {
			fmt.Print(3)
		}
	}
}
