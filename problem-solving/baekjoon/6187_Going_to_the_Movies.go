// https://www.acmicpc.net/problem/6187
package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var c, n int
	fmt.Fscanln(reader, &c, &n)
	var w = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &w[i])
	}
	fmt.Fprintln(writer, getHeaviest(c, w))
}

func getHeaviest(c int, w []int) int {
	maxSum := 0
	for i := len(w); i > 0; i-- {
		for _, v := range Combinations(w, i) {
			sum := 0
			for j := 0; j < len(v); j++ {
				if sum+v[j] <= c {
					sum += v[j]
				} else {
					break
				}
			}
			if maxSum < sum {
				maxSum = sum
			}
			if maxSum == c {
				return maxSum
			}
		}
	}
	return maxSum
}

func Combinations(set []int, n int) (subsets [][]int) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []int

		for object := uint(0); object < length; object++ {
			if (subsetBits>>object)&1 == 1 {
				subset = append(subset, set[object])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}
