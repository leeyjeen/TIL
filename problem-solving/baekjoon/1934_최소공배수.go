// https://www.acmicpc.net/problem/1934
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

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintln(writer, getLCM(a, b))
	}

}

// 최소 공배수 구하기 (유클리드 알고리즘 이용)
func getLCM(first, second int) (lcm int) {
	return first * second / getGCD(first, second) // 두 수의 곱 나누기 최대공약수
}

// 최대 공약수 구하기 (유클리드 알고리즘 이용)
func getGCD(first, second int) (gcd int) {
	if first < second { // fn에 큰 값을 오게 하기
		second, first = first, second
	}

	for second != 0 { // second가 0이 될때까지 반복
		first, second = second, first%second
	}
	return first
}
