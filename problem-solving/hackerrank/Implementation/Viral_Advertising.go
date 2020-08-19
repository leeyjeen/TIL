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
	fmt.Fprintln(writer, viralAdvertising(n))
}

func viralAdvertising(n int) int {
	var shared = 5
	var totalLiked int
	for i := 1; i <= n; i++ {
		var liked = shared / 2
		totalLiked += liked
		shared = liked * 3
	}
	return totalLiked
}
