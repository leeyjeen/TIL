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

	var h = make([]int, 26)
	for i := 0; i < 26; i++ {
		fmt.Fscanf(reader, "%d ", &h[i])
	}
	var word string
	fmt.Fscanln(reader, &word)
	fmt.Fprintln(writer, designerPdfViewer(h, word))
}

func designerPdfViewer(h []int, word string) int {
	var max int
	for _, v := range word {
		if h[v-97] > max {
			max = h[v-97]
		}
	}
	return max * len(word)
}
