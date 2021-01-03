// https://www.acmicpc.net/problem/11719
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 100; i++ {
		var input string
		input, _ = reader.ReadString('\n') // 공백 포함하여 입력 받기 위해 ReadString() 사용
		input = strings.ReplaceAll(input, "\n", "")
		fmt.Println(input)
	}
}
