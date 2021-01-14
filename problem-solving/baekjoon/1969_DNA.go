// https://www.acmicpc.net/problem/1969
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var counts = make([]map[string]int, m)
	for i := 0; i < m; i++ {
		counts[i] = map[string]int{} // DNA의 각 위치에 존재하는 뉴클레오티드 종류별 개수를 저장하기 위한 맵 생성
	}
	for i := 0; i < n; i++ {
		var dna string
		fmt.Fscanln(reader, &dna)
		nucleotide := strings.Split(dna, "")
		for j := 0; j < len(nucleotide); j++ {
			counts[j][nucleotide[j]]++
		}
	}

	var answerDNA string
	var answerCount int
	for i := 0; i < len(counts); i++ {
		count := counts[i]
		var pairs []pair
		for k, v := range count {
			pairs = append(pairs, pair{k, v})
		}
		sort.Slice(pairs, func(i, j int) bool {
			if pairs[i].value > pairs[j].value { // 각 위치별로 가장 많이 존재하는 뉴클레오티드를 채택
				return true
			} else if pairs[i].value == pairs[j].value {
				return pairs[i].key < pairs[j].key // 동일한 개수일 경우 사전순으로 채택
			}
			return false
		})
		answerDNA += pairs[0].key         // 채택된 뉴클레오티드 값 추가
		answerCount += n - pairs[0].value // DNA 개수에서 채택된 뉴클레오티드의 개수를 뺀 값을 hamming distance로 계산
	}

	fmt.Fprintln(writer, answerDNA)
	fmt.Fprintln(writer, answerCount)
}

type pair struct {
	key   string
	value int
}
