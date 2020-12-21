// https://www.acmicpc.net/problem/5635
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type studentFormat struct {
	name  string
	day   int
	month int
	year  int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var students []studentFormat
	for i := 0; i < n; i++ {
		student := studentFormat{}
		fmt.Fscanln(reader, &student.name, &student.day, &student.month, &student.year)
		students = append(students, student)
	}
	sort.Slice(students, func(i, j int) bool {
		if students[i].year < students[j].year {
			return true
		} else if students[i].year == students[j].year {
			if students[i].month < students[j].month {
				return true
			} else if students[i].month == students[j].month {
				return students[i].day < students[j].day
			}
		}
		return false
	})
	fmt.Fprintln(writer, students[len(students)-1].name)
	fmt.Fprintln(writer, students[0].name)
}
