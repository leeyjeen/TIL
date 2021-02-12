// https://www.acmicpc.net/problem/3048
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n1, n2 int
	fmt.Fscanln(reader, &n1, &n2)
	first, _ := reader.ReadString('\n')
	first = strings.ReplaceAll(first, "\n", "")
	second, _ := reader.ReadString('\n')
	second = strings.ReplaceAll(second, "\n", "")
	var t int
	fmt.Fscanln(reader, &t)
	fmt.Fprintln(writer, ants(first, second, t))
}

func ants(first, second string, t int) string {
	first = reverseString(first)
	r := []rune(first)
	r2 := []rune(second)
	for i := range r2 {
		r2[i] = -r2[i]
	}
	r = append(r, r2...)
	for i := 0; i < t; i++ {
		swapped := false
		for j := 0; j < len(r)-1; j++ {
			if r[j] > 0 && r[j+1] < 0 {
				if !swapped {
					r[j], r[j+1] = r[j+1], r[j]
					swapped = true
				}
			} else {
				swapped = false
			}
		}
	}

	for i := 0; i < len(r); i++ {
		if r[i] < 0 {
			r[i] = -r[i]
		}
	}

	return string(r)
}

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
