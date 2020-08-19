package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	
	var fibMap = make(map[int]int, n)
	fibMap[0] = 0
	fibMap[1] = 1
	fmt.Fprintln(writer, fibonacci(fibMap, n))
}

func fibonacci(fibMap map[int]int, n int) (result int) {
	for i:=2; i<=n; i++ {
		fibMap[i] = fibMap[i-1]+fibMap[i-2]
	}
	return fibMap[n]
}