package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscanln(reader, &t)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i :=0; i<t; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		lengths := calculate(n)
		fmt.Fprintf(writer, "%d\n", lengths[n-1])
	}
}

func calculate(n int) (lengths []int) {
	lengths = append(lengths, 1, 1, 1)
	if n <= 3 {
		return
	}
	for i := 3; i < n+1 ; i++ {
		lengths = append(lengths, lengths[i-2]+ lengths[i-3])
	}
	return
}