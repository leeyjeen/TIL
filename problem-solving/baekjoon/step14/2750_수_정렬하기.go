package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var numbers = make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Fscanln(reader, &numbers[i])
	}

	for i:=0; i<n-1; i++ {
		for j:=i+1; j<n; j++ {
			if numbers[i] > numbers[j] {
				var temp = numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = temp
			}
		}
	}
	
	for i:=0; i <n; i++ {
		fmt.Println(numbers[i])
	}
}