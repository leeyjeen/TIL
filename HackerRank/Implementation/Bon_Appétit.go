package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k, b int
	fmt.Fscanln(reader, &n, &k)

	var bill = make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Fscanf(reader, "%d ", &bill[i])
	}
	fmt.Fscanln(reader, &b)

	var sum, billActual int
	for i:=0; i<n; i++ {
		if i == k {
			continue
		}
		sum += bill[i]
	}
	billActual = sum/2
	if b- billActual >0 {
		fmt.Fprintln(writer, b-billActual)
	} else {
		fmt.Fprintln(writer, "Bon Appetit")
	}

}