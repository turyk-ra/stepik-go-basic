package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	slice := make([]int, x)
	for i := 0; i < len(slice); i++ {
		fmt.Scan(&slice[i])
	}
	if x == 1 {
		fmt.Print(slice[0])
	} else {
		if x%2 != 0 {
			for i := 0; i < x-1; {
				tmp := slice[i+1]
				slice[i+1] = slice[i]
				slice[i] = tmp
				i += 2
			}
		} else {
			for i := 0; i < x; {
				tmp := slice[i+1]
				slice[i+1] = slice[i]
				slice[i] = tmp
				i += 2
			}
		}
		for i := 0; i < len(slice); i++ {
			fmt.Print(slice[i], " ")
		}
	}

}
