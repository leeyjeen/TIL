// https://www.acmicpc.net/problem/2910
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, c int
	fmt.Fscanln(reader, &n, &c)
	var frequency = map[int]int{}
	var order = map[int]int{}
	for i := 0; i < n; i++ {
		var number int
		fmt.Fscanf(reader, "%d ", &number)
		frequency[number]++
		if _, ok := order[number]; !ok {
			order[number] = i + 1
		}
	}

	var frequencies = []frequencyFormat{}
	for key, val := range frequency {
		frequencies = append(frequencies, frequencyFormat{key, val, order[key]})
	}

	sort.Slice(frequencies, func(i, j int) bool {
		if frequencies[i].frequency > frequencies[j].frequency {
			return true
		} else if frequencies[i].frequency == frequencies[j].frequency {
			return frequencies[i].order < frequencies[j].order
		}
		return false
	})

	for i := 0; i < len(frequencies); i++ {
		tmp := frequencies[i]
		for j := 0; j < tmp.frequency; j++ {
			fmt.Fprintf(writer, "%d ", tmp.number)
		}
	}
}

type frequencyFormat struct {
	number    int
	frequency int
	order     int
}
