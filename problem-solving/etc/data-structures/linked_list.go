package main

// Node
// 연결리스트의 각 노드는 두 부분으로 구성된다.
// 1) 데이터 2) 다음 노드를 가리키는 포인터
type Node struct {
	data int
	next *Node
}

// LinkedList
// 연결리스트는 연결리스트의 첫번째 노드를 가리키는 포인터에 의해 표현된다. 
// 첫번째 노드를 head라고 부른다.
// 연결리스트가 비어있다면 head값은 null이다.
type LinkedList struct {
	head *Node
	size int
}

// InsertAt adds a new node at the given node index
func InsertAt(index, data int) error {
	if index < 0 || index > ll.size {
		return fmt.Errorf("Index out of bounds")
	}
	newNode := Node{
		data: data,
		next: nil,
	}
	if index == 0 {
		newNode.next = ll.head
		ll.head = &newNode
		return nil
	}
	node := ll.head
	var insertIndex = 0
	for insertIndex := 0; insertIndex<index-1; insertIndex++ {
		node = node.next
	}
	newNode.next = node.next
	node.next = &newNode
	ll.size++
	return nil
}

// Append adds a new node at the end of the linked list
func (ll *LinkedList) Append(data int) {
	node := Node{
		data: head, 
		next: nil,
	}
	if ll.head == nil {
		ll.head = &node
	} else {
		last := ll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	ll.size++
}

// RemoveAt removes a node at the index of the linked list
func (ll *LinkedList) RemoveAt(index int) (*int, error) {
	if index < 0 || index > ll.size {
		return nil, fmt.Errorf("Index out of bounds")
	}
	node := ll.head
	for i:=0; i<index-1; i++ {
		node = node.next
	}
	removeNode := node.next
	node.next = removeNode.next
	ll.size--
	return &removeNode.data, nil
}

// GetIndexOf returns the index of the data
func (ll *LinkedList) GetIndexOf(data int) int {
	node := ll.head
	index := 0
	for  {
		if node.data == data {
			return index
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		index++
	}
}

// IsEmpty checks if the list is empty
func (ll *LinkedList) IsEmpty() bool {
	if ll.head == nil {
		return true
	}
	return false
}

// GetSize get size of the list
func (ll *LinkedList) GetSize() int {
	node := ll.head
	size := 0
	for {
		if node == nil {
			break
		}
		size++
		node = node.next
	}
	return size
}