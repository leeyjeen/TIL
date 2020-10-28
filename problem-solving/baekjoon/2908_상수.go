package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var a, b, result int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &a, &b)

	a = (a%10)*100 + (a/10)%10*10 + a/100
	b = (b%10)*100 + (b/10)%10*10 + b/100
	if a > b {
		result = a
	} else {
		result = b
	}
	fmt.Println(result)
}