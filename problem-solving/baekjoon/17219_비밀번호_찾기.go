// https://www.acmicpc.net/problem/17219
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var passwordMap = map[string]string{}
	for i := 0; i < n; i++ {
		var input string
		input, _ = reader.ReadString('\n')
		input = strings.ReplaceAll(input, "\n", "")
		inputs := strings.Split(input, " ")
		passwordMap[inputs[0]] = inputs[1]
	}

	for i := 0; i < m; i++ {
		var site string
		fmt.Fscanln(reader, &site)
		fmt.Fprintln(writer, passwordMap[site])
	}
}
