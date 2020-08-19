package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var mat []int
	for i := n; i != 0; i /= 10 {
		mat = append(mat, i%10)
	}

	for i := 0; i < len(mat)-1; i++ {
		for j := i + 1; j < len(mat); j++ {
			var tmp = mat[i]
			if tmp < mat[j] {
				mat[i] = mat[j]
				mat[j] = tmp
			}
		}
	}

	for i := 0; i < len(mat); i++ {
		fmt.Print(mat[i])
	}
	fmt.Println()
}
