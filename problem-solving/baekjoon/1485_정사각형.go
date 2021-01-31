// https://www.acmicpc.net/problem/1485
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var coordinates []coordinateFormat
		for j := 0; j < 4; j++ {
			var coordinate = coordinateFormat{}
			fmt.Fscanln(reader, &coordinate.x, &coordinate.y)
			coordinates = append(coordinates, coordinate)
		}
		sort.Slice(coordinates, func(i, j int) bool {
			if coordinates[i].x < coordinates[j].x {
				return true
			} else if coordinates[i].x == coordinates[j].x {
				return coordinates[i].y < coordinates[j].y
			}
			return false
		})
		fmt.Fprintln(writer, checkSquare(coordinates))
	}
}

type coordinateFormat struct {
	x float64
	y float64
}

// if square return 1 else return 0
func checkSquare(c []coordinateFormat) int {
	sideOne := math.Pow(c[0].x-c[1].x, 2.0) + math.Pow(c[0].y-c[1].y, 2.0)
	sideTwo := math.Pow(c[1].x-c[3].x, 2.0) + math.Pow(c[1].y-c[3].y, 2.0)
	if sideOne != sideTwo {
		return 0
	}
	sideThree := math.Pow(c[3].x-c[2].x, 2.0) + math.Pow(c[3].y-c[2].y, 2.0)
	if sideTwo != sideThree {
		return 0
	}
	sideFour := math.Pow(c[2].x-c[0].x, 2.0) + math.Pow(c[2].y-c[0].y, 2.0)
	if sideThree != sideFour {
		return 0
	}
	diagonalOne := math.Pow(c[0].x-c[3].x, 2.0) + math.Pow(c[0].y-c[3].y, 2.0)
	diagonalTwo := math.Pow(c[2].x-c[1].x, 2.0) + math.Pow(c[2].y-c[1].y, 2.0)
	if diagonalOne != diagonalTwo || sideOne+sideTwo != diagonalOne {
		return 0
	}
	return 1
}
