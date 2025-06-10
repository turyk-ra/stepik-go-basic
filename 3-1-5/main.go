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
	a := sc.Text()
	_ = sc.Scan()
	b := sc.Text()

	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)

	if aInt%bInt == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
