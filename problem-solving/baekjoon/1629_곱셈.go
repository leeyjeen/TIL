// https://www.acmicpc.net/problem/1629
// 분할 정복으로 거듭제곱을 빠르게 계산하는 문제
package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c int64
	fmt.Fscanln(reader, &a, &b, &c)
	fmt.Fprintln(writer, getPower(a, b, c))
}

func getPower(a, b, c int64) (result int64) {
	if b == 0 {
		result = 1
		return
	}
	if b == 1 {
		result = a % c
		return
	}
	recur := getPower(a, b/2, c) % c
	temp := new(big.Int)
	m := new(big.Int)
	temp = temp.Mul(big.NewInt(recur), big.NewInt(recur))
	m = temp.Mod(temp, big.NewInt(c))
	if b%2 == 0 {
		result = m.Int64()
		return
	}
	result = m.Int64() * a % c
	return
}
