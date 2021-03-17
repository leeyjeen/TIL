// https://www.acmicpc.net/problem/10974
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
	numbers := []int{}
	for i := 0; i < n; i++ {
		numbers = append(numbers, i+1)
	}
	permutations := getPermutations(numbers)
	for i := 0; i < len(permutations); i++ {
		permutationStr := fmt.Sprintf("%v", permutations[i])
		fmt.Fprintln(writer, permutationStr[1:len(permutationStr)-1])
	}
}

func getPermutations(elements []int) [][]int {
	permutations := [][]int{}
	if len(elements) == 1 {
		permutations = [][]int{elements}
		return permutations
	}
	for i := range elements {
		el := append([]int(nil), elements...)

		for _, perm := range getPermutations(append(el[0:i], el[i+1:]...)) {
			permutations = append(permutations, append([]int{elements[i]}, perm...))
		}
	}
	return permutations
}
