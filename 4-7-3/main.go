package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	sum := 0
	for i := x; i > 1; i-- {
		j := i
		//fmt.Println("j---", j)
		for j > 0 {
			last := j % 10
			//fmt.Println("last ---", last)
			sum += last
			//fmt.Println("sum ---", sum)
			j = j / 10
			//fmt.Println("j again ==", j)
		}
		fmt.Println(i, sum)
		sum = 0
		//TODO
	}

}
