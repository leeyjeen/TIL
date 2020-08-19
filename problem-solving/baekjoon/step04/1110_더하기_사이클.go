package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, cycle int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var newN = n
	for true {
		cycle++
		var sum int
		if newN/10 == 0 {
			sum = newN
		} else {
			sum = newN/10 + newN%10
		}
		newN = newN%10*10 + sum%10
		if newN == n {
			break
		}
	}
	fmt.Println(cycle)
}
