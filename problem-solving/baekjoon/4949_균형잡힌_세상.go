// https://www.acmicpc.net/problem/4949
// 주어진 문자열이 올바른 괄호열인지 판단하는 문제2
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

	for {
		var input string
		input, _ = reader.ReadString('\n') // 공백 포함하여 입력 받기 위해 ReadString() 사용
		convInput := strings.ReplaceAll(input, "\n", "")
		if convInput == "." {
			break
		}
		fmt.Fprintln(writer, isBalanced(convInput))
	}
}

func isBalanced(input string) (result string) {
	var stack []string
	for _, v := range input {
		strV := string(v)
		if strV == "(" || strV == "[" {
			stack = append(stack, strV)
		} else if strV == ")" {
			if len(stack) > 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			} else {
				return "no"
			}
		} else if strV == "]" {
			if len(stack) > 0 && stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
			} else {
				return "no"
			}
		} else {
			continue
		}
	}
	if len(stack) > 0 {
		return "no"
	}
	return "yes"
}


