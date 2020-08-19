package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var n, a, count int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	for i:=0; i<n; i++ {
		fmt.Fscanf(reader, "%d ", &a)

		var divisorCount int
		for j:=0; j<a; j++ {
			if a % (j+1) == 0 {
				divisorCount++
			}
		}
		if divisorCount == 2 {
			count++
		}
	}

	fmt.Println(count)
}