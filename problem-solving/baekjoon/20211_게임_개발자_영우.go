// https://www.acmicpc.net/problem/20211
/*
시간초과 발생
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var input string
	fmt.Fscanln(reader, &input)
	coins := strings.Split(input, "")

	var x, y = 1, 1
	var results = [][]int{}
	for {
		var level, exp int
		for i, v := range coins {
			// 동전 던지기
			if v == "H" {
				exp += 3
			} else if v == "T" {
				if exp%2 == 0 {
					exp += 5
				} else {
					exp--
				}
			}

			// 경험치 체크
			if exp >= x {
				level++
				exp = 0
			}

			// 레벨 체크
			if level == y {
				if i == len(coins)-1 {
					results = append(results, []int{x, y})
				}
				break
			}
		}
		if y == len(coins) {
			y = 1
			x++
		} else {
			y++
		}
		if level == 0 {
			break
		}
	}

	fmt.Fprintln(writer, len(results))
	for i := 0; i < len(results); i++ {
		for j := 0; j < 2; j++ {
			fmt.Fprintf(writer, "%d ", results[i][j])
		}
		fmt.Fprintln(writer, "")
	}
}
*/

