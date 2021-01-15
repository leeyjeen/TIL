// https://www.acmicpc.net/problem/1145
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

	var numbers = make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscanf(reader, "%d ", &numbers[i])
	}

	var minLCM int = 10000000000
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			lcmOfTwo := getLCM(numbers[i], numbers[j])
			for k := j + 1; k < 5; k++ {
				lcmOfThree := getLCM(lcmOfTwo, numbers[k])
				if lcmOfThree < minLCM {
					minLCM = lcmOfThree
				}
			}
		}
	}
	fmt.Fprintln(writer, minLCM)
}

func getGCD(first, second int) (gcd int) {
	if first < second {
		second, first = first, second
	}

	for second != 0 {
		first, second = second, first%second
	}
	return first
}

func getLCM(first, second int) (lcm int) {
	return first * second / getGCD(first, second)
}
