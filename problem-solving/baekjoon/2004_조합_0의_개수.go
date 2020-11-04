// https://www.acmicpc.net/problem/2004
// nCm의 끝에 0이 얼마나 많이 오는지 구하는 문제
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

	var n, m, count, twoCount, fiveCount int
	fmt.Fscanln(reader, &n, &m)

	twoCount = getTwoCount(n) - getTwoCount(n-m) - getTwoCount(m)
	fiveCount = getFiveCount(n) - getFiveCount(n-m) - getFiveCount(m)

	if twoCount > fiveCount {
		count = fiveCount
	} else {
		count = twoCount
	}
	fmt.Fprintln(writer, count)
}

func getTwoCount(num int) (twoCount int) {
	for i := 2; i <= num; i *= 2 {
		twoCount += num / i
	}
	return twoCount
}

func getFiveCount(num int) (fiveCount int) {
	for i := 5; i <= num; i *= 5 {
		fiveCount += num / i
	}
	return fiveCount
}

/* 시간초과..
// reader := bufio.NewReader(os.Stdin)
// writer := bufio.NewWriter(os.Stdout)
// defer writer.Flush()

// var n, m, count, twoCount, fiveCount int
// fmt.Fscanln(reader, &n, &m)
// for i := n - m + 1; i <= n; i++ {
// 	var temp = i
// 	for temp%2 == 0 && temp > 0 {
// 		twoCount++
// 		temp /= 2
// 	}
// 	temp = i
// 	for temp%5 == 0 && temp > 0 {
// 		fiveCount++
// 		temp /= 5
// 	}
// }
// for i := 1; i <= m; i++ {
// 	var temp = i
// 	for temp%2 == 0 && temp > 0 {
// 		twoCount--
// 		temp /= 2
// 	}
// 	temp = i
// 	for temp%5 == 0 && temp > 0 {
// 		fiveCount--
// 		temp /= 5
// 	}
// }
// if twoCount > fiveCount {
// 	count = fiveCount
// } else {
// 	count = twoCount
// }
// fmt.Fprintln(writer, count)
*/
