package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var t int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &t)
	defer writer.Flush()

	for i := 0; i < t; i++ {
		var x1, y1, r1, x2, y2, r2 int
		fmt.Fscanln(reader, &x1, &y1, &r1, &x2, &y2, &r2)

		var distanceX = x1 - x2
		var distanceY = y1 - y2

		var addR = math.Pow(float64(r1+r2), 2)
		var subR = math.Pow(float64(r1-r2), 2)

		var d = math.Pow(float64(distanceX), 2) + math.Pow(float64(distanceY), 2)

		if d < addR && d > subR {
			fmt.Fprintln(writer, 2)
		} else if d == addR || d == subR && d != 0 {
			fmt.Fprintln(writer, 1)
		} else if d < subR || d > addR {
			fmt.Fprintln(writer, 0)
		} else if d == 0 && r1 == r2 {
			fmt.Fprintln(writer, -1)
		}
	}
}
