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
	aStr := sc.Text()
	_ = sc.Scan()
	bStr := sc.Text()
	aFloat, _ := strconv.ParseFloat(aStr, 64)
	bFloat, _ := strconv.ParseFloat(bStr, 64)

	fmt.Print((aFloat + bFloat) / 2.0)
}
