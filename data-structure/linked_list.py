# 노드 클래스
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

# 연결 리스트 클래스
class LinkedList:
    def __init__(self):
        self.head = None

    # 연결 리스트의 내용을 Head부터 시작해서 출력하는 함수
    def printList(self):
        temp = self.head
        while temp:
            print(temp.data)
            temp = temp.next
            
if __name__ == "__main__":
    linked_list = LinkedList()
    
    linked_list.head = Node(1)
    second = Node(2)
    third = Node(3)

    linked_list.head.next = second
    second.next = third

    linked_list.printList()
