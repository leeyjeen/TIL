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

	var creator int
	for i:=0; i<n; i++ {
		var splitSum = getSplitSum(i)
		if splitSum == n {
			creator = i
			break
		}
	}
	fmt.Println(creator)
}

func getSplitSum(n int) (result int) {
	result = n
	for n != 0 {
		result += n%10
		n /= 10
	}
	return
}