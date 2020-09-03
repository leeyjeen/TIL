# Binary Tree | 이진 트리
```이진 트리```는 각 부모 노드가 **최대 두 개의 자식**을 가질 수 있는 트리 자료 구조이다.

EX. 아래 이미지에서 각 요소는 최대 두 개의 자식을 갖는다.

<img src="./images/17.Binary_Tree.png" width="50%" alt="Binary Tree">
<br><br>

## 이진 트리의 종류
### Full Binary Tree | 정이진 트리
```정이진 트리```는 모든 부모 노드/내부 노드에 **자식 노드가 2개 또는 0개**인 특별한 유형의 이진 트리이다.

<img src="./images/18.Full_Binary_Tree.png" width="50%" alt="Full Binary Tree">
<br><br>

### Perfect Binary Tree | 포화 이진 트리
```포화 이진 트리```는 모든 내부 노드에 정확히 **2개의 자식 노드**가 있고 모든 잎사귀 노드가 **동일한 레벨**에 있는 이진 트리의 유형이다.  

<img src="./images/19.Perfect_Binary_Tree.png" width="70%" alt="Perfect Binary Tree">
<br><br>

### Complete Binary Tree | 완전 이진 트리
```완전 이진 트리```는 정이진 트리와 비슷하지만, 아래와 같은 차이점이 있다.

1. 모든 레벨은 완전히 채워져야 한다.
1. 모든 잎사귀 요소는 **왼쪽으로 기울어져야** 한다.
1. 마지막 잎사귀 요소에는 **오른쪽 형제 노드가 없을 수 있다**. 즉, 완전 이진 트리가 정이진 트리일 필요는 없다.

<img src="./images/20.Complete_Binary_Tree.png" width="70%" alt="Complete Binary Tree">
<br><br>

### Degenerate or Pathological Tree | 변질 트리
```변질 트리```는 왼쪽 또는 오른쪽에 **오직 하나의 자식 노드**를 갖는 트리이다.

<img src="./images/21.Pathological_Tree.png" width="40%" alt="Pathological Tree">
<br><br>

### Skewed Binary Tree | 편향 이진 트리
```편향 이진 트리```는 **왼쪽 또는 오른쪽 노드에 의해 지배**되는 변질 트리이다. 따라서, 편향 이진 트리는 좌편향 이진 트리, 우편향 이진 트리의 두 가지 종류가 있다.

<img src="./images/22.Skewed_Binary_Tree.png" width="80%" alt="Skewed Binary Tree">
<br><br>

### Balanced Binary Tree | 균형 이진 트리
```균형 이진 트리```는 각 노드의 왼쪽과 오른쪽 하위 트리의 **차이가 0 또는 1**인 이진 트리의 종류이다.

<img src="./images/23.Balanced_Binary_Tree.png" width="50%" alt="Balanced Binary Tree">
<br><br>

## 이진 트리 표현
이진 트리의 노드는 다른 노드에 대한 두 개의 포인터와 데이터 부분을 포함하는 구조로 표현된다.

```python
class Node:
    def __init__(self, key):
        self.left = None
        self.right = None
        self.val = key
```


## 코드 예시
```python
# Binary Tree in Python
class Node:
    def __init__(self, key):
        self.left = None
        self.right = None
        self.val = key

    # Traverse preorder
    def traversePreOrder(self):
        print(self.val, end=' ')
        if self.left:
            self.left.traversePreOrder()
        if self.right:
            self.right.traversePreOrder()

    # Traverse inorder
    def traverseInOrder(self):
        if self.left:
            self.left.traverseInOrder()
        print(self.val, end=' ')
        if self.right:
            self.right.traverseInOrder()

    # Traverse postorder
    def traversePostOrder(self):
        if self.left:
            self.left.traversePostOrder()
        if self.right:
            self.right.traversePostOrder()
        print(self.val, end=' ')


root = Node(1)

root.left = Node(2)
root.right = Node(3)

root.left.left = Node(4)

print("Pre order Traversal: ", end="")
root.traversePreOrder()
print("\nIn order Traversal: ", end="")
root.traverseInOrder()
print("\nPost order Traversal: ", end="")
root.traversePostOrder()
```
Output:
```bash
Pre order Traversal: 1 2 4 3 
In order Traversal: 4 2 1 3 
Post order Traversal: 4 2 3 1
```

## 이진 트리의 활용
- 쉽고 빠른 데이터 접근
- 라우터 알고리즘
- 힙 자료 구조 구현
- 구문 트리


## Reference
- https://www.programiz.com/dsa/binary-tree