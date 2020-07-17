package main

import (
	"os"
	"fmt"
	"bufio"
)

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0 ; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		zeroCount, oneCount := result(n)
		fmt.Fprintf(writer, "%d %d\n", zeroCount[n], oneCount[n])
	}
}

func result(n int) (zeroCount, oneCount[]int) {
	zeroCount = []int{1, 0}
	oneCount = []int{0, 1}

	if n <= 1 {
		return
	}
	for i := 2; i< n+1; i++ {
		zeroCount = append(zeroCount, zeroCount[i-1]+zeroCount[i-2])
		oneCount = append(oneCount, oneCount[i-1]+oneCount[i-2])
	}
	return
}