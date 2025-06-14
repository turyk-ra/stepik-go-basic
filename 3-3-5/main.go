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
	a, _ := strconv.ParseFloat(aStr, 64)
	_ = sc.Scan()
	bStr := sc.Text()
	b, _ := strconv.ParseFloat(bStr, 64)
	_ = sc.Scan()
	symbol := sc.Text()

	if symbol == "/" {
		if b == 0 {
			fmt.Print("На ноль делить нельзя!")
		} else {
			fmt.Print((a / b))
		}
	} else {
		if symbol == "*" {
			fmt.Print((a * b))
		} else {
			if symbol == "+" {
				fmt.Print((a + b))
			} else {
				if symbol == "-" {
					fmt.Print((a - b))
				} else {
					fmt.Print("Неверная операция")
				}
			}
		}
	}
}
