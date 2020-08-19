package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var arr = make([]int, 5)
	var sum int
	for i := 0; i < 5; i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
		sum += arr[i]
	}
	sort.Ints(arr)
	fmt.Fprintln(writer, sum-arr[4], sum-arr[0])
}
