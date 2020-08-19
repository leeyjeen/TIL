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
		fmt.Fprintln(writer, utopianTree(n))
	}
}

func utopianTree(n int) int {
	var h = 1
	for i := 1; i < n+1; i++ {
		if i%2 == 1 {
			h *= 2
		} else {
			h++
		}
	}
	return h
}
