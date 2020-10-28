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

	var currentNum = 1
	var room = 1
	for true {
		if currentNum >= n {
			break
		}
		currentNum += 6*room
		room++
	}

	fmt.Println(room)
}