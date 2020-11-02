// https://www.acmicpc.net/problem/9375
// 경우의 수 연습문제
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

		var clothesMap = make(map[string]int, n)
		for j := 0; j < n; j++ {
			var name, kind string
			fmt.Fscanln(reader, &name, &kind)
			clothesMap[kind]++
		}
		var result int = 1
		for _, value := range clothesMap {
			result *= (value + 1)
		}
		result--
		fmt.Fprintln(writer, result)
	}
}
