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

	var word string
	var groupCount int
	for i:=0; i<n; i++ {
		fmt.Fscanln(reader, &word)
		var letters = make(map[uint8]bool)
		var prev uint8
		var isGroup = true
		for j:=0; j<len(word); j++ {
			_, exist := letters[word[j]]
			if !exist {
				letters[word[j]] = true	
			} else if prev != word[j] {
				isGroup = false
				break
			}
			prev = word[j]
		}
		if isGroup {
			groupCount++
		}
	}

	fmt.Println(groupCount)
}