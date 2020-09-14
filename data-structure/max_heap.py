import sys

class MaxHeap:
    def __init__(self, maxsize):
        self.maxsize = maxsize
        self.size = 0
        self.Heap = [0] * (self.maxsize + 1)
        self.Heap[0] = sys.maxsize
        self.FRONT = 1

    def parent(self, pos):
        return pos // 2

    def leftChild(self, pos):
        return 2 * pos

    def rightChild(self, pos):
        return (2 * pos) + 1

    def isLeaf(self, pos):
        if pos >= (self.size // 2) and pos <= self.size:
            return True
        return False

    def swap(self, fpos, spos):
        self.Heap[fpos], self.Heap[spos] = self.Heap[spos], self.Heap[fpos]

    def maxHeapify(self, pos):
        # 비잎사귀노드이고 자식 노드보다 작은 경우
        if (not self.isLeaf(pos) and 
            (self.Heap[pos] < self.Heap[self.leftChild(pos)] or 
            self.Heap[pos] < self.Heap[self.rightChild(pos)])):
            # 왼쪽 자식 노드와 교환한 후 왼쪽 자식 노드를 heapify한다
            if self.Heap[self.leftChild(pos)] > self.Heap[self.rightChild(pos)]:
                self.swap(pos, self.leftChild(pos))
                self.maxHeapify(self.leftChild(pos))
            # 오른쪽 자식 노드와 교환한 후 오른쪽 자식 노드를 heapify한다
            else: 
                self.swap(pos, self.rightChild(pos))
                self.maxHeapify(self.rightChild(pos))

    def insert(self, element):
        if self.size >= self.maxsize:
            return
        self.size += 1
        self.Heap[self.size] = element

        current = self.size

        while self.Heap[current] > self.Heap[self.parent(current)]:
            self.swap(current, self.parent(current))
            current = self.parent(current)
            
    def Print(self):
        for i in range(1, (self.size // 2) + 1):
            print("Parent : " + str(self.Heap[i]) + " Left child : " + str(self.Heap[2 * i]) + "Right child : " + str(self.Heap[2 * i + 1]))

    def extractMax(self):
        popped = self.Heap[self.FRONT]
        self.Heap[self.FRONT] = self.Heap[self.size]
        self.Heap[self.FRONT] = self.Heap[self.size]
        self.size -= 1
        self.maxHeapify(self.FRONT)
        return popped

if __name__ == "__main__":
    maxHeap = MaxHeap(15) 
    maxHeap.insert(5) 
    maxHeap.insert(3) 
    maxHeap.insert(17) 
    maxHeap.insert(10) 
    maxHeap.insert(84) 
    maxHeap.insert(19) 
    maxHeap.insert(6) 
    maxHeap.insert(22) 
    maxHeap.insert(9) 
  
    maxHeap.Print() 
      
    print("The Max val is " + str(maxHeap.extractMax()))