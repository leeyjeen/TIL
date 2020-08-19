package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		fmt.Fprintln(writer, findDigits(n))
	}
}

func findDigits(n int) int {
	digit := n % 10
	dividend := n
	var count = 0
	for dividend != 0 {
		if digit != 0 && n%digit == 0 {
			count++
		}
		dividend /= 10
		digit = dividend % 10
	}
	return count
}
