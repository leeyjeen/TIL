// https://www.acmicpc.net/problem/5568
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
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &k)
	var cards = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &cards[i])
	}

	fmt.Fprintln(writer, getCountOfNumbers(cards, k))
}

func getCountOfNumbers(cards []string, k int) int {
	var resultMap = map[string]int{}
	for i := 0; i < len(cards); i++ {
		resultMap = recursion(i, cards, []string{cards[i]}, k, resultMap)
	}
	return len(resultMap)
}

func recursion(i int, cards []string, curList []string, k int, resultMap map[string]int) map[string]int {
	if len(curList) == k {
		permutations := getPermutations(curList)

		for i := 0; i < len(permutations); i++ {
			var key string
			for j := 0; j < len(permutations[i]); j++ {
				key += permutations[i][j]
			}
			resultMap[key] = 1
		}
		return resultMap
	}
	for j := i + 1; j < len(cards); j++ {
		resultMap = recursion(j, cards, append(curList, cards[j]), k, resultMap)
	}
	return resultMap
}

func getPermutations(elements []string) [][]string {
	permutations := [][]string{}
	if len(elements) == 1 {
		permutations = [][]string{elements}
		return permutations
	}
	for i := range elements {
		el := append([]string(nil), elements...)

		for _, perm := range getPermutations(append(el[0:i], el[i+1:]...)) {
			permutations = append(permutations, append([]string{elements[i]}, perm...))
		}
	}
	return permutations
}
