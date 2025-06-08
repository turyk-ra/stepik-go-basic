package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := bufio.NewScanner(os.Stdin)
	_ = a.Scan()
	first := a.Text()
	_ = a.Scan()
	second := a.Text()
	_ = a.Scan()
	third := a.Text()

	fmt.Println(third)
	fmt.Println(second)
	fmt.Print(first)
}
