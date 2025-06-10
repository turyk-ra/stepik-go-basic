package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)
	if age < 12 {
		fmt.Print("Доступ запрещен")
	} else {
		fmt.Print("Доступ разрешен")
	}
}
