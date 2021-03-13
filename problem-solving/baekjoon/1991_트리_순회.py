# https://www.acmicpc.net/problem/1991
import sys

class Node():
    def __init__(self, name, left, right):
        self.name = name
        self.left = left
        self.right = right
    
def preorder(node):
    print(node.name, end="")
    if node.left != ".":
        preorder(nodes[node.left])
    if node.right != ".":
        preorder(nodes[node.right])

def inorder(node):
    if node.left != ".":
        inorder(nodes[node.left])
    print(node.name, end="")
    if node.right != ".":
        inorder(nodes[node.right])

def postorder(node):
    if node.left != ".":
        postorder(nodes[node.left])
    if node.right != ".":
        postorder(nodes[node.right])
    print(node.name, end="")

nodes = {}

if __name__ == "__main__":
    n = int(sys.stdin.readline())
    for i in range(n):
        node_name, left_name, right_name = sys.stdin.readline().split()
        nodes[node_name] = Node(node_name, left_name, right_name)
    preorder(nodes["A"])
    print()
    inorder(nodes["A"])
    print()
    postorder(nodes["A"])