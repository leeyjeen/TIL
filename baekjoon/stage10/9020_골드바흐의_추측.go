package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
)

func main() {
	var t, n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	primeNumbers := getPrimeNumbers(10000)
	fmt.Fscanln(reader,&t)
	
	for i:=0; i<t; i++ {
		fmt.Fscanln(reader, &n)
		printGoldBach(n,primeNumbers, writer)
	}
}

func getPrimeNumbers(num int) (primeNumbers map[int]bool) {
	primeNumbers = make(map[int]bool, num)
	for i:=0; i<num; i++ {
		primeNumbers[i+1] = true
	}
	for i:=0; i<num; i++ {
		if primeNumbers[i+1] == false {
			continue
		}
		var divisionCount int
		for j:=0; j<int(math.Sqrt(float64(i+1))); j++ {
			if (i+1) % (j+1) == 0 {
				divisionCount++
				if divisionCount >1 {
					for k:=1; k*(i+1)<=num; k++ {
						primeNumbers[k*(i+1)] = false
					}
					break
				}
			}
		}
	}
	return primeNumbers
}

func printGoldBach(num int, primeNumbers map[int]bool, writer *bufio.Writer) {
	for i:=num/2; i>0; i-- {
		if primeNumbers[i] && primeNumbers[num-i] {
			fmt.Fprintf(writer, "%d %d\n", i, num-i)
			break
		}
	}
}