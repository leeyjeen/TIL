package main

import "fmt"

func HanoiTowerMove(num int, from, by, to string) {
	if num == 1 {
		fmt.Printf("moved %d from %s to %s \n", num, from, to)
	} else {
		HanoiTowerMove(num-1, from, to, by)
		fmt.Printf("moved %d from %s to %s \n", num, from, to)
		HanoiTowerMove(num-1, by, from, to)
	}
}

func main() {
	HanoiTowerMove(3, "A", "B", "C")
}