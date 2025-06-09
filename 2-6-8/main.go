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
	_ = sc.Scan()
	y1 := sc.Text()
	_ = sc.Scan()
	x2 := sc.Text()
	_ = sc.Scan()
	y2 := sc.Text()

	x1Int, _ := strconv.Atoi(x1)
	y1Int, _ := strconv.Atoi(y1)
	x2Int, _ := strconv.Atoi(x2)
	y2Int, _ := strconv.Atoi(y2)

	distance := math.Sqrt((math.Pow(float64(x1Int)-float64(x2Int), 2)) + (math.Pow(float64(y1Int)-float64(y2Int), 2)))
	fmt.Print(distance)
}
