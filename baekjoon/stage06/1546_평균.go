package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var max, sum int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var scores = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &scores[i])
		sum += scores[i]
		if max < scores[i] {
			max = scores[i]
		}
	}

	fmt.Println(float64(sum) / float64(n) / float64(max) * 100.0)
}
