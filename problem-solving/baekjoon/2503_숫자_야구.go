// https://www.acmicpc.net/problem/2503
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

	var games []Game
	for i := 0; i < n; i++ {
		game := Game{}
		fmt.Fscanln(reader, &game.number, &game.strike, &game.ball)
		games = append(games, game)
	}

	var count int                 // 가능한 답의 개수
	for i := 123; i <= 987; i++ { // 모든 경우 계산
		var isAnswer bool = true
		a, b, c := i/100, (i/10)%10, i%10
		if !checkValid(a, b, c) { // 각 자릿수가 1-9의 서로 다른 숫자 세개로 구성되었는지 체크
			continue
		}

		for j := 0; j < n; j++ { // i값이 민혁이의 질문과 영수의 대답의 모든 경우에 대해 일치하는 경우 count 증가
			strike, ball, number := 0, 0, games[j].number
			a2, b2, c2 := number/100, number/10%10, number%10
			if a == a2 {
				strike++
			}
			if a == b2 || a == c2 {
				ball++
			}
			if b == b2 {
				strike++
			}
			if b == a2 || b == c2 {
				ball++
			}
			if c == c2 {
				strike++
			}
			if c == a2 || c == b2 {
				ball++
			}
			if !(strike == games[j].strike && ball == games[j].ball) {
				isAnswer = false
				break
			}
		}
		if isAnswer {
			count++
		}
	}

	fmt.Fprintln(writer, count)
}

type Game struct {
	number int
	strike int
	ball   int
}

func checkValid(a, b, c int) bool {
	if b == 0 || c == 0 {
		return false
	}
	if a == b || b == c || c == a {
		return false
	}
	return true
}
