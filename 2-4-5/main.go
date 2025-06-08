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
	pupilesStr := sc.Text()
	pupiles, _ := strconv.Atoi(pupilesStr)
	_ = sc.Scan()
	applesStr := sc.Text()
	apples, _ := strconv.Atoi(applesStr)
	fmt.Print(apples / pupiles)
}
