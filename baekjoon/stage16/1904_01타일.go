package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	sequence := countBinarySequence(n)
	fmt.Println(sequence[n-1])
}

func countBinarySequence(n int) (sequence []int) {
	sequence = append(sequence, 1, 2)
	for i := 3; i<n+1; i++ {
		sequence = append(sequence, (sequence[i-2]+sequence[i-3])%15746)
	}
	return
}