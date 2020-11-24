// https://www.acmicpc.net/problem/10815
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

	var cards = map[int]int{}
	for i := 0; i < n; i++ {
		var input int
		if i == n-1 {
			fmt.Fscanln(reader, &input)
		} else {
			fmt.Fscan(reader, &input)
		}
		cards[input]++
	}

	var m int
	fmt.Fscanln(reader, &m)

	for i := 0; i < m; i++ {
		var num int
		fmt.Fscanf(reader, "%d ", &num)
		fmt.Fprintf(writer, "%s ", hasCard(cards, num))
	}
}

func hasCard(cards map[int]int, num int) string {
	if cards[num] != 0 {
		return "1"
	}
	return "0"
}
