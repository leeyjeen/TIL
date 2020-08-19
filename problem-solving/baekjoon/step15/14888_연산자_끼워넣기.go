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

	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}

	var counts = make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Fscanf(reader, "%d ", &counts[i])
	}

	var max, min, sum, idx = -1000000000, 1000000000, numbers[0], 0
	solve(numbers, counts, &max, &min, sum, idx)

	fmt.Fprintln(writer, max)
	fmt.Fprintln(writer, min)
}

func solve(numbers, counts []int, maxptr, minptr *int, sum, idx int) {
	if idx == len(numbers)-1 {
		if *maxptr < sum {
			*maxptr = sum
		}
		if *minptr > sum {
			*minptr = sum
		}
		return
	}

	if counts[0] > 0 {
		counts[0]--
		solve(numbers, counts, maxptr, minptr, sum+numbers[idx+1], idx+1)
		counts[0]++
	}

	if counts[1] > 0 {
		counts[1]--
		solve(numbers, counts, maxptr, minptr, sum-numbers[idx+1], idx+1)
		counts[1]++
	}

	if counts[2] > 0 {
		counts[2]--
		solve(numbers, counts, maxptr, minptr, sum*numbers[idx+1], idx+1)
		counts[2]++
	}

	if counts[3] > 0 {
		counts[3]--
		solve(numbers, counts, maxptr, minptr, sum/numbers[idx+1], idx+1)
		counts[3]++
	}
}
