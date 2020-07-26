package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	var ar = make([]int, n)
	for i:=0 ; i<n; i++ {
		fmt.Fscanf(reader, "%d ", &ar[i])
	}

	var values = map[int]int{}
	var pair int
	for i:=0; i<n; i++ {
		if _, ok := values[ar[i]]; !ok {
			values[ar[i]] = 1
		} else {
			values[ar[i]]++
		}
	}
	for _, val := range values {
		pair += int(val/2)
	}
	fmt.Fprintln(writer, pair)
}