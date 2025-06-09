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
	x1 := sc.Text()
	x1Int, _ := strconv.Atoi(x1)
	_ = sc.Scan()
	y1 := sc.Text()
	y1Int, _ := strconv.Atoi(y1)
	_ = sc.Scan()
	x2 := sc.Text()
	x2Int, _ := strconv.Atoi(x2)
	_ = sc.Scan()
	y2 := sc.Text()
	y2Int, _ := strconv.Atoi(y2)

	distance := math.Abs(float64(x1Int)-float64(x2Int)) + math.Abs(float64(y1Int)-float64(y2Int))

	fmt.Print(distance)
}
