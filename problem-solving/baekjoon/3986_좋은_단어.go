// https://www.acmicpc.net/problem/3986
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

	var n int
	fmt.Fscanln(reader, &n)
	var count int
	for i := 0; i < n; i++ {
		var input string
		input, _ = reader.ReadString('\n') // 공백 포함하여 입력 받기 위해 ReadString() 사용
		inputs := strings.Split(strings.ReplaceAll(input, "\n", ""), "")
		stack := []string{inputs[0]}
		for j := 1; j < len(inputs); j++ {
			if len(stack) == 0 {
				stack = []string{inputs[j]}
			} else if stack[len(stack)-1] == inputs[j] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, inputs[j])
			}
		}
		if len(stack) == 0 {
			count++
		}
	}
	fmt.Fprintln(writer, count)
}
