// https://www.acmicpc.net/problem/7785
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

	var workers = map[string]int{}
	for i := 0; i < n; i++ {
		var name, commute string
		fmt.Fscanln(reader, &name, &commute)
		if commute == "enter" {
			workers[name] = 1
		} else if commute == "leave" {
			delete(workers, name)
		}
	}

	var keys = []string{}
	for key, _ := range workers {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	for _, key := range keys {
		fmt.Fprintln(writer, key)
	}
}
