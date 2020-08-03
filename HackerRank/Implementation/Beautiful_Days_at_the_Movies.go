package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var i, j, k int
	fmt.Fscanln(reader, &i, &j, &k)
	fmt.Fprintln(writer, beautifulDays(i, j, k))
}

func beautifulDays(i, j, k int) int {
	var count int
	for v := i; v < j+1; v++ {
		strV := fmt.Sprint(v)
		arrV := strings.Split(strV, "")
		var reversed int
		for j := 0; j < len(arrV); j++ {
			n, _ := strconv.Atoi(arrV[j])
			reversed += n * int(math.Pow(10.0, float64(j)))
		}
		var diff = reversed - v
		if diff%k == 0 {
			count++
		}
	}
	return count
}
