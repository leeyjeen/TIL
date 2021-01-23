// https://www.acmicpc.net/problem/1620
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

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var dogam = map[string]string{}
	for i := 0; i < n; i++ {
		var name string
		fmt.Fscanln(reader, &name)
		dogam[name] = fmt.Sprintf("%d", i+1)
		dogam[fmt.Sprintf("%d", i+1)] = name
	}
	for i := 0; i < m; i++ {
		var input string
		fmt.Fscanln(reader, &input)
		fmt.Fprintln(writer, dogam[input])
	}
}
