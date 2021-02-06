// https://www.acmicpc.net/problem/10546
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

	var runners = map[string]int{}
	for i := 0; i < n; i++ {
		var runner string
		fmt.Fscanln(reader, &runner)
		runners[runner]++
	}
	for i := 0; i < n-1; i++ {
		var finished string
		fmt.Fscanln(reader, &finished)
		runners[finished]--
		if runners[finished] == 0 {
			delete(runners, finished)
		}
	}
	for key := range runners {
		fmt.Fprintln(writer, key)
	}
}
