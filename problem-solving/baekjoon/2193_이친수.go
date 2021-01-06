// https://www.acmicpc.net/problem/2193
/*
N = 1일때 count = 1
1

N = 2일때 count = 1
10

N = 3일때 count = 2
100
101

N = 4일때 count = 3
1000
1010
1001

N = 5일때 count = 5
10000
10100
10101
10010
10001

N = 6일때 count = 8
100000
101000
101010
101001
100100
100101
100010
100001

...
*/
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

	var count = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if i < 2 {
			count[i] = 1
		} else {
			count[i] = count[i-1] + count[i-2]
		}
	}
	fmt.Fprintln(writer, count[n])
}
