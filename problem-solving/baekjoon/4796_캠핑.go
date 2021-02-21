// https://www.acmicpc.net/problem/4796
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 1; ; i++ {
		var l, p, v int
		fmt.Fscanln(reader, &l, &p, &v)
		if l == 0 && p == 0 && v == 0 {
			break
		}

		fmt.Fprintf(writer, "Case %d: %d\n", i, (v/p)*l+int(math.Min(float64(v%p), float64(l))))
	}

}
