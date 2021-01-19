// https://www.acmicpc.net/problem/2303
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

	var cards = make([][]int, n)
	for i := 0; i < n; i++ {
		cards[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			fmt.Fscanf(reader, "%d ", &cards[i][j])
		}
	}
	fmt.Fprintln(writer, getWinner(cards))
}

func getWinner(cards [][]int) int {
	var maxSum, winner int
	var sumOne, sumTwo, sumThree int
	for i := 0; i < len(cards); i++ {
		for j := 0; j < 5; j++ {
			sumOne = cards[i][j]
			for k := j + 1; k < 5; k++ {
				sumTwo = sumOne + cards[i][k]
				for l := k + 1; l < 5; l++ {
					sumThree = sumTwo + cards[i][l]
				}
				if maxSum <= sumThree%10 {
					maxSum = sumThree % 10
					winner = i + 1
				}
			}
		}
	}
	return winner
}
