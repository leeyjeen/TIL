package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &input)

	var letters = make(map[uint8]int)
	for i:=0; i<26; i++ {
		letters[uint8(i)+65] = 0
	}

	for i:=0; i<len(input); i++ {
		ascii := input[i]
		if ascii > 90 {
			ascii -= 32
		}
		letters[ascii]++
	}

	var maxVal = -1
	var maxKey string
	for key, val := range letters {
		if val > maxVal {
			maxVal = val
			maxKey = string(key)
		} else if val == maxVal {
			maxKey = "?"
		}
	}

	fmt.Println(maxKey)
}