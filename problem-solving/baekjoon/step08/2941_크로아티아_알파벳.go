package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var word string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &word)

	word = strings.Replace(word, "c=", "!", -1)
	word = strings.Replace(word, "c-", "@", -1)
	word = strings.Replace(word, "dz=", "#", -1)
	word = strings.Replace(word, "d-", "$", -1)
	word = strings.Replace(word, "lj", "%", -1)
	word = strings.Replace(word, "nj", "^", -1)
	word = strings.Replace(word, "s=", "&", -1)
	word = strings.Replace(word, "z=", "*", -1)

	var alphabets []string
	for i := 0; i < len(word); i++ {
		alphabets = append(alphabets, string(word[i]))
	}

	fmt.Println(len(alphabets))
}
