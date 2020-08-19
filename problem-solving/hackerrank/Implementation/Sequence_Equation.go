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
	var p = make([]int, n)
	for i := range p {
		fmt.Fscanf(reader, "%d ", &p[i])
	}
	for _, v := range permutationEquation(p) {
		fmt.Fprintln(writer, v)
	}
}

func permutationEquation(p []int) []int {
	pInverse := make([]int, len(p)) // p()는 일대일 대응 함수이므로 역함수가 성립함
	for i, v := range p {
		pInverse[v-1] = i + 1
	}
	var result = []int{}
	for i := range p {
		result = append(result, pInverse[pInverse[i]-1])
	}
	return result
}
