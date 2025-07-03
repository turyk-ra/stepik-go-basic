package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	list := make([]int, x)
	for i := range list {
		fmt.Scan(&list[i])
	}
	arg := false
	for j := 10001; j >= -10000; j-- {
		cnt := 0
		for i := range list {
			if j == list[i] {
				cnt += 1
			}
		}
		if cnt >= 2 {
			arg = true
			break
		} else {
			arg = false
		}
	}
	if arg {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
