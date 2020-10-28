package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &n, &m)
	defer writer.Flush()

	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}

	var nearest int
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				var val = numbers[i] + numbers[j] + numbers[k]
				if val <= m && nearest < val {
					nearest = val
				} else {
					continue
				}
			}
		}
	}
	fmt.Println(nearest)
}
