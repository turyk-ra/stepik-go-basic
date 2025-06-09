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
	floatA, _ := strconv.ParseFloat(a, 64)
	_ = sc.Scan()
	b := sc.Text()
	floatB, _ := strconv.ParseFloat(b, 64)

	c := math.Sqrt(math.Pow(floatA, 2) + math.Pow(floatB, 2))
	fmt.Print(floatA + floatB + c)
}
