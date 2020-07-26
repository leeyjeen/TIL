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
	var steps = make([]rune, n)
	var height, count int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%c", &steps[i])
		var input string = fmt.Sprintf("%c", steps[i])
		if input == "U" {
			height++
			if height == 0 {
				count++
			}
		} else if input == "D" {
			height--
		}
	}
	fmt.Fprintln(writer, count)
}
