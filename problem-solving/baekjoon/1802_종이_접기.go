// https://www.acmicpc.net/problem/1802
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		input, _ := reader.ReadString('\n')
		rules := strings.Split(strings.Replace(input, "\n", "", 1), "")
		if check(rules) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}
}

func check(rules []string) bool {
	if len(rules) == 1 {
		return true
	}
	for i := 0; i < len(rules)/2; i++ {
		if rules[i] == rules[len(rules)-1-i] { // 가운데를 기준으로 접히는 부분은 서로 다른 값이어야 함
			return false
		}
	}
	return check(rules[0:len(rules)/2]) && check(rules[len(rules)/2+1:len(rules)]) // 가운데를 기준으로 왼쪽과 오른쪽을 체크
}
