package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	values := [][]int{}
	for i := 0; i < n; i++ {
		inputs := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			fmt.Fscanf(reader, "%d ", &inputs[j])
		}
		values = append(values, inputs)
	}
	maxVal := calculate(values, n)
	fmt.Fprintln(writer, maxVal)
}

func calculate(values [][]int, n int) int {
	if n == 1 {
		return values[0][0]
	}
	maxValues := [][]int{}
	maxValues = append(maxValues, []int{values[0][0]})

	var maxVal int

	for i := 1; i < n; i++ {
		val := []int{}
		for j := 0; j < i+1; j++ {
			if j-1 < 0 {
				maxVal = maxValues[i-1][j]
			} else if j > i-1 {
				fmt.Println(j, i-1)
				maxVal = maxValues[i-1][j-1]
			} else {
				maxVal = int(math.Max(float64(maxValues[i-1][j]), float64(maxValues[i-1][j-1])))
			}
			val = append(val, maxVal+values[i][j])
		}
		maxValues = append(maxValues, val)
	}
	return maxOfArray(maxValues[n-1])
}

func maxOfArray(arr []int) (max int) {
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return
}
