package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, k int
	fmt.Scan(&a, &b, &k)
	for i := a; i <= b; i++ {
		counter := 0
		//fmt.Println("i", i)
		for j := 1; j <= b; j++ {
			//fmt.Println("j", j)
			if i%j == 0 {
				counter += 1
				//fmt.Println("counter", counter)
				if counter == k {
					fmt.Print(strconv.Itoa(i) + " ")
					//fmt.Println("num", i)
				}
			}
		}
	}
}
