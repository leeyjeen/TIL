package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var numbers = make([]int, 9)
	reader := bufio.NewReader(os.Stdin)
	var max = 0
	var maxIndex = 0
	for i := range numbers {
		fmt.Fscanf(reader, "%d\n", &numbers[i])
		if numbers[i] > max {
			max = numbers[i]
			maxIndex = i
		}
	}

	fmt.Println(max)
	fmt.Println(maxIndex+1)
}