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
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
	}
	for _, v := range cutTheSticks(arr) {
		fmt.Fprintln(writer, v)
	}
}

func cutTheSticks(arr []int) []int {
	var minStick int = 1000
	var count int = 0
	var result = []int{}
	for {
		for _, v := range arr {
			if v == 0 {
				continue
			}
			if v < minStick {
				minStick = v
			}
		}
		for i, v := range arr {
			if v == 0 {
				continue
			}
			arr[i] -= minStick
			count++
		}
		if count == 0 {
			break
		}
		result = append(result, count)
		count = 0
		minStick = 1000
	}
	return result
}
