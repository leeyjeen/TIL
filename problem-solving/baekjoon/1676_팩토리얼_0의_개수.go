// https://www.acmicpc.net/problem/1676
// 소인수분해의 성질을 활용하여 N!의 끝에 0이 얼마나 많이 오는지 구하는 문제
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

	// 팩토리얼로 풀면 오버플로우 발생하여 값이 이상해짐 -> 소인수분해 활용하여 풀어야 한다
	// 끝에 0이 오는 개수는 소인수분해 했을 때 2*5 쌍의 개수와 같다
	// 소인수분해 했을 때 2의 개수가 5의 개수보다 훨씬 많으므로, 5의 개수만 카운팅하면 된다
	var fiveCount int

	for i := 1; i <= n; i++ {
		var temp = i
		for temp%5 == 0 && temp > 0 {
			fiveCount++
			temp /= 5
		}
	}

	fmt.Fprintln(writer, fiveCount)
}
