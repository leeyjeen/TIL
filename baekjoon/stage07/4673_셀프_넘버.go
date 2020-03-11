package main

import "fmt"

func main() {
	selfNumbers := selfNumber(10000)
	for i := 1; i < len(selfNumbers); i++ {
		if selfNumbers[i] == false {
			fmt.Println(i)
		}
	}
}

func selfNumber(n int) (checkSelfNumber map[int]bool) {
	checkSelfNumber = make(map[int]bool, n+1)
	for i := 0; i < n+1; i++ {
		checkSelfNumber[i] = false
	}

	for i := 0; i < n+1; i++ {
		var sum = i
		var number = i
		for j := number; j != 0; j /= 10 {
			sum += j % 10
		}
		if sum <= n {
			checkSelfNumber[sum] = true
		}
	}

	return checkSelfNumber
}
