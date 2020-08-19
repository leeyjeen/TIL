package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var seqFirst, seqSecond string
	fmt.Fscanln(reader, &seqFirst)
	fmt.Fscanln(reader, &seqSecond)

	s1 := strings.Split(seqFirst, "")
	s2 := strings.Split(seqSecond, "")

	var result = lcs(s1, s2)
	fmt.Fprintln(writer, result)
}

func lcs(s1, s2 []string) int {
	var LCS = make([][]int, len(s1)+1)
	for i := range LCS {
		LCS[i] = make([]int, len(s2)+1)
	}
	for i := 0; i < len(s1)+1; i++ {
		for j := 0; j < len(s2)+1; j++ {
			if i == 0 || j == 0 {
				continue
			}
			if s1[i-1] == s2[j-1] {
				LCS[i][j] = LCS[i-1][j-1] + 1
			} else {
				LCS[i][j] = int(math.Max(float64(LCS[i-1][j]), float64(LCS[i][j-1])))
			}
		}
	}
	return LCS[len(s1)][len(s2)]
}
