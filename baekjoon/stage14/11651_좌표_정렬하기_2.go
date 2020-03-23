package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type coordinate struct {
	x int
	y int
}

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var mat = make([]coordinate, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &mat[i].x, &mat[i].y)
	}

	sort.Slice(mat, func(i, j int) bool {
		if mat[i].y < mat[j].y {
			return true
		} else if mat[i].y == mat[j].y {
			return mat[i].x < mat[j].x
		} else {
			return false
		}
	})

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d %d\n", mat[i].x, mat[i].y)
	}
}
