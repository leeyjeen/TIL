// https://www.acmicpc.net/problem/4659
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		var password string
		fmt.Fscanln(reader, &password)
		if password == "end" {
			break
		}

		if checkPasswordAcceptable(password) {
			fmt.Fprintf(writer, "<%s> is acceptable.\n", password)
		} else {
			fmt.Fprintf(writer, "<%s> is not acceptable.\n", password)
		}
	}
}

func checkPasswordAcceptable(password string) bool {
	var prevLetter string
	var totalVowelCount, vowelCount, consonantCount int
	for i := 0; i < len(password); i++ {
		var curLetter = string(password[i])
		if isSame(prevLetter, curLetter) {
			return false
		}
		if isVowel(curLetter) {
			vowelCount++
			totalVowelCount++
			if vowelCount == 3 {
				return false
			}
			consonantCount = 0
		} else {
			consonantCount++
			if consonantCount == 3 {
				return false
			}
			vowelCount = 0
		}
		prevLetter = curLetter
	}
	if totalVowelCount == 0 {
		return false
	}
	return true
}

func isVowel(letter string) bool {
	var vowels = []string{"a", "e", "i", "o", "u"}
	for _, vowel := range vowels {
		if letter == vowel {
			return true
		}
	}
	return false
}

func isSame(prevLetter, letter string) bool {
	if prevLetter == letter && letter != "e" && letter != "o" {
		return true
	}
	return false
}
