package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	if a == 0 {
		fmt.Print("Ошибка")
	} else {
		d := math.Pow(b, 2) - (4 * a * c)
		if d > 0 {
			x1 := (-b - math.Sqrt(d)) / (2 * a)
			x2 := ((-b + math.Sqrt(d)) / (2 * a))
			fmt.Println(math.Min(x1, x2))
			fmt.Print(math.Max(x1, x2))
		} else {
			if d == 0 {
				fmt.Print(-b / (2 * a))
			}
		}
	}
}
