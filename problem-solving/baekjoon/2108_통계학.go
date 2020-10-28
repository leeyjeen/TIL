package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	var n, sum, modeCnt int
	var max = -4000
	var min = 4000
	var counts = make([]int, 8001)
	var modes []int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &numbers[i])
		sum += numbers[i]
		if min > numbers[i] {
			min = numbers[i]
		}
		if max < numbers[i] {
			max = numbers[i]
		}
		var idx = numbers[i] + 4000
		counts[idx]++
		if modeCnt < counts[idx] {
			modes = []int{numbers[i]}
			modeCnt = counts[idx]
		} else if modeCnt == counts[idx] {
			modes = append(modes, numbers[i])
		}
	}

	sort.Ints(numbers)
	sort.Ints(modes)

	fmt.Fprintln(writer, math.Round(float64(sum)/float64(n)))
	fmt.Fprintln(writer, numbers[n/2])
	if len(modes) > 1 {
		fmt.Fprintln(writer, modes[1])
	} else {
		fmt.Fprintln(writer, modes[0])
	}
	fmt.Fprintln(writer, max-min)
}
