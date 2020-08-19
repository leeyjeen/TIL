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

	var t, n, k int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n, &k)
		var a = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &a[j])
		}
		fmt.Fprintln(writer, angryProfessor(k, a))
	}
}

func angryProfessor(k int, a []int) string {
	var onTime int
	for _, v := range a {
		if v <= 0 {
			onTime++
		}
		if onTime >= k {
			return "NO"
		}
	}
	return "YES"
}
