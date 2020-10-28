package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var r int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &r)

	fmt.Printf("%.6f\n", float64(r*r)*math.Pi)
	fmt.Printf("%.6f\n", float64(2*r*r))
}
