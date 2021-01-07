// https://www.acmicpc.net/problem/14697v
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

	var a, b, c, n int
	fmt.Fscanln(reader, &a, &b, &c, &n)
	fmt.Fprintln(writer, checkRooms(a, b, c, n))
}

func checkRooms(a, b, c int, n int) int {
	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			for k := 0; k <= n-i-j; k++ {
				if a*i+b*j+c*k == n {
					return 1
				}
			}
		}
	}
	return 0
}
