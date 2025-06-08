package main

import (
	"fmt"
)

func main() {
	var minutes int
	fmt.Scan(&minutes)
	fmt.Print(minutes, " мин - это ", minutes/60, " час ", minutes%60, " минут.")
}
