// https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	var c = make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Fscanf(reader, "%d ", &c[i])
	}
	fmt.Fprintln(writer, jumpingOnClouds(c))
}

func jumpingOnClouds(c []int) int {
	var leap int = 0
	for i:=0; i<len(c)-1; i++ {
		if c[i] == 0 {
			if i+2 <= len(c)-1 && c[i+2] == 0 {
				leap++
				i++
			} else if i+1 <= len(c)-1 && c[i+1] == 0 {
				leap++
			}
		}
	}	
	return leap
}