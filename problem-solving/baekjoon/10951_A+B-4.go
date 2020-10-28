package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for true {
		val, _ := fmt.Fscanln(reader, &a, &b)
		if val != 2 {
			break
		}
		fmt.Fprintln(writer, a+b)
	}
}
