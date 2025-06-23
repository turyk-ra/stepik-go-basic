package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	var maxSum, sum float64 = 0, 0
	numOfMaxSum := 0
	for i := x; i > 1; i-- {
		j := i
		for j > 0 {
			last := j % 10
			sum += float64(last)
			j = j / 10
		}
		if sum > maxSum || (sum == maxSum && i < numOfMaxSum) {
			maxSum = sum
			numOfMaxSum = i
		}
		sum = 0
	}
	fmt.Println(numOfMaxSum, maxSum)
}
