package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	if num <= 13 {
		fmt.Print("детство")
	} else if num >= 14 && num <= 24 {
		fmt.Print("молодость")
	} else if num >= 25 && num <= 59 {
		fmt.Print("зрелость")
	} else {
		fmt.Print("старость")
	}
}
