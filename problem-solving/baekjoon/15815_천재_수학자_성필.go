// https://www.acmicpc.net/problem/15815
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var formula string
	fmt.Fscanln(reader, &formula)
	formulas := strings.Split(formula, "")

	numbers := []int{}
	for len(formulas) > 0 {
		popped := formulas[0]
		formulas = formulas[1:]

		if popped == "+" || popped == "-" || popped == "*" || popped == "/" {
			first := numbers[len(numbers)-2]
			second := numbers[len(numbers)-1]
			numbers = append(numbers[:len(numbers)-2], compute(popped, first, second))
		} else {
			number, _ := strconv.Atoi(popped)
			numbers = append(numbers, number)
		}
	}
	fmt.Fprintln(writer, numbers[0])
}

func compute(operator string, first, second int) int {
	if operator == "+" {
		return first + second
	} else if operator == "-" {
		return first - second
	} else if operator == "*" {
		return first * second
	} else {
		return first / second
	}
}
