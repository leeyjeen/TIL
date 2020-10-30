// https://www.acmicpc.net/problem/3036
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
	var inputs = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &inputs[i])
		if i == 0 {
			continue
		}
		gcm := gcm(inputs[0], inputs[i])
		fmt.Fprintf(writer, "%d/%d\n", inputs[0]/gcm, inputs[i]/gcm)
	}
}

func gcm(a, b int) int {
	if b == 0 {
		return a
	}
	a, b = b, a%b
	return gcm(a, b)
}
