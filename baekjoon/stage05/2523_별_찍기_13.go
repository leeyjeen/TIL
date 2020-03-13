package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n)

	for i := 0; i < 2*n-1; i++ {
		if i < n {
			for j := 0; j < i+1; j++ {
				fmt.Print("*")
			}
		} else {
			for j := i; j < 2*n-1; j++ {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
