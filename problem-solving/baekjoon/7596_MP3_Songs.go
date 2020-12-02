// https://www.acmicpc.net/problem/7596
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

	var count int = 1
	for {
		var n int
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}
		var playlist []string
		for i := 0; i < n; i++ {
			var input string
			input, _ = reader.ReadString('\n') // 공백 포함하여 입력 받기 위해 ReadString() 사용
			playlist = append(playlist, strings.TrimRight(input, "\n"))
		}
		sort.Strings(playlist)
		fmt.Fprintln(writer, count)
		for _, v := range playlist {
			fmt.Fprintln(writer, v)
		}
		count++
	}
}
