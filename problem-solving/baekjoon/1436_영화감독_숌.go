package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var number int
	var cur int
	for true {
		number++
		var tmpNum = number
		var found = false
		for true {
			if tmpNum == 0 {
				break
			}
			if tmpNum%1000 == 666 {
				cur++
				if cur == n {
					found = true
				}
				break
			}
			tmpNum /= 10
		}
		if found {
			fmt.Println(number)
			break
		}
	}
}
