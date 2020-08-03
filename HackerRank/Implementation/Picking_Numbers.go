package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &a[i])
	}
	fmt.Fprintln(writer, pickingNumbers(a))
}

func pickingNumbers(a []int) int {
	numberMap := map[int]int{}
	for _, v := range a {
		if _, ok := numberMap[v]; ok {
			numberMap[v]++
		} else {
			numberMap[v] = 1
		}
	}
	var maxLen int
	for i := 0; i < len(a); i++ {
		var length = int(math.Max(float64(numberMap[a[i]]+numberMap[a[i]-1]), float64(numberMap[a[i]]+numberMap[a[i]+1])))
		if length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}
