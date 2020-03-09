package main

import (
	"fmt"
)

func main() {
	var score int
	var grade string
	fmt.Scanf("%d", &score)
	if 90 <= score {
		grade ="A"
	} else if 80 <= score{
		grade = "B"
	} else if 70 <= score {
		grade = "C"
	} else if 60 <= score {
		grade = "D"
	} else {
		grade = "F"
	}
	fmt.Println(grade)
}