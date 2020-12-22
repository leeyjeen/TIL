// https://www.acmicpc.net/problem/4573
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var menu int
	for {
		menu++
		var n int
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}
		var selectedDiamater int
		var lowPricePerSquare = math.MaxFloat64
		for i := 0; i < n; i++ {
			var diameter, price int
			fmt.Fscanln(reader, &diameter, &price)
			var radius = float64(diameter) / 2.0
			var pricePerSquare = float64(price) / (radius * radius * math.Phi)
			if pricePerSquare < lowPricePerSquare {
				lowPricePerSquare = pricePerSquare
				selectedDiamater = diameter
			}
		}
		fmt.Fprintf(writer, "Menu %d: %d\n", menu, selectedDiamater)
	}
}
