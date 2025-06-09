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
	aInt, _ := strconv.Atoi(a)
	_ = sc.Scan()
	b := sc.Text()
	bInt, _ := strconv.Atoi(b)

	fmt.Print(math.Sqrt(math.Pow(float64(aInt), 2) + math.Pow(float64(bInt), 2)))
}
