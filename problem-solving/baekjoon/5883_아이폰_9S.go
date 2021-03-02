// https://www.acmicpc.net/problem/5883
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
	var b = make([]int, n)

	keyMap := map[int]bool{}
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &b[i])
		keyMap[b[i]] = true
	}

	var maxContinuousCount = 1
	for key := range keyMap {
		keyToDelete := key
		var newLine []int
		for i := 0; i < n; i++ {
			if b[i] != keyToDelete {
				newLine = append(newLine, b[i])
			}
		}
		if len(newLine) == 0 {
			maxContinuousCount = 0
			break
		}
		var prev = newLine[0]
		var continuousCount = 1
		for i := 1; i < len(newLine); i++ {
			var cur = newLine[i]
			if cur == prev {
				continuousCount++
				if maxContinuousCount < continuousCount {
					maxContinuousCount = continuousCount
				}
			} else {
				continuousCount = 1
			}
			prev = cur
		}
	}
	fmt.Fprintln(writer, maxContinuousCount)
}
