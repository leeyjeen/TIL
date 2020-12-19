// https://www.acmicpc.net/problem/1392
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

	var n, q int
	fmt.Fscanln(reader, &n, &q)

	var scores []int
	for i := 0; i < n; i++ {
		var seconds int
		fmt.Fscanln(reader, &seconds)
		if i == 0 {
			scores = append(scores, seconds)
		} else {
			scores = append(scores, scores[i-1]+seconds)
		}
	}

	var questions []int
	for i := 0; i < q; i++ {
		var seconds int
		fmt.Fscanln(reader, &seconds)
		questions = append(questions, seconds)
	}

	for i := 0; i < q; i++ {
		var question = questions[i]
		for j := 0; j < n; j++ {
			if scores[j] > question {
				fmt.Fprintln(writer, j+1)
				break
			}
		}
	}
}
