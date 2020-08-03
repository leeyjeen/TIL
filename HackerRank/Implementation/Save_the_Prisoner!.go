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

	var t int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		var n, m, s int
		fmt.Fscanln(reader, &n, &m, &s)
		fmt.Fprintln(writer, saveThePrisoner(n, m, s))
	}
}

func saveThePrisoner(n, m, s int) int {
	var answer = ((s + m - 1) % n)
	if answer == 0 {
		answer = n
	}
	return answer
}
