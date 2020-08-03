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

	var n, p int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &p)

	fmt.Fprintln(writer, pageCount(n, p))
}

func pageCount(n, p int) int {
	return int(math.Min(float64(n/2-p/2), float64(p/2)))
}
