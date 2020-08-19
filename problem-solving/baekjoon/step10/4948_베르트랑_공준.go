package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n, count int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for true {
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}
		var numbers = make(map[int]bool, n)
		if n == 1 {
			fmt.Fprintln(writer, 1)
			continue
		}
		for i := n; i < 2*n; i++ {
			numbers[i+1] = true
		}

		for i := n; i < 2*n; i++ {
			if numbers[i+1] == false {
				break
			}
			for j := 1; j < int(math.Sqrt(float64(i+1))); j++ {
				if (i+1)%(j+1) == 0 {
					for k := 1; k*i <= 2*n; k++ {
						numbers[k*(i+1)] = false
					}
					break
				}
			}
		}

		count = 0
		for _, v := range numbers {
			if v {
				count++
			}
		}
		fmt.Fprintln(writer, count)
	}
}
