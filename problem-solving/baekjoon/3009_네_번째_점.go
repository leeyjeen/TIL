package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x = make([]int, 4)
	var y = make([]int, 4)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 3; i++ {
		fmt.Fscanf(reader, "%d %d\n", &x[i], &y[i])
	}

	if x[0] == x[1] {
		x[3] = x[2]
	} else if x[0] == x[2] {
		x[3] = x[1]
	} else if x[1] == x[2] {
		x[3] = x[0]
	}

	if y[0] == y[1] {
		y[3] = y[2]
	} else if y[0] == y[2] {
		y[3] = y[1]
	} else if y[1] == y[2] {
		y[3] = y[0]
	}

	fmt.Printf("%d %d\n", x[3], y[3])
}
