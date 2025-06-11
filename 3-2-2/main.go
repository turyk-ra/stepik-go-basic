package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num > -1 && num < 17 {
		fmt.Print("Принадлежит")
	} else {
		fmt.Print("Не принадлежит")
	}
}
