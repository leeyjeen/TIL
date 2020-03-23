package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var mat = make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &mat[i][0], &mat[i][1])
	}

	sort.Slice(mat, func(i, j int) bool {
		for idx := range mat[i] {
			if mat[i][idx] == mat[j][idx] {
				continue
			}
			return mat[i][idx] < mat[j][idx]
		}
		return false
	})

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d %d\n", mat[i][0], mat[i][1])
	}
}
