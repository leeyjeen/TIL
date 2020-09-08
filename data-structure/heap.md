# Heap | 힙
`힙 자료 구조`는 **힙 속성**을 충족하는 **완전 이진 트리**(Complete Binary Tree)이다. `이진 힙`이라고도 한다.

완전 이진 트리는 다음과 같은 특별한 이진 트리이다.
- 마지막 레벨을 제외한 모든 레벨이 채워진다.
- 모든 노드가 가능한 한 왼쪽으로 채워진다.

힙 속성은 다음과 같은 노드의 속성을 갖는 것을 의미한다.
- 최대 힙(max heap): 각 노드의 데이터가 항상 자식 노드보다 크며 루트 노드의 데이터는 모든 노드 중 가장 크다.
<br>
<img src="./images/34.Max-heap.png" width="50%" alt="Max-heap">
<br><br>

- 최소 힙(min heap): 각 노드의 데이터가 항상 자식 노드보다 작으며 루트 노드의 데이터는 모든 노드 중 가장 작다.
<br>
<img src="./images/35.Min-heap.png" width="50%" alt="Min-heap">
<br><br>

## 힙 연산
힙에서 수행되는 중요한 연산 중 일부를 아래에서 알고리즘과 함께 설명한다.

### Heapify
`Heapify`는 이진 트리로부터 힙 자료 구조를 생성하는 과정이다. 최소 힙 또는 최대 힙을 생성하기 위해 사용된다.

1. 입력 배열을 다음과 같이 둔다.
<br>
<img src="./images/36.heap_initial_array.png" width="50%" alt="heap initial array">
<br><br>

2. 배열에서 완전 이진 트리를 생성한다.
<br>
<img src="./images/37.Complete_binary_tree.png" width="50%" alt="Complete binary tree">
<br><br>

3. 인덱스가 `n/2 - 1`인 비잎사귀 노드의 첫 번째 인덱스로 부터 시작해보자.
<br>
<img src="./images/38.heapify.png" width="50%" alt="heapify">
<br><br>

4. 현재 요소 `i`를 `largest`로 설정한다.
5. 왼쪽 자식의 인덱스는 `2i + 1`, 오른쪽 자식의 인덱스는 `2i + 2`이다.
`leftChild`가 `currentElement`보다 크다면 (i.e. `ith` 인덱스의 요소), `leftChildIndex`를 `largest`로 설정한다.
`rightChild`가 `largest`의 요소보다 크다면, `rightChildIndex`를 `largest`로 설정한다.
6. `largest`와 `currentElement`를 교환한다.
<br>
<img src="./images/39.heapify.png" width="50%" alt="heapify">
<br><br>

7. 하위 트리들이 heapify될 때까지 3-7 단계를 반복한다. 

### 알고리즘
To be continued..


## Reference
- https://www.programiz.com/dsa/heap-data-structure