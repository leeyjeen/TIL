// https://www.acmicpc.net/problem/3085
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

	var bonbons [][]string
	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n')
		inputs := strings.Split(strings.ReplaceAll(input, "\n", ""), "")
		bonbons = append(bonbons, inputs)
	}

	var maxCount int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != n-1 {
				bonbons = swap(i, j, i, j+1, bonbons)
				maxCount = checkMax(maxCount, bonbons)
				bonbons = swap(i, j, i, j+1, bonbons)
			}

			if i != n-1 {
				bonbons = swap(i, j, i+1, j, bonbons)
				maxCount = checkMax(maxCount, bonbons)
				bonbons = swap(i, j, i+1, j, bonbons)
			}
		}
	}
	fmt.Fprintln(writer, maxCount)
}

func swap(i, j, k, l int, bonbons [][]string) [][]string {
	bonbons[i][j], bonbons[k][l] = bonbons[k][l], bonbons[i][j]
	return bonbons
}

func checkMax(maxCount int, bonbons [][]string) int {
	max := 0
	prev := ""
	for i := 0; i < len(bonbons); i++ {
		count := 0
		for j := 0; j < len(bonbons); j++ {
			cur := bonbons[i][j]
			if prev == cur {
				count++
				if max < count {
					max = count
				}
			} else {
				count = 1
			}
			prev = cur
		}
	}
	prev = ""
	for i := 0; i < len(bonbons); i++ {
		count := 0
		for j := 0; j < len(bonbons); j++ {
			cur := bonbons[j][i]
			if prev == cur {
				count++
				if max < count {
					max = count
				}
			} else {
				count = 1
			}
			prev = cur
		}
	}
	if max < maxCount {
		max = maxCount
	}
	return max
}
