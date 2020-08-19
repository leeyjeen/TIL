package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
)

func main() {
	var a,b,c float64
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &a, &b, &c)

	if c-b <= 0 {
		fmt.Println(-1)
		return
	}
	
	roundUp := int(math.Ceil(a/(c-b)))
	if roundUp == int(a/(c-b)) {
		fmt.Println(roundUp + 1)
	} else {
		fmt.Println(roundUp)
	}
}