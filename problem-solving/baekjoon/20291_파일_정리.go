// https://www.acmicpc.net/problem/20291
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	extMap := map[string]int{}
	for i := 0; i < n; i++ {
		var filepath string
		fmt.Fscanln(reader, &filepath)
		ext := strings.Split(filepath, ".")[1]
		extMap[ext]++
	}

	sortedKeys := []string{}
	for key := range extMap {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		fmt.Fprintf(writer, "%s %d\n", key, extMap[key])
	}
}
