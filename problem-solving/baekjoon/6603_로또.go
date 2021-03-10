// https://www.acmicpc.net/problem/6603
package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	index := 0
	for {
		var k int
		fmt.Fscanf(reader, "%d ", &k)
		if k == 0 {
			break
		}
		if index != 0 {
			fmt.Fprintln(writer, "")
		}
		index++
		var s = make([]int, k)
		for i := 0; i < k; i++ {
			fmt.Fscanf(reader, "%d ", &s[i])
		}
		combinations := Combinations(s, 6)
		sort.Slice(combinations[:], func(i, j int) bool {
			for x := range combinations[i] {
				if combinations[i][x] == combinations[j][x] {
					continue
				}
				return combinations[i][x] < combinations[j][x]
			}
			return false
		})

		for i := 0; i < len(combinations); i++ {
			combiStr := fmt.Sprintf("%v", combinations[i])
			combiStr = combiStr[1 : len(combiStr)-1]
			fmt.Fprintln(writer, combiStr)
		}
	}
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
