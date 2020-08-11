package main

type Node struct {
	value int   // data
	next  *Node // Pointer to the next Node in the list
}

type LinkedList struct {
	first  *Node // The first Node
	last   *Node // The last Node
	length int   // The number of nodes
}

func (l *LinkedList) Append(n *Node) {
	if l.length == 0 {
		l.first = n
		l.last = n
		l.length++
		return
	}

	l.last.next = n
	l.last = n
	l.length++
}

func (l *LinkedList) Prepend(n *Node) {
	if l.length == 0 {
		l.first = n
		l.last = n
		l.length++
		return
	}

	firstNode := l.first
	l.first = n
	l.first.next = firstNode
	l.length++
}

func (l *LinkedList) Lookup() []int {
	if l.length == 0 {
		return nil
	}

	var nodes []int
	node := l.first

	for node != nil {
		nodes = append(nodes, node.value)
		node = node.next
	}

	return nodes
}

func (l *LinkedList) Delete(v int) {
	if l.length == 0 {
		return
	}

	if l.first.value == v {
		l.first = l.first.next
		l.length--
		return
	}

	var prevNode *Node
	node := l.first
	for node.value != v {
		if node == nil {
			return
		}
		prevNode = node
		node = node.next
	}

	prevNode.next = node.next
	l.length--
}
