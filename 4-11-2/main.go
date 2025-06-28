package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var res string
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()
	x := sc.Text()
	intX, _ := strconv.Atoi(x)
	for intX > 0 {
		for i := 0; i <= intX; i++ {
			if i == intX%10 {
				//fmt.Println(intX % 10)
				res += strconv.Itoa(i)

			}
		}
		intX = intX / 10
	}
	if res == x {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
