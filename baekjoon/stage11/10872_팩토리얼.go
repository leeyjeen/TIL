package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	fmt.Println(factorial(n))
}

func factorial(n int) (result int) {
	if n == 1 || n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
