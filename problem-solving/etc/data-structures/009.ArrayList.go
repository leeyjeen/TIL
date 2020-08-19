package main

import (
	"fmt"
)

const LIST_LEN int = 100

type ArrayList struct {	// 배열 기반 리스트 구조체
	arr [LIST_LEN]int	// 리스트 저장소인 배열
	numOfData int	// 저장된 데이터의 수
	curPosition int // 데이터 참조 위치
}

type IArrayList interface {
	ListInit()	// 초기화
	LInsert(data int) // 데이터 저장
	LFirst() (int, bool) // 첫 데이터 참조
	LNext() (int, bool) // 두 번째 이후 데이터 참조
	LRemove() int // 참조한 데이터 삭제
	LCount() int // 저장된 데이터 수 반환
}

func (al *ArrayList) ListInit() {
	al.numOfData = 0
	al.curPosition = -1
}

func (al *ArrayList) LInsert(data int) {
	if al.numOfData > LIST_LEN {
		fmt.Println("Cannot save..")
		return
	}
	al.arr[al.numOfData] = data
	al.numOfData++
}
func (al *ArrayList) LFirst() (int, bool) {
	if al.numOfData == 0 {
		return 0, false
	}
	al.curPosition = 0
	return al.arr[al.curPosition], true
}
func (al *ArrayList) LNext() (int, bool) {
	if al.curPosition >= al.numOfData - 1 {
		return 0, false
	}
	al.curPosition++
	return  al.arr[al.curPosition], true
}
func (al *ArrayList) LRemove() int {
	var rPos = al.curPosition
	var num = al.numOfData
	var rData = al.arr[rPos]
	for i := rPos; i< num-1; i++ {
		al.arr[i] = al.arr[i+1]
	}
	al.numOfData--
	al.curPosition--
	return rData
}

func (al *ArrayList) LCount() int  {
	return al.numOfData
}
