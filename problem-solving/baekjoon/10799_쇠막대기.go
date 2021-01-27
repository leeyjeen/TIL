// https://www.acmicpc.net/problem/10799
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

	var input string
	fmt.Fscanln(reader, &input)
	inputs := strings.Split(input, "")
	fmt.Fprintln(writer, getBarCount(inputs))
}

func getBarCount(inputs []string) int {
	var stack []string
	var prev string
	var barCount int
	for i := 0; i < len(inputs); i++ {
		cur := inputs[i]
		if cur == "(" {
			stack = append(stack, cur)
		} else if cur == ")" {
			stack = stack[0 : len(stack)-1]
			if prev == "(" { // "(" 다음 ")"가 바로 들어오면 레이저
				barCount += len(stack) // 레이저인 경우 스택 안의 괄호 갯수만큼 막대 카운트.
			} else {
				barCount++ // ")" 다음 ")"가 들어온 경우 막대기 끝의 개수 하나 카운트.
			}
		}
		prev = cur
	}
	return barCount
}
