package main

import (
	"fmt"
	"bufio"
	"os"
	"math"
	)

func main() {
	var a,b,c float64
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	
	for true {
		fmt.Fscanln(reader, &a, &b, &c)
		if a ==0 &&b==0&& c==0 {
			break
		}
		aPow := math.Pow(a, 2)
		bPow := math.Pow(b, 2)
		cPow := math.Pow(c, 2)

		var result = "wrong"
		if aPow > bPow && aPow > cPow {
			if aPow - bPow - cPow == 0 {
				result = "right"
			}
		} else if bPow > cPow && bPow > aPow {
			if bPow -  aPow - cPow == 0 {
				result = "right"
			}
		} else {
			if cPow - aPow - bPow == 0 {
				result = "right"
			}
		}
		fmt.Fprintln(writer, result)
	}
}