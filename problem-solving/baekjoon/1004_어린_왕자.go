// https://www.acmicpc.net/problem/1004
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
		var count int
		var startX, startY, endX, endY, n int
		fmt.Fscanln(reader, &startX, &startY, &endX, &endY)
		fmt.Fscanln(reader, &n)
		for j := 0; j < n; j++ {
			var planetX, planetY, r int
			fmt.Fscanln(reader, &planetX, &planetY, &r)
			isStartContained := checkPlanetContains(startX, startY, planetX, planetY, r)
			isEndContained := checkPlanetContains(endX, endY, planetX, planetY, r)
			if isEndContained && !isStartContained || !isEndContained && isStartContained {
				count++
			}
		}
		fmt.Fprintln(writer, count)
	}

}

func checkPlanetContains(x, y, planetX, planetY, r int) bool {
	dist := (x-planetX)*(x-planetX) + (y-planetY)*(y-planetY)
	if dist > r*r {
		return false
	}
	return true
}
