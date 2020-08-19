package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var numbers = make([]int, 10)
	var mods []int
	for i := range numbers {
		fmt.Fscanf(reader, "%d\n", &numbers[i])
		var mod = numbers[i] % 42
		if !contains(mods, mod) {
			mods = append(mods, mod)
		}
	}

	fmt.Println(len(mods))
}

func contains(numbers []int, number int) bool {
	for _, n := range numbers {
		if number == n {
			return true
		}
	}
	return false
}
