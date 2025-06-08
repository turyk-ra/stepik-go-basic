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
	rublesString := sc.Text()
	rubles, _ := strconv.Atoi(rublesString)
	rublesInCoins := rubles * 100
	_ = sc.Scan()
	coinsString := sc.Text()
	coins, _ := strconv.Atoi(coinsString)
	_ = sc.Scan()
	cakesString := sc.Text()
	cakes, _ := strconv.Atoi(cakesString)

	costPerOneCake := rublesInCoins + coins
	cakesCostInCoins := costPerOneCake * cakes

	fmt.Print(cakesCostInCoins/100, cakesCostInCoins%100)
}
