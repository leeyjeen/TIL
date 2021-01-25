// https://www.acmicpc.net/problem/1302
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var sold = map[string]int{}
	for i := 0; i < n; i++ {
		var title string
		fmt.Fscanln(reader, &title)
		sold[title]++
	}

	counts := []soldCount{}
	for k, v := range sold {
		counts = append(counts, soldCount{k, v})
	}
	sort.Slice(counts, func(i, j int) bool {
		if counts[i].count > counts[j].count {
			return true
		} else if counts[i].count == counts[j].count {
			return counts[i].title < counts[j].title
		}
		return false
	})
	fmt.Fprintln(writer, counts[0].title)
}

type soldCount struct {
	title string
	count int
}
