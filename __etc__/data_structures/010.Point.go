package main

import "fmt"

type Point struct {
	xPos int
	yPos int
}

type IPoint interface {
	SetPointPos(xPos, yPos int) // Point 변수의 xPos, yPos 값 설정
	PrintPointPos()             // Point 변수의 xPos, yPos 정보 출력
	PointComp(pos2 *Point) int  // 두 Point 변수 비교
}

func (p *Point) SetPointPos(xPos, yPos int) {
	p.xPos = xPos
	p.yPos = yPos
}

func (p *Point) PrintPointPos() {
	fmt.Printf("[%d %d] \n", p.xPos, p.yPos)
}

func (p *Point) PointComp(pos2 *Point) int {
	if p.xPos == pos2.xPos && p.yPos == pos2.yPos {
		return 0
	} else if p.xPos == pos2.xPos {
		return 1
	} else if p.yPos == pos2.yPos {
		return 2
	}
	return -1
}

const LIST_LEN int = 100

type ArrayList struct { // 배열 기반 리스트 구조체
	arr         [LIST_LEN]*Point // 리스트 저장소인 배열
	numOfData   int              // 저장된 데이터의 수
	curPosition int              // 데이터 참조 위치
}

type IArrayList interface {
	ListInit()              // 초기화
	LInsert(data *Point)    // 데이터 저장
	LFirst() (*Point, bool) // 첫 데이터 참조
	LNext() (*Point, bool)  // 두 번째 이후 데이터 참조
	LRemove() int           // 참조한 데이터 삭제
	LCount() int            // 저장된 데이터 수 반환
}

func (al *ArrayList) ListInit() {
	al.numOfData = 0
	al.curPosition = -1
}

func (al *ArrayList) LInsert(data *Point) {
	if al.numOfData > LIST_LEN {
		fmt.Println("Cannot save..")
		return
	}
	al.arr[al.numOfData] = data
	al.numOfData++
}
func (al *ArrayList) LFirst() (*Point, bool) {
	if al.numOfData == 0 {
		return nil, false
	}
	al.curPosition = 0
	return al.arr[al.curPosition], true
}
func (al *ArrayList) LNext() (*Point, bool) {
	if al.curPosition >= al.numOfData-1 {
		return nil, false
	}
	al.curPosition++
	return al.arr[al.curPosition], true
}
func (al *ArrayList) LRemove() *Point {
	var rPos = al.curPosition
	var num = al.numOfData
	var rData = al.arr[rPos]
	for i := rPos; i < num-1; i++ {
		al.arr[i] = al.arr[i+1]
	}
	al.numOfData--
	al.curPosition--
	return rData
}

func (al *ArrayList) LCount() int {
	return al.numOfData
}
