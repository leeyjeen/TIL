package main

import "fmt"

func solveMeFirst(a int, b int) int {
	return a + b
}

func main() {
	var a, b, res int
	fmt.Scanf("%v\n%v", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}
