package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var word string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &word)

	var seconds = []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 10, 10, 10, 10}
	var result int
	for i := 0; i < len(word); i++ {
		result += seconds[int(word[i])-65]
	}
	fmt.Println(result)
}
