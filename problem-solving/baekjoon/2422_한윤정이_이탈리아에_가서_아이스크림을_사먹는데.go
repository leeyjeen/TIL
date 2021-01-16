// https://www.acmicpc.net/problem/2422
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

	var n, m int
	fmt.Fscanln(reader, &n, &m)
	var badCombination = make([][]bool, n)
	for i := 0; i < n; i++ {
		badCombination[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		var iceOne, iceTwo int
		fmt.Fscanln(reader, &iceOne, &iceTwo)
		badCombination[iceOne-1][iceTwo-1] = true
		badCombination[iceTwo-1][iceOne-1] = true
	}

	fmt.Fprintln(writer, getCountOfGoodCombination(badCombination))
}

func getCountOfGoodCombination(badCombination [][]bool) (count int) {
	for i := 0; i < len(badCombination)-2; i++ {
		for j := i + 1; j < len(badCombination)-1; j++ {
			if badCombination[i][j] {
				continue
			}
			for k := j + 1; k < len(badCombination); k++ {
				if badCombination[i][k] || badCombination[j][k] {
					continue
				}
				count++
			}
		}
	}
	return
}
