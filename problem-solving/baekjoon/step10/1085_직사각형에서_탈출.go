package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var x, y, w, h float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &x, &y, &w,&h)

	fmt.Println(math.Min(math.Min(x, w-x), math.Min(y, h-y)))
}