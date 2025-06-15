package main

import "fmt"

func main() {
	var num float64
	fmt.Scan(&num)

	if num == 0 {
		fmt.Print("Обратного числа не существует")
	} else {
		fmt.Print(1 / num)
	}
}
