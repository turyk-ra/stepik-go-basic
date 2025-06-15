package main

import "fmt"

func main() {
	var age int
	var sex string
	fmt.Scan(&age, &sex)

	switch sex {
	case "f":
		fmt.Print("NO")
	case "m":
		if age >= 12 && age <= 18 {
			fmt.Print("YES")
		} else {
			fmt.Print("NO")
		}
	}
}
