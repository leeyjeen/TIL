package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var h, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &h, &m)

	if m-45 >= 0 {
		fmt.Printf("%d %d\n", h, m-45)
	} else if h == 0 {
		fmt.Printf("%d %d\n", 23, 15+m)
	} else {
		fmt.Printf("%d %d\n", h-1, 15+m)
	}
}
