package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Scanf("%d", &count)

	var mat = make([][]int, count)
	for i := range mat {
		mat[i] = make([]int, 2)
	}

	for i:=0; i<count; i++ {
		fmt.Scanf("%d %d", &mat[i][0], &mat[i][1])
	}

	for i:=0; i<count; i++ {
		fmt.Println(mat[i][0]+mat[i][1])
	}
}