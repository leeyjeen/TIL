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

	var n int
	fmt.Fscanln(reader, &n)

	var steps = make([][]int, n+1)
	for i := range steps {
		steps[i] = make([]int, 10)
	}

	for i := 1; i < 10; i++ {
		steps[1][i] = 1
	}

	for i := 2; i < n+1; i++ {
		for j := 0; j < 10; j++ {
			if j == 0 {
				steps[i][j] = steps[i-1][j+1] % 1000000000
			} else if j == 9 {
				steps[i][j] = steps[i-1][j-1] % 1000000000
			} else {
				steps[i][j] = (steps[i-1][j+1] + steps[i-1][j-1]) % 1000000000
			}
		}
	}
	var result int
	for i := 0; i < 10; i++ {
		result += steps[n][i]
	}
	fmt.Fprintln(writer, result%1000000000)
}