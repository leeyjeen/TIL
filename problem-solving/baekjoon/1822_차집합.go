// https://www.acmicpc.net/problem/1822
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var aCount, bCount int
	fmt.Fscanln(reader, &aCount, &bCount)
	var elements = map[int]bool{}
	for i := 0; i < aCount; i++ {
		var elem int
		fmt.Fscanf(reader, "%d ", &elem)
		elements[elem] = true
	}
	for i := 0; i < bCount; i++ {
		var elem int
		fmt.Fscanf(reader, "%d ", &elem)
		delete(elements, elem)
	}
	fmt.Fprintln(writer, len(elements))
	var elemArray []int
	for key := range elements {
		elemArray = append(elemArray, key)
	}
	sort.Ints(elemArray)
	arrToStr := fmt.Sprint(elemArray)
	fmt.Fprint(writer, arrToStr[1:len(arrToStr)-1])
}
