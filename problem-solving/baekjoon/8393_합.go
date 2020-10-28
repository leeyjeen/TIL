package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var sum int
	for i:=1; i<n+1; i++ {
		sum += i
	}

	fmt.Println(sum)
}