# 원형 연결 리스트의 노드 클래스
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

# 원형 연결 리스트 클래스
class CircularLinkedList:
    def __init__(self):
        self.tail = None

    # This function is only for empty list
    def addToEmpty(self, data):
        if self.tail != None:
            return self.tail
        new_node = Node(data)
        self.tail = new_node

        self.tail.next = self.tail
        return self.tail

    def addBegin(self, data):
        if self.tail == None:
            return self.addToEmpty(data)

        new_node = Node(data)
        new_node.next = self.tail.next
        self.tail.next = new_node

        return self.tail

    def addEnd(self, data):
        if self.tail == None:
            return self.addToEmpty(data)

        new_node = Node(data)
        new_node.next = self.tail.next
        self.tail.next = new_node
        self.tail = new_node

        return self.tail

    def addAfter(self, data, item):
        if self.tail == None:
            return None
        
        new_node = Node(data)
        p = self.tail.next
        while p:
            if p.data == item:
                new_node.next = p.next
                p.next = new_node

                if p == self.tail:
                    self.tail = new_node
                return self.tail
            
            p = p.next
            if p == self.tail.next:
                print(item, "이(가) 리스트에 존재하지 않습니다.")
                break

    def traverse(self):
        if self.tail == None:
            print("리스트가 비어 있습니다.")
            return
        
        node = self.tail.next
        while node:
            print(node.data, end=" ")
            node = node.next
            if node == self.tail.next:
                break

if __name__ == "__main__":
    cll = CircularLinkedList()

    tail = cll.addToEmpty(6)    # 6
    tail = cll.addBegin(4)      # 4 6
    tail = cll.addBegin(2)      # 2 4 6
    tail = cll.addEnd(8)        # 2 4 6 8
    tail = cll.addEnd(12)       # 2 4 6 8 12
    tail = cll.addAfter(10, 8)  # 2 4 6 8 10 12

    cll.traverse()