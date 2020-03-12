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

	var count int

	for true {
		if n%5 != 0 {
			if n < 3 {
				count = -1
				break
			}
			n = n - 3
			count++
		} else {
			break
		}
	}
	if count != -1 {
		count += n / 5
	}
	fmt.Println(count)
}
