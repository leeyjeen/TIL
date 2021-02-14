// https://www.acmicpc.net/problem/12789
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	var students = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &students[i])
	}
	fmt.Fprintln(writer, getResult(students))
}

func getResult(students []int) string {
	var stack []int
	var line []int
	for len(students) > 0 {
		if len(stack) == 0 {
			stack = append(stack, students[0])
			students = students[1:]
		} else {
			if stack[len(stack)-1] > students[0] {
				stack = append(stack, students[0])
				students = students[1:]
			} else {
				line = append(line, stack[len(stack)-1])
				students = students[1:]
			}
		}
	}
	for len(stack) > 0 {
		line = append(line, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	if sort.SliceIsSorted(line, func(i, j int) bool {
		return line[i] < line[j]
	}) {
		return "Nice"
	}
	return "Sad"
}
