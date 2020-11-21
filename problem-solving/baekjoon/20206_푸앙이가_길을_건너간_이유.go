// https://www.acmicpc.net/problem/20206
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

	var a, b, c float64
	fmt.Fscanln(reader, &a, &b, &c)

	var x1, x2, y1, y2 float64
	fmt.Fscanln(reader, &x1, &x2, &y1, &y2)

	fmt.Fprintln(writer, checkSecurity(a, b, c, x1, x2, y1, y2))
}

func checkSecurity(a, b, c, x1, x2, y1, y2 float64) string {
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	y := (a*x1 + c) / (-b)
	if y < y2 && y1 < y {
		return "Poor"
	}

	y = (a*x2 + c) / (-b)
	if y < y2 && y1 < y {
		return "Poor"
	}

	x := (b*y1 + c) / (-a)
	if x < x2 && x1 < x {
		return "Poor"
	}

	x = (b*y2 + c) / (-a)
	if x < x2 && x1 < x {
		return "Poor"
	}
	return "Lucky"
}
