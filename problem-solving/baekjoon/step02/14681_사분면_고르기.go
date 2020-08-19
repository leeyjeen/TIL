package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x, y int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &x)
	fmt.Fscanln(reader, &y)

	if x > 0 {
		if y > 0 {
			fmt.Println(1)
		} else {
			fmt.Println(4)
		}
	} else {
		if y > 0 {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	}
}
