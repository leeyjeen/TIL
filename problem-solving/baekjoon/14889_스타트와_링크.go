package main

import (
	"fmt"
	"os"
	"bufio"
	"math"
)

var (
	n int
	result = 1000
	numbers [][]int
	visited = make([]bool, 21)
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)

	numbers = make([][]int, n)
	for i := range numbers {
		numbers[i] = make([]int, n)
	}

	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			fmt.Fscanf(reader, "%d ", &numbers[i][j])
		}
	}
	dfs(0, 0)
	fmt.Fprintln(writer, result)
}

func dfs(index, len int) {
	if n/2 == len {
		computeResult()
	} else {
		for i := index + 1 ; i < n; i++ {
			if !visited[i] {
				visited[i] = true
				dfs(i, len + 1)
			}
		} 
	}
	visited[index] = false
}

func computeResult() {
	start := make([]int, n/2)
	link := make([]int, n/2)
	var startIdx, linkIdx int
	for i:=0; i<n; i++ {
		if visited[i] {
			start[startIdx] = i
			startIdx++
		} else {
			link[linkIdx] = i
			linkIdx++
		}
	}

	var diff, startTotal, linkTotal int
	for i:=0; i<n/2; i++ {
		for j:=i+1; j<n/2; j++ {
			startTotal = startTotal + numbers[start[i]][start[j]] + numbers[start[j]][start[i]]
			linkTotal = linkTotal + numbers[link[i]][link[j]] + numbers[link[j]][link[i]]
		}
	}
	diff = int(math.Abs(float64(linkTotal-startTotal)))
	if diff < result {
		result = diff
	}
}