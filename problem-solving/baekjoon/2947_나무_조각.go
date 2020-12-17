// https://www.acmicpc.net/problem/2947
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

	var arr = make([]int, 5)
	fmt.Fscanln(reader, &arr[0], &arr[1], &arr[2], &arr[3], &arr[4])

	for arr[0] != 1 || arr[1] != 2 || arr[2] != 3 || arr[3] != 4 || arr[4] != 5 {
		for i := 0; i < 4; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				for j := 0; j < 4; j++ {
					fmt.Fprintf(writer, "%d ", arr[j])
				}
				fmt.Fprintf(writer, "%d\n", arr[4])
			}
		}
	}
}
