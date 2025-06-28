package main

import "fmt"

func main() {
	var a, n int
	fmt.Scan(&a, &n)
	switch n {
	case 0:
		fmt.Print(1)
	case 1:
		fmt.Print(a)
	default:
		tmp := a
		for i := 1; i < n; i++ {
			tmp = tmp * a
		}
		fmt.Print(tmp)
	}
}
