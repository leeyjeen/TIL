package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &s)

	var letters = make(map[int]int)
	for i:=0; i<26; i++ {
		letters[i+97] = -1
	}

	for i:=0; i<len(s);i++{
		ascii , _ := strconv.Atoi(fmt.Sprintf("%d", s[i]))
		if letters[ascii] == -1 {
			letters[ascii] = i
		}
	}

	for i:=0; i<26; i++ {
		fmt.Printf("%d ", letters[i+97])
	}
	fmt.Print("\n")
}