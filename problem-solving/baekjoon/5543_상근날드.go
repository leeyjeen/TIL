package main

import (
	"fmt"
)

func main() {
	burgers := make([]int, 3)
	for i := range burgers {
		fmt.Scanf("%d", &burgers[i])
	}

	beverages := make([]int, 2)
	for i := range beverages {
		fmt.Scanf("%d", &beverages[i])
	}

	var minBurger, minBeverage = 2000, 2000
	for _, burger := range burgers {
		if burger < minBurger {
			minBurger = burger
		}
	}
	for _, beverage := range beverages {
		if beverage < minBeverage {
			minBeverage = beverage
		}
	}
	fmt.Println(minBurger + minBeverage - 50)
}