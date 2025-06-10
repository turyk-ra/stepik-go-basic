package main

import (
	"bufio"
	"fmt"
	"math"
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

	if a != b {
		fmt.Print(math.Max(float64(aInt), float64(bInt)))
	} else {
		fmt.Print(a)
	}
}
