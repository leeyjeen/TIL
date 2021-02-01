// https://www.acmicpc.net/problem/1543
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var document, word string
	document, _ = reader.ReadString('\n')
	document = strings.ReplaceAll(document, "\n", "")
	word, _ = reader.ReadString('\n')
	word = strings.ReplaceAll(word, "\n", "")
	fmt.Println(strings.Count(document, word))
}
