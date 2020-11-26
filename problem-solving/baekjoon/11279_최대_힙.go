// https://www.acmicpc.net/problem/11279
// 최댓값을 빠르게 뽑는 자료구조를 배우는 문제
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
			fmt.Fprintln(writer, deleteHeap(&heap))
		} else {
			insertHeap(&heap, x)
		}
	}
}

type Heap struct {
	heapArr   []uint
	numOfData int
}

// deleteHeap: 힙에서 데이터 삭제.
// 힙의 마지막 노드를 루트 노드의 위치에 올린 다음,
// 자식 노드와의 비교과정을 거치면서 자신의 위치를 찾을 때까지 내린다.
func deleteHeap(heap *Heap) (delVal uint) {
	if heap.numOfData == 0 {
		return
	}
	delVal = heap.heapArr[1]
	lastVal := heap.heapArr[heap.numOfData]
	var parentIdx int = 1 // 마지막 노드가 저장될 인덱스
	var childIdx int = getPriorityChildIdx(heap, parentIdx)
	for childIdx != -1 { // 자식 노드가 존재하는 동안 반복
		if lastVal > heap.heapArr[childIdx] { // 마지막 노드가 우선순위가 높은 경우 반복문 탈출
			break
		}
		heap.heapArr[parentIdx] = heap.heapArr[childIdx]
		parentIdx = childIdx
		childIdx = getPriorityChildIdx(heap, parentIdx)
	}
	heap.heapArr[parentIdx] = lastVal
	heap.numOfData--
	return
}

// getPriorityChildIdx:
// 자식 노드 중 우선순위가 높은 자식의 인덱스 반환.
// 자식 노드가 존재하지 않는 경우 -1을 반환한다.
func getPriorityChildIdx(heap *Heap, parentIdx int) int {
	if parentIdx*2 > heap.numOfData {
		return -1
	} else if parentIdx*2 == heap.numOfData {
		return parentIdx * 2
	} else {
		if heap.heapArr[parentIdx*2] < heap.heapArr[parentIdx*2+1] {
			return parentIdx*2 + 1
		}
		return parentIdx * 2
	}
}

// insertHeap: 힙에 데이터 저장.
// 마지막 노드에서부터 부모 노드와 우선순위 비교과정을 거쳐 자신의 위치를 찾는다.
func insertHeap(heap *Heap, x uint) {
	heap.numOfData++
	var idx = heap.numOfData // 새 요소가 저장될 인덱스 값을 idx에 저장

	// 새 요소가 저장될 위치가 루트 노드의 위치가 아닌 경우 반복
	for idx != 1 {
		// 새 노드와 부모 노드의 우선순위 비교
		if x > heap.heapArr[idx/2] { // 새 노드의 우선순위가 높다면
			heap.heapArr[idx] = heap.heapArr[idx/2] // 부모 노드를 한 레벨 내린다
			idx = idx / 2                           // 새 노드를 한 레벨 올리기 위해 인덱스 값 갱신
		} else { // 새 노드의 우선순위가 높지 않다면
			break
		}
	}
	heap.heapArr[idx] = x // 새 노드를 배열에 저장
}
