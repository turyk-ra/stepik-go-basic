package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	_ = sc.Scan()
	a := sc.Text()
	_ = sc.Scan()
	b := sc.Text()

	if a != b {
		fmt.Print("Пароль не принят")
	} else {
		fmt.Print("Пароль принят")
	}
}
