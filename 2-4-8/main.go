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
	device := sc.Text()
	deviceCost, _ := strconv.Atoi(device)
	_ = sc.Scan()
	deviceCase := sc.Text()
	deviceCaseCost, _ := strconv.Atoi(deviceCase)
	_ = sc.Scan()
	charger := sc.Text()
	chargerCost, _ := strconv.Atoi(charger)
	_ = sc.Scan()
	headphones := sc.Text()
	headphonesCost, _ := strconv.Atoi(headphones)
	fmt.Print((deviceCost + deviceCaseCost + chargerCost + headphonesCost) * 3)
}
