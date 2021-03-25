// https://www.acmicpc.net/problem/1431
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var serials = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &serials[i])
	}

	sort.Slice(serials, func(i, j int) bool {
		if len(serials[i]) < len(serials[j]) {
			return true
		} else if len(serials[i]) == len(serials[j]) {
			var sumI, sumJ int
			for k := 0; k < len(serials[i]); k++ {
				if val, err := strconv.Atoi(string(serials[i][k])); err == nil {
					sumI += val
				}
				if val, err := strconv.Atoi(string(serials[j][k])); err == nil {
					sumJ += val
				}
			}
			if sumI < sumJ {
				return true
			} else if sumI == sumJ {
				return serials[i] < serials[j]
			}
			return false
		}
		return false
	})
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, serials[i])
	}
}
