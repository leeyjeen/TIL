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

	var s string
	fmt.Fscanln(reader, &s)

	var splits []string
	splits = strings.Split(s, ":")
	hours, _ := strconv.Atoi(splits[0])
	minutes := splits[1]
	seconds := splits[2][:2]
	if splits[2][2:] == "PM" && hours != 12 {
		hours = hours + 12
	} else if splits[2][2:] == "AM" && hours == 12 {
		hours = hours - 12
	}
	hourString := fmt.Sprint(hours)
	if len(hourString) == 1 {
		hourString = "0" + hourString
	}
	fmt.Fprintln(writer, hourString+":"+minutes+":"+seconds)
}
