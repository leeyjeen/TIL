package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &t)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var h, w, n int
	for i:=0; i<t; i++ {
		fmt.Fscanln(reader, &h, &w, &n)
		var roomNumber, stage int
		if n%h == 0 {
			roomNumber = n/h
			stage = h
		} else {
			roomNumber = n/h + 1
			stage = n%h
		}
		fmt.Fprintln(writer, stage * 100  + roomNumber)
	}
}