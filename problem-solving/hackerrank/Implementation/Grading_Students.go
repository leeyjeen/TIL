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

	var grades = make([]int, n)
	var newGrades = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &grades[i])
		if grades[i] > 36 && grades[i]%5 >= 3 {
			newGrades[i] = grades[i] + 5 - grades[i]%5
		} else {
			newGrades[i] = grades[i]
		}
		fmt.Fprintln(writer, newGrades[i])
	}
}
