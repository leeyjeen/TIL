// https://www.acmicpc.net/problem/10866
// 덱의 개념을 익히고 실습하는 문제
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

	var deque []int
	for i := 0; i < n; i++ {
		var command string
		var num int
		fmt.Fscanln(reader, &command, &num)

		if command == "push_back" {
			deque = pushBack(deque, num)
		} else if command == "push_front" {
			deque = pushFront(deque, num)
		} else if command == "pop_front" {
			var front int
			deque, front = popFront(deque)
			fmt.Fprintln(writer, front)
		} else if command == "pop_back" {
			var back int
			deque, back = popBack(deque)
			fmt.Fprintln(writer, back)
		} else if command == "size" {
			fmt.Fprintln(writer, size(deque))
		} else if command == "empty" {
			if isEmpty(deque) {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		} else if command == "front" {
			fmt.Fprintln(writer, front(deque))
		} else if command == "back" {
			fmt.Fprintln(writer, back(deque))
		}
	}
}

func pushBack(deque []int, num int) []int {
	deque = append(deque, num)
	return deque
}

func pushFront(deque []int, num int) []int {
	deque = append([]int{num}, deque...)
	return deque
}

func popFront(deque []int) ([]int, int) {
	if isEmpty(deque) {
		return deque, -1
	}
	front := deque[0]
	deque = deque[1:]
	return deque, front
}

func popBack(deque []int) ([]int, int) {
	if isEmpty(deque) {
		return deque, -1
	}
	back := deque[len(deque)-1]
	deque = deque[:len(deque)-1]
	return deque, back
}

func size(deque []int) int {
	return len(deque)
}

func isEmpty(deque []int) bool {
	if size(deque) == 0 {
		return true
	}
	return false
}

func front(deque []int) int {
	if isEmpty(deque) {
		return -1
	}
	return deque[0]
}

func back(deque []int) int {
	if isEmpty(deque) {
		return -1
	}
	return deque[len(deque)-1]
}
