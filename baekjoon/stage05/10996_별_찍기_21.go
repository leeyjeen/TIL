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

	if n == 1 {
		fmt.Println("*")
		return
	}
	for i:=0; i<2*n; i++ {
		for j:=0; j<n; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}