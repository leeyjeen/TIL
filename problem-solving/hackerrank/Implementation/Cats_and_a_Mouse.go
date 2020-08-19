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

	var q int
	fmt.Fscanln(reader, &q)
	for i := 0; i < q; i++ {
		var x, y, z int
		fmt.Fscanln(reader, &x, &y, &z)
		fmt.Fprintln(writer, catAndMouse(x, y, z))
	}
}

func catAndMouse(x, y, z int) string {
	xToZ := z - x
	if xToZ < 0 {
		xToZ = -xToZ
	}
	yToZ := z - y
	if yToZ < 0 {
		yToZ = -yToZ
	}
	if xToZ > yToZ {
		return "Cat B"
	} else if yToZ > xToZ {
		return "Cat A"
	}
	return "Mouse C"
}
