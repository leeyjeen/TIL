// https://www.acmicpc.net/problem/2960
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

	var n, k int
	fmt.Fscanln(reader, &n, &k)
	fmt.Fprintln(writer, eratos(n, k))
}

func eratos(n, k int) (kth int) {
	primeNumbers := initPrimeNumbers(n)
	for i := 2; i < n+1; i++ {
		if primeNumbers[i] {
			for j := i; j < n+1; j += i {
				if primeNumbers[j] {
					primeNumbers[j] = false
					k--
					if k == 0 {
						kth = j
						return
					}
				}
			}
		}
	}
	return
}

func initPrimeNumbers(n int) []bool {
	var primeNumbers = make([]bool, n+1)
	for i := range primeNumbers {
		primeNumbers[i] = true
	}
	return primeNumbers
}
