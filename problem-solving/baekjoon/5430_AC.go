// https://www.acmicpc.net/problem/5430
// 덱을 활용하여 시간복잡도를 향상시키는 문제
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

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var command string
		fmt.Fscanln(reader, &command)
		var n int
		fmt.Fscanln(reader, &n)
		var array string
		fmt.Fscanln(reader, &array)

		array = strings.ReplaceAll(array, "[", "")
		array = strings.ReplaceAll(array, "]", "")
		arr := strings.Split(array, ",")
		if n == 0 {
			arr = []string{}
		}

		var isAsc = true
		var isError = false
		for j := 0; j < len(command); j++ {
			if string(command[j]) == "R" {
				isAsc = !isAsc
			} else if string(command[j]) == "D" {
				if len(arr) == 0 {
					isError = true
					break
				}
				if isAsc {
					arr = arr[1:]
				} else {
					arr = arr[:len(arr)-1]
				}
			}
		}
		if isError {
			fmt.Fprintln(writer, "error")
		} else {
			fmt.Fprint(writer, "[")
			if isAsc {
				for k := 0; k < len(arr); k++ {
					if k == 0 {
						fmt.Fprint(writer, arr[k])
					} else {
						fmt.Fprint(writer, ","+arr[k])
					}
				}
			} else {
				for k := len(arr) - 1; k >= 0; k-- {
					if k == len(arr)-1 {
						fmt.Fprint(writer, arr[k])
					} else {
						fmt.Fprint(writer, ","+arr[k])
					}
				}
			}
			fmt.Fprintln(writer, "]")
		}
	}
}
