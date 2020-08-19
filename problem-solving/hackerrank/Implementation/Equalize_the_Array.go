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
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
	}
	fmt.Fprintln(writer, equalizeArray(arr))
}

type Pair struct {
	Key   int
	Value int
}

func equalizeArray(arr []int) int {
	arrMap := map[int]int{}
	for _, v := range arr {
		if _, ok := arrMap[v]; ok {
			arrMap[v]++
		} else {
			arrMap[v] = 1
		}
	}

	pairs := []Pair{}
	for k, v := range arrMap {
		pairs = append(pairs, Pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})
	return len(arr) - pairs[0].Value
}
