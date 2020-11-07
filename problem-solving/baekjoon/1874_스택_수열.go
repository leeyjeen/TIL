// https://www.acmicpc.net/problem/1874
// 스택을 활용하는 문제
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

	var stackNum int = 1
	var stackSequence []int
	var result []string
	for i := 0; i < n; i++ {
		var input int
		fmt.Fscanln(reader, &input)
		if len(stackSequence) > 0 && input > stackSequence[len(stackSequence)-1] || input > stackNum {
			for j := stackNum; j < input+1; j++ {
				stackSequence = append(stackSequence, j)
				result = append(result, "+")
				stackNum++
			}
			stackSequence = stackSequence[:len(stackSequence)-1]
			result = append(result, "-")
		} else if input == stackNum {
			stackSequence = append(stackSequence, input)
			result = append(result, "+")
			stackNum++
			stackSequence = stackSequence[:len(stackSequence)-1]
			result = append(result, "-")
		} else if len(stackSequence) > 0 && input == stackSequence[len(stackSequence)-1] {
			stackSequence = stackSequence[:len(stackSequence)-1]
			result = append(result, "-")
		}
	}

	if len(stackSequence) > 0 {
		fmt.Fprintln(writer, "NO")
	} else {
		for _, v := range result {
			fmt.Fprintln(writer, v)
		}
	}
}
