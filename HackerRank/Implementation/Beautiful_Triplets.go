// https://www.hackerrank.com/challenges/beautiful-triplets/problem
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, d int
	fmt.Fscanln(reader, &n, &d)
	var arr = make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
	}
	fmt.Fprintln(writer, beautifulTriplets(d, arr))
}

func beautifulTriplets(d int, arr []int) int {
	var count int
	for i:=0; i<len(arr)-2; i++ {
		for j:=i+1; j<len(arr)-1; j++ {
			if arr[i]+d > arr[j] {
				continue
			}
			if arr[i]+d < arr[j] {
				break
			}
			for k:=j+1; k<len(arr); k++ {
				if arr[j]+d > arr[k] {
					continue
				}
				if arr[j]+d < arr[k] {
					break
				}
				count++
			}
		}
	}
	return count
}