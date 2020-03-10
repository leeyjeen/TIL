package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b = 9, 9
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for true {
		fmt.Fscanln(reader, &a, &b)
		if a == 0 || b == 0 {
			break
		}
		fmt.Fprintln(writer, a+b)
	}
	writer.Flush()
}
