package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m)

	var mat = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%v\n", &mat[i])
	}
	var min = n * m
	for i := 0; i < n-7; i++ {
		for j := 0; j < m-7; j++ {
			var cntChange1 = float64(0)
			var cntChange2 = float64(0)
			for k := i; k < i+8; k++ {
				for l := j; l < j+8; l++ {
					if (k+l)%2 == 0 {
						if string(mat[k][l]) == "B" {
							cntChange1++
						} else {
							cntChange2++
						}
					} else {
						if string(mat[k][l]) == "B" {
							cntChange2++
						} else {
							cntChange1++
						}
					}
				}
			}
			if min > int(math.Min(cntChange1, cntChange2)) {
				min = int(math.Min(cntChange1, cntChange2))
			}
		}
	}
	fmt.Println(min)
}
