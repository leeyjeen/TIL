package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var y int
	fmt.Fscanln(reader, &y)
	daysOfMonth := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if y < 1918 {
		if y % 4 == 0 {
			daysOfMonth[1] = 29
		}
	} else if y == 1918 {
		daysOfMonth[1] = 15
	} else {
		if y %400 == 0 || y%4 == 0 && y % 100 != 0 {
			daysOfMonth[1] = 29
		}
	}
	var sum, month, day int
	for i, v := range daysOfMonth {
		month = i+1
		if sum + v > 256 {
			day = 256 - sum
			break
		} 
		sum += v
	}

	var monthStr string = fmt.Sprintf("%d", month)
	if month/10 < 1 {
		monthStr = "0"+monthStr
	}
	fmt.Fprintln(writer, fmt.Sprintf("%d.%s.%d", day, monthStr, y))
}