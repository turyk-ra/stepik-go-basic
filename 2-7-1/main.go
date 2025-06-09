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
	aInt, _ := strconv.Atoi(a)
	_ = sc.Scan()
	b := sc.Text()
	bInt, _ := strconv.Atoi(b)
	_ = sc.Scan()
	c := sc.Text()
	cInt, _ := strconv.Atoi(c)

	fmt.Print((aInt + bInt + cInt), (aInt * bInt * cInt))
}
