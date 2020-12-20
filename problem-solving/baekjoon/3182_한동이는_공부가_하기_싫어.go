// https://www.acmicpc.net/problem/3182
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph   []int
	visited []bool
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	graph = make([]int, n+1)
	visited = make([]bool, n+1)
	for i := 0; i < n; i++ {
		var response int
		fmt.Fscanln(reader, &response)
		graph[i+1] = response
	}

	var maxContacts int
	var whoToContact = n + 1
	for i := 1; i < n+1; i++ {
		var numOfContacts int
		if graph[i] != 0 {
			visited[i] = true
			numOfContacts = dfs(graph[i], numOfContacts)
		}
		if numOfContacts > maxContacts {
			maxContacts = numOfContacts
			whoToContact = i
		} else if numOfContacts == maxContacts {
			if whoToContact > i {
				whoToContact = i
			}
		}
		initVisited()
	}
	fmt.Fprintln(writer, whoToContact)
}

func dfs(start int, numOfContacts int) int {
	visited[start] = true
	numOfContacts++
	if graph[graph[start]] != 0 && !visited[graph[start]] {
		numOfContacts = dfs(graph[start], numOfContacts)
	}
	return numOfContacts
}

func initVisited() {
	for i := 0; i < len(visited); i++ {
		visited[i] = false
	}
}
