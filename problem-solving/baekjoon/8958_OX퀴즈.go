package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)

	var results = make([]string, t)
	var scores = make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &results[i])
		var correct, score int
		for j := 0; j < len(results[i]); j++ {
			if fmt.Sprintf("%c", results[i][j]) == "O" {
				correct++
				score += correct
			} else {
				correct = 0
			}
		}
		scores[i] = score
	}

	for _, val := range scores {
		fmt.Println(val)
	}
}
