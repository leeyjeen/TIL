package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t, x, y int

	fmt.Fscanln(reader, &t)

	for i := 0; i < t ; i ++ {
		fmt.Fscanln(reader, &x, &y)
		var temp = 1
		for true {
			if y - x  <= temp * temp {
				fmt.Fprintln(writer, temp*2-1)
				break
			} else if y - x <= temp * temp + temp {
				fmt.Fprintln(writer, temp*2)
				break
			}
			temp++
		}
	}
}