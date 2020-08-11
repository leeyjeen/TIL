package main

import "fmt"

func Recursive(num int) {
	if num <= 0 { // escape condition
		return // escape
	}
	fmt.Printf("Recursive call! %d \n", num)
	Recursive(num-1)
}

func main() {
	Recursive(3)
}