// https://www.acmicpc.net/problem/2667
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	graph       [][]string
	visited     [][]bool
	numOfHouses int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	graph = make([][]string, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]string, n)
	}
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		var input string
		input, _ = reader.ReadString('\n') // 공백 포함하여 입력 받기 위해 ReadString() 사용
		convInput := strings.ReplaceAll(input, "\n", "")
		inputs := strings.Split(convInput, "")
		for j, input := range inputs {
			graph[i][j] = input
		}
	}
	numbers := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == "1" && !visited[i][j] {
				numOfHouses = 0
				dfs(i, j)
				numbers = append(numbers, numOfHouses)
			}
		}
	}

	sort.Ints(numbers)
	fmt.Fprintln(writer, len(numbers))
	for _, num := range numbers {
		fmt.Fprintln(writer, num)
	}
}

func dfs(row, col int) {
	visited[row][col] = true
	numOfHouses++

	if row+1 < len(graph) && graph[row+1][col] == "1" && !visited[row+1][col] {
		dfs(row+1, col)
	}
	if row-1 >= 0 && graph[row-1][col] == "1" && !visited[row-1][col] {
		dfs(row-1, col)
	}
	if col+1 < len(graph[row]) && graph[row][col+1] == "1" && !visited[row][col+1] {
		dfs(row, col+1)
	}
	if col-1 >= 0 && graph[row][col-1] == "1" && !visited[row][col-1] {
		dfs(row, col-1)
	}
}
