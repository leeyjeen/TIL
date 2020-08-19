// https://www.hackerrank.com/challenges/happy-ladybugs/problem
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

	var g int
	fmt.Fscanln(reader, &g)
	for i := 0; i < g; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		var b string
		fmt.Fscanln(reader, &b)
		fmt.Fprintln(writer, happyLadybugs(b))
	}

}

func happyLadybugs(b string) string {
	ladyMap := map[rune]int{}
	var alreadyHappy bool = true
	var sameCount int = 1
	for i, v := range b {
		if i > 0 {
			if b[i-1] == byte(v) {
				sameCount++
			} else if b[i-1] != byte(v) && sameCount == 1 {
				alreadyHappy = false
				break
			} else if sameCount > 1 {
				sameCount = 1
			}
		}

	}
	for _, v := range b {
		if _, ok := ladyMap[v]; ok {
			ladyMap[v]++
		} else {
			ladyMap[v] = 1
		}
	}
	for key, val := range ladyMap {
		if val <= 1 && fmt.Sprintf("%c", key) != "_" {
			return "NO"
		}
	}
	if _, ok := ladyMap[95]; !ok && !alreadyHappy {
		return "NO"
	}
	return "YES"
}
