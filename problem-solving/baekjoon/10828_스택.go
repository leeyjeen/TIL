package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	var s = NewStack()
	s.Init()
	for i := 0; i < n; i++ {
		var command string
		command, _ = reader.ReadString('\n')
		command = strings.ReplaceAll(command, "\n", "")
		commands := strings.Split(command, " ")
		if commands[0] == "push" {
			num, _ := strconv.Atoi(commands[1])
			s.Push(num)
		} else if commands[0] == "pop" {
			fmt.Fprintln(writer, s.Pop())
		} else if commands[0] == "size" {
			fmt.Fprintln(writer, s.GetSize())
		} else if commands[0] == "empty" {
			if s.IsEmpty() {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		} else if commands[0] == "top" {
			fmt.Fprintln(writer, s.GetTop())
		}
	}
}

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(int)
	Pop() int
	GetTop() int
	Init()
	IsFull() bool
}

type stack struct {
	arr [10000]int
	top int
}

func NewStack() Stack {
	return &stack{}
}

func (s *stack) Init() {
	s.top = -1
}

func (s *stack) GetSize() int {
	return s.top + 1
}
func (s *stack) IsEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *stack) IsFull() bool {
	if s.GetSize() >= 10000 {
		return true
	}
	return false
}

func (s *stack) Push(n int) {
	if s.IsFull() {
		return
	}
	s.top++
	s.arr[s.top] = n
}

func (s *stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	result := s.arr[s.top]
	s.top--
	return result
}

func (s *stack) GetTop() int {
	if s.IsEmpty() {
		return -1
	}
	return s.arr[s.top]
}
