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

	var count = getCountOfHansu(n)
	fmt.Println(count)
}

func getCountOfHansu(number int) (count int) {
	if number < 100 {
		count = number
		return
	}
	for i := 100; i <= number; i++ {
		one := i % 10
		ten := i / 10 % 10
		hund := i / 100
		if hund-ten == ten-one {
			count++
		}
	}
	count += 99
	return count
}
