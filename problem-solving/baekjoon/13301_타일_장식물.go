// https://www.acmicpc.net/problem/13301
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

	var n int
	fmt.Fscanln(reader, &n)

	if n == 1 {
		fmt.Fprintln(writer, 4)
		return
	} else if n == 2 {
		fmt.Fprintln(writer, 6)
		return
	}

	var sides = make([]int, n)
	sides = computeSides(sides)
	fmt.Fprintln(writer, getPerimeter(sides))
}

func computeSides(sides []int) []int {
	sides[0] = 1
	sides[1] = 1
	for i := 2; i < len(sides); i++ {
		sides[i] = sides[i-1] + sides[i-2]
	}
	return sides
}

func getPerimeter(sides []int) (perimeter int) {
	a := sides[len(sides)-1]
	b := sides[len(sides)-2]
	perimeter = a*4 + b*2
	return
}
