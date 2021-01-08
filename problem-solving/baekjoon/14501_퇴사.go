// https://www.acmicpc.net/problem/14501
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	result int
	infos  []Info
	n      int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var time, profit int
		fmt.Fscanln(reader, &time, &profit)
		infos = append(infos, Info{time, profit})
	}
	getMax(0, 0)
	fmt.Fprintln(writer, result)
}

type Info struct {
	time   int
	profit int
}

func getMax(index, profit int) {
	if index == n {
		if result < profit {
			result = profit
		}
		return
	}

	if index > n {
		return
	}

	getMax(index+infos[index].time, profit+infos[index].profit) // 해당 날짜에 상담을 하는 경우
	getMax(index+1, profit)                                     // 해당 날짜에 상담을 하지 않는 경우
}
