// https://www.acmicpc.net/problem/1991
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
	var nodes = map[string]*Node{}
	for i := 0; i < n; i++ {
		var nodeName, leftNodeName, rightNodeName string
		fmt.Fscanln(reader, &nodeName, &leftNodeName, &rightNodeName)
		node := new(Node)
		if nodes[nodeName] != nil {
			node = nodes[nodeName]
		} else {
			node.setName(nodeName)
			nodes[nodeName] = node
		}
		if leftNodeName != "." {
			leftNode := new(Node)
			leftNode.setName(leftNodeName)
			node.setLeftChild(leftNode)
			nodes[leftNodeName] = leftNode
		}
		if rightNodeName != "." {
			rightNode := new(Node)
			rightNode.setName(rightNodeName)
			node.setRightChild(rightNode)
			nodes[rightNodeName] = rightNode
		}
	}
	fmt.Fprintln(writer, preOrder(nodes["A"], ""))
	fmt.Fprintln(writer, inOrder(nodes["A"], ""))
	fmt.Fprintln(writer, postOrder(nodes["A"], ""))
}

type Node struct {
	name      string
	leftNode  *Node
	rightNode *Node
}

func (n *Node) setLeftChild(l *Node) {
	n.leftNode = l
}

func (n *Node) setRightChild(r *Node) {
	n.rightNode = r
}

func (n *Node) setName(name string) {
	n.name = name
}

func preOrder(startNode *Node, printStr string) string {
	printStr += startNode.name
	if startNode.leftNode != nil {
		printStr = preOrder(startNode.leftNode, printStr)
	}
	if startNode.rightNode != nil {
		printStr = preOrder(startNode.rightNode, printStr)
	}
	return printStr
}

func inOrder(startNode *Node, printStr string) string {
	if startNode.leftNode != nil {
		printStr = inOrder(startNode.leftNode, printStr)
	}
	printStr += startNode.name
	if startNode.rightNode != nil {
		printStr = inOrder(startNode.rightNode, printStr)
	}
	return printStr
}

func postOrder(startNode *Node, printStr string) string {
	if startNode.leftNode != nil {
		printStr = postOrder(startNode.leftNode, printStr)
	}
	if startNode.rightNode != nil {
		printStr = postOrder(startNode.rightNode, printStr)
	}
	printStr += startNode.name
	return printStr
}
