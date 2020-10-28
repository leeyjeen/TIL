package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k int
	fmt.Fscanln(reader, &k)

	var s = NewStack()
	s.Init()
	for i := 0; i < k; i++ {
		var num int
		fmt.Fscanln(reader, &num)
		if num == 0 {
			s.Pop()
			continue
		}
		s.Push(num)
	}

	var sum int
	for true {
		val := s.Pop()
		if val == -1 {
			break
		}
		sum += val
	}
	fmt.Println(sum)
}

type Stack interface {
	GetSize() int
	IsEmpty() bool
	Push(int)
	Pop() int
	GetTop() int
	Init()
}

type stack struct {
	arr [100000]int
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

func (s *stack) Push(n int) {
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
