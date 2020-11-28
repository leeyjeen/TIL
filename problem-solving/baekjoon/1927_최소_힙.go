// https://www.acmicpc.net/problem/1927
// 최솟값을 빠르게 뽑는 문제
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

	var n int
	fmt.Fscanln(reader, &n)

	var heap Heap
	heap.heapArr = make([]uint, n+1)

	for i := 0; i < n; i++ {
		var x uint
		fmt.Fscanln(reader, &x)
		if x == 0 {
			fmt.Fprintln(writer, heap.deleteHeap())
		} else {
			heap.insertHeap(x)
		}
	}
}

type Heap struct {
	heapArr   []uint
	numOfData int
}

func (heap *Heap) deleteHeap() (delVal uint) {
	if heap.numOfData == 0 {
		return
	}
	delVal = heap.heapArr[1]
	lastVal := heap.heapArr[heap.numOfData]
	var parentIdx = 1
	var childIdx = heap.getPriorityChildIdx(parentIdx)
	for childIdx != -1 {
		if lastVal < heap.heapArr[childIdx] {
			break
		}
		heap.heapArr[parentIdx] = heap.heapArr[childIdx]
		parentIdx = childIdx
		childIdx = heap.getPriorityChildIdx(parentIdx)
	}
	heap.heapArr[parentIdx] = lastVal
	heap.numOfData--
	return
}

func (heap *Heap) getPriorityChildIdx(parentIdx int) int {
	if parentIdx*2 > heap.numOfData {
		return -1
	} else if parentIdx*2 == heap.numOfData {
		return parentIdx * 2
	} else {
		if heap.heapArr[parentIdx*2] > heap.heapArr[parentIdx*2+1] {
			return parentIdx*2 + 1
		}
		return parentIdx * 2
	}
}

func (heap *Heap) insertHeap(x uint) {
	heap.numOfData++
	var idx = heap.numOfData

	for idx != 1 {
		if x < heap.heapArr[idx/2] {
			heap.heapArr[idx] = heap.heapArr[idx/2]
			idx = idx / 2
		} else {
			break
		}
	}
	heap.heapArr[idx] = x
}
