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
	frst := sc.Text()
	first, _ := strconv.Atoi(frst)
	_ = sc.Scan()
	scnd := sc.Text()
	second, _ := strconv.Atoi(scnd)
	_ = sc.Scan()
	thrd := sc.Text()
	third, _ := strconv.Atoi(thrd)
	fmt.Print(first * second * third)
}
