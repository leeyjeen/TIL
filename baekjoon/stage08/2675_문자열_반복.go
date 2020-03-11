package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, r int
	var str string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &t)
	defer writer.Flush()

	for i:=0; i<t; i++ {
		fmt.Fscanf(reader, "%d %s\n", &r, &str)
		for j:=0; j<len(str); j++ {
			for k:=0; k<r; k++ {
				fmt.Fprint(writer, string(str[j]))
			}
		}
		fmt.Fprint(writer, "\n")
	}
}