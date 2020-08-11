package main

import "fmt"

func Fibo(n int) int {
	fmt.Printf("func call param %d \n", n)

	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	return Fibo(n-1) + Fibo(n-2)
}

func main() {
	Fibo(7)
}
