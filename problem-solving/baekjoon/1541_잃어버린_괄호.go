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
	var formula string
	fmt.Fscanln(reader, &formula)
	formulas := strings.Split(formula, "-")
	fmt.Println(greedyAlgorithm(formulas))
}

func greedyAlgorithm(formulas []string) (result int) {
	var values = make([]int, len(formulas))
	for i := 0; i < len(formulas); i++ {
		for _, v := range strings.Split(formulas[i], "+") {
			intVal, _ := strconv.Atoi(v)
			values[i] += intVal
		}
		if i == 0 {
			result = values[i]
		} else {
			result -= values[i]
		}
	}
	return
}
