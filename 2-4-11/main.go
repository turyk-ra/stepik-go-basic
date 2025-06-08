package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println("Следующее за числом", num, "число:", num+1)
	fmt.Print("Для числа ", num, " предыдущее число: ", num-1)
}
