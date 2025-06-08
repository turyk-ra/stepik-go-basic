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
	floatA, _ := strconv.ParseFloat(a, 64)
	_ = sc.Scan()
	b := sc.Text()
	floatB, _ := strconv.ParseFloat(b, 64)
	fmt.Print(0.5 * floatA * floatB)
}
