package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var t, k, n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i:=0; i<t; i++ {
		fmt.Fscanln(reader, &k)
		fmt.Fscanln(reader, &n)
		fmt.Fprintln(writer,getCount(k, n))	
	}
}

func getCount(k, n int) (count int) {
	if k == 1 {
		for i:=0; i<n; i++ {
			count += (i+1)
		}
		return count
	}
	for i:=0; i<n; i++ {
		count += getCount(k-1, i+1)
	}
	return count
}