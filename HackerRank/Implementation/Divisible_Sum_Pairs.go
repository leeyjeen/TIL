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

	var n, k, count int
	fmt.Fscanln(reader, &n, &k)
	var ar = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &ar[i])
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (ar[i]+ar[j])%k == 0 {
				count++
			}
		}
	}
	fmt.Fprintln(writer, count)
}
