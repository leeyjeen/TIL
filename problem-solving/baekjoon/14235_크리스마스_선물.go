// https://www.acmicpc.net/problem/14235
// runtime error..
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	writer := bufio.NewWriter(os.Stdout)
// 	defer writer.Flush()

// 	var n int
// 	fmt.Fscanln(reader, &n)
// 	var heap Heap
// 	heap.heapArr = make([]uint64, 5001)

// 	for i := 0; i < n; i++ {
// 		var a string
// 		a, _ = reader.ReadString('\n') // 공백 포함하여 입력 받기 위해 ReadString() 사용
// 		aArr := strings.Split(strings.ReplaceAll(a, "\n", ""), " ")
// 		if aArr[0] == "0" {
// 			if heap.numOfData == 0 {
// 				fmt.Fprintln(writer, -1)
// 			} else {
// 				fmt.Fprintln(writer, heap.deleteHeap())
// 			}
// 		} else {
// 			aLen, _ := strconv.Atoi(aArr[0])
// 			for j := 1; j < aLen+1; j++ {
// 				u64, _ := strconv.ParseUint(aArr[j], 10, 64)
// 				heap.insertHeap(u64)
// 			}
// 		}
// 	}
// }

// type Heap struct {
// 	heapArr   []uint64
// 	numOfData int
// }

// func (heap *Heap) deleteHeap() (delVal uint64) {
// 	delVal = heap.heapArr[1]
// 	lastVal := heap.heapArr[heap.numOfData]
// 	var parentIdx int = 1
// 	var childIdx int = heap.getPriorityChildIdx(parentIdx)
// 	for childIdx != -1 {
// 		if lastVal > heap.heapArr[childIdx] {
// 			break
// 		}
// 		heap.heapArr[parentIdx] = heap.heapArr[childIdx]
// 		parentIdx = childIdx
// 		childIdx = heap.getPriorityChildIdx(parentIdx)
// 	}
// 	heap.heapArr[parentIdx] = lastVal
// 	heap.numOfData--
// 	return
// }

// func (heap *Heap) getPriorityChildIdx(parentIdx int) int {
// 	if parentIdx*2 > heap.numOfData {
// 		return -1
// 	} else if parentIdx*2 == heap.numOfData {
// 		return parentIdx * 2
// 	} else {
// 		if heap.heapArr[parentIdx*2] < heap.heapArr[parentIdx*2+1] {
// 			return parentIdx*2 + 1
// 		}
// 		return parentIdx * 2
// 	}
// }

// func (heap *Heap) insertHeap(x uint64) {
// 	heap.numOfData++
// 	var idx = heap.numOfData

// 	for idx != 1 {
// 		if x > heap.heapArr[idx/2] {
// 			heap.heapArr[idx] = heap.heapArr[idx/2]
// 			idx = idx / 2
// 		} else {
// 			break
// 		}
// 	}
// 	heap.heapArr[idx] = x
// }
