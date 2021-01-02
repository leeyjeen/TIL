// https://www.acmicpc.net/problem/18243
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

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	friends := make([][]int, n)
	for i := 0; i < n; i++ {
		friends[i] = make([]int, n)
	}

	for i := 0; i < k; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		friends[a-1][b-1] = 1
		friends[b-1][a-1] = 1
	}

	friends = floydWarshall(friends)
	isBig := checkNetwork(friends)
	if isBig {
		fmt.Fprintln(writer, "Big World!")
	} else {
		fmt.Fprintln(writer, "Small World!")
	}
}

func floydWarshall(friends [][]int) [][]int {
	for i := 0; i < len(friends); i++ {
		for j := 0; j < len(friends); j++ {
			for k := 0; k < len(friends); k++ {
				if j == k { // 자기 자신과 친구인 경우는 없다
					continue
				}
				if friends[j][i] != 0 && friends[i][k] != 0 { // j<->i, i<->k 가 연결되어 있을 때
					if friends[j][k] == 0 || friends[j][k] > friends[j][i]+friends[i][k] {
						friends[j][k] = friends[j][i] + friends[i][k] // j<->k가 연결되어 있지 않거나, i를 거치는 것이 더 적은 단계를 거칠 경우
					}
				}
			}
		}
	}
	return friends
}

func checkNetwork(friends [][]int) (isBig bool) {
	for i := 0; i < len(friends); i++ {
		for j := 0; j < len(friends[i]); j++ {
			if friends[i][j] == 0 && i != j || friends[i][j] > 6 { // 연결되어있지 않거나 6단계를 초과하여 연결된 경우
				isBig = true
				return
			}
		}
	}
	return
}
