package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &n)
	defer writer.Flush()

	var starMat = make([][]bool, n)
	for i := 0; i < n; i++ {
		starMat[i] = make([]bool, n)
	}

	makeStar(n, 0, 0, starMat)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if starMat[i][j] {
				fmt.Fprint(writer, "*")
			} else {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprint(writer, "\n")
	}
}

func makeStar(size, row, col int, starMat [][]bool) {
	if size == 3 {
		var cur = 0
		for i := row; i < row+size; i++ {
			for j := col; j < col+size; j++ {
				cur++
				if cur == 5 {
					starMat[i][j] = false
				} else {
					starMat[i][j] = true
				}
			}
		}
		return
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				continue
			} else {
				makeStar(size/3, row+i*size/3, col+j*size/3, starMat)
			}
		}
	}
}
