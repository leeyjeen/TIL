package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords) // for fast scan
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text()) 

	var counts = make([]int, 10001)
	for i:=0; i<n; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		counts[num]++	// 카운팅정렬 
	}

	for i:=1; i<10001; i++ {
		for j:=0; j<counts[i]; j++ {
			fmt.Fprintln(writer, i)
		}
	}
}