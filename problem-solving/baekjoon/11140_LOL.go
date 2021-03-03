// https://www.acmicpc.net/problem/11140
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		var input string
		fmt.Fscanln(reader, &input)
		fmt.Fprintln(writer, solve(input))
	}
}

func solve(input string) int {
	match, _ := regexp.MatchString("l([a-z])l", input)
	if strings.Contains(input, "lol") {
		return 0
	} else if strings.Contains(input, "lo") || strings.Contains(input, "ol") || strings.Contains(input, "ll") || match {
		return 1
	} else if strings.Contains(input, "l") || strings.Contains(input, "o") {
		return 2
	}
	return 3
}
