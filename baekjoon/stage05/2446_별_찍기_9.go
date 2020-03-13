package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)

	for i:=0; i<n; i++ {
		for j:=0; j<i; j++ {
			fmt.Print(" ")
		}
		for j:=0; j<2*(n-i)-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i:=0; i <n-1; i++ {
		for j:=i; j<n-2; j++ {
			fmt.Print(" ")
		}
		for j:=0; j<2*i+3; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}