// https://www.acmicpc.net/problem/1021
// 덱을 활용하여 큐를 회전시키는 문제
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var q []queue
	for i := 0; i < n; i++ {
		tmp := queue{}
		tmp.curPos = i + 1
		tmp.initPos = i + 1
		q = append(q, tmp)
	}

	var selectPos []int
	for i := 0; i < m; i++ {
		var pos int
		fmt.Fscanf(reader, "%d ", &pos)
		selectPos = append(selectPos, pos)
	}

	var count int
	for len(selectPos) > 0 && len(q) > 0 {
		for i := 0; i < len(q); i++ {
			if q[i].initPos == selectPos[0] {
				if i < len(q)-i {
					count += i
					q = moveLeft(q, i)
				} else {
					count += (len(q) - i)
					q = moveRight(q, len(q)-i)
				}
				q = q[1:]
				selectPos = selectPos[1:]
				break
			}
		}
	}
	fmt.Fprintln(writer, count)
}

func moveLeft(q []queue, interval int) []queue {
	q = append(q[interval:], q[0:interval]...)
	return q
}

func moveRight(q []queue, interval int) []queue {
	q = append(q[len(q)-interval:len(q)], q[0:len(q)-interval]...)
	return q
}

type queue struct {
	curPos  int
	initPos int
}
