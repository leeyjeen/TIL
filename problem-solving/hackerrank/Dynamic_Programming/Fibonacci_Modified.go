package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t1, t2 big.Int // use big package to prevent integer overflow
	var n int
	fmt.Fscanln(reader, &t1, &t2, &n)
	fmt.Fprintln(writer, fibonacci(t1, t2, n))
}

func fibonacci(t1, t2 big.Int, n int) string {
	var f = make([]big.Int, n)
	f[0], f[1] = t1, t2

	var multiplied = new(big.Int)
	for i := 2; i < int(n); i++ {
		f[i].Add(&f[i-2], multiplied.Mul(&f[i-1], &f[i-1]))
	}
	return f[n-1].String()
}
