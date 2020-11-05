// https://www.acmicpc.net/problem/9012
// 주어진 문자열이 올바른 괄호열인지 판단하는 문제
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
		var inputStr string
		fmt.Fscanln(reader, &inputStr)
		fmt.Fprintln(writer, isVPS(inputStr))
	}
}

func isVPS(inputStr string) string {
	var open, close int
	for _, v := range inputStr {
		if string(v) == "(" {
			open++
		} else if string(v) == ")" {
			close++
		}
		if open < close { // 닫힌 괄호 개수가 더 많아지면 "NO" 리턴
			return "NO"
		}
	}
	if open != close { // 최종 괄호 개수가 다르면 "NO" 리턴
		return "NO"
	}
	return "YES"
}

// // 스택으로 풀기
// func isVPS2(inputStr string) string {
// 	var stack []string
// 	for _, v := range inputStr {
// 		strValue := string(v)
// 		if strValue == "(" {
// 			stack = append(stack, strValue)
// 		} else if strValue == ")" {
// 			if len(stack) > 0 {
// 				stack = stack[:len(stack)-1]
// 			} else {
// 				return "NO"
// 			}
// 		}
// 	}
// 	if len(stack) > 0 {
// 		return "NO"
// 	}
// 	return "YES"
// }
