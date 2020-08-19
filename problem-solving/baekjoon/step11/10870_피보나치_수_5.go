package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var result = fibonacci(n)
	fmt.Println(result)
}

func fibonacci(n int) (result int) {
	if n == 0 || n == 1 {
		result = n
		return result
	}
	return fibonacci(n-1) + fibonacci(n-2)
}