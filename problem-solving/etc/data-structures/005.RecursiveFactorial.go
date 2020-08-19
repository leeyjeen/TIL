package main

import "fmt"

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

func main() {
	fmt.Printf("1! = %d \n", Factorial(1))
	fmt.Printf("2! = %d \n", Factorial(2))
	fmt.Printf("3! = %d \n", Factorial(3))
	fmt.Printf("4! = %d \n", Factorial(4))
	fmt.Printf("9! = %d \n", Factorial(9))
}
