// https://www.acmicpc.net/problem/1475
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

	var n int
	fmt.Fscanln(reader, &n)
	if n == 0 {
		fmt.Fprintln(writer, 1)
		return
	}
	numberCount := map[int]int{}
	for n > 0 {
		numberCount[n%10]++
		n /= 10
	}
	sixCount, nineCount := numberCount[6], numberCount[9]
	if (sixCount+nineCount)%2 == 0 {
		numberCount[6] = (sixCount + nineCount) / 2
		numberCount[9] = numberCount[6]
	} else {
		numberCount[6] = (sixCount + nineCount) / 2
		numberCount[9] = numberCount[6] + 1
	}
	var setCount int
	for _, v := range numberCount {
		if setCount < v {
			setCount = v
		}
	}
	fmt.Fprintln(writer, setCount)
}
