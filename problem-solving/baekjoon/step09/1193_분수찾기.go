package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var x int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &x)

	var level = 1
	var currentNum = 1
	for true {
		if currentNum > x {
			break
		}
		level++
		for i := level; i > 0; i-- {
			currentNum++
			if currentNum == x {
				if level%2 == 1 {
					fmt.Println(strconv.Itoa(i) + "/" + strconv.Itoa(level+1-i))
				} else {
					fmt.Println(strconv.Itoa(level+1-i) + "/" + strconv.Itoa(i))
				}
			}
		}
	}
}
