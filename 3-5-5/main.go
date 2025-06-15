package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	if num >= 16 {
		fmt.Print("старший разработчик")
	} else if num >= 8 && num <= 15 {
		fmt.Print("средний разработчик")
	} else if num >= 4 && num <= 7 {
		fmt.Print("младший разработчик")
	} else {
		fmt.Print("начинающий")
	}
}
