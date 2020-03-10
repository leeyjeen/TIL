package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &n)

	for i:=0; i<n; i++ {
		for j:=0; j<i+1; j++ {
			fmt.Fprint(writer,"*")
		}
		fmt.Fprint(writer,"\n")
	}
	writer.Flush()
}