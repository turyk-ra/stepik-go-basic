package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()
	x := sc.Text()
	intX, _ := strconv.Atoi(x)
	var res, not string
	for intX > 0 {
		step := intX % 10
		if step == 7 || step == 5 {
			not += strconv.Itoa(step)
			intX /= 10
		} else {
			res = strconv.Itoa(step) + res
			intX /= 10
		}
	}
	ress, _ := strconv.Atoi(res)
	fmt.Println(ress)
}
