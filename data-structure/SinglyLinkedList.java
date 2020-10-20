import java.io.*;

// singly linked list
public class SinglyLinkedList {
    Node head;

    // linked list Node.
    // static -> main() can access it.
    static class Node {
        int data;
        Node next;

        Node(int d) {
            data = d;
        }
    }

    // 시작 부분에 새로운 노드를 삽입하는 함수
    public static SinglyLinkedList push(SinglyLinkedList list, int data) {
        Node new_node = new Node(data); // 노드 할당 및 데이터 초기화
        new_node.next = list.head;  // 새 노드의 next가 현재의 head를 가리키도록 설정
        list.head = new_node;   // head가 새 노드를 가리키도록 설정
        return list;
    }

    // 주어진 노드 뒤에 새 노드를 추가하는 함수
    public static SinglyLinkedList insertAfter(SinglyLinkedList list, Node prev_node, int new_data) {
        if (prev_node == null) { // 주어진 노드의 존재 여부 검사
            System.out.println("prev_node는 연결 리스트 내부에 존재하지 않습니다.");
            return list;
        }

        Node new_node = new Node(new_data); // 새 노드 생성 및 데이터 초기화

        new_node.next = prev_node.next; // 새 노드가 이전 노드가 가리키던 것을 가리키도록 설정

        prev_node.next = new_node; // 이전 노드가 새 노드를 가리키도록 설정

        return list;
    }


    // 맨 끝에 노드를 추가하는 함수
    public static SinglyLinkedList append(SinglyLinkedList list, int data) {
        Node new_node = new Node(data); // 새 노드 생성 및 데이터 초기화
        new_node.next = null;

        // 연결리스트가 비어있는 경우, 새 노드를 head로 지정
        if (list.head == null) {
            list.head = new_node;
        } else {
            Node last = list.head;
            while (last.next != null) { // 마지막 노드까지 순회
                last = last.next;
            }
            last.next = new_node; // 마지막 노드 변경
        }
        return list;
    }

    // 연결 리스트에서 data의 첫 번째 항목을 삭제하는 함수
    public static void deleteNode(SinglyLinkedList list, int data) {
        Node temp = list.head; // head 노드 임시 보관
        Node prev = null;

        // head 노드 자체가 삭제할 data를 가지고 있다면
        if (temp != null) {
            if (temp.data == data) {
                list.head = temp.next;
                temp = null;
                return;
            }
        }

        // 삭제할 data를 탐색한다
        // prev.next를 변경하기 위해 이전 노드를 찾는다.
        while (temp != null) {
            if (temp.data == data) {
                break;
            }
            prev = temp;
            temp = temp.next;
        }

        // 연결 리스트에 data가 존재하지 않는다면
        if (temp == null) {
            return;
        }

        // 연결 리스트에서 노드 연결 해제
        prev.next = temp.next;

        temp = null;
    }

    // 리스트의 head로의 reference와 position이 주어졌을 때
    // 주어진 position의 노드를 삭제하는 함수
    public static void deleteNodeAt(SinglyLinkedList list, int position) {
        // 연결 리스트가 비어 있는 경우
        if (list.head == null) {
            return;
        }

        Node temp = list.head; // head 노드 임시 보관

        // 삭제할 노드가 head인 경우
        if (position == 0) {
            list.head = temp.next;
            temp = null;
            return;
        }

        // 삭제할 노드의 이전 노드 찾기
        for (int i = 0; i < position - 1; i++) {
            temp = temp.next;
            if (temp == null) {
                break;
            }
        }

        // position이 노드 개수보다 큰 경우
        if (temp == null) {
            return;
        }
        if (temp.next == null) {
            return;
        }

        // temp.next의 노드가 삭제할 노드인 경우
        // 삭제할 노드의 next를 가리키는 포인터를 보관
        Node next = temp.next.next;

        // 연결 리스트에서 노드 연결 해제
        temp.next = next;
    }

    // 연결 리스트에서 주어진 인덱스 위치의 데이터를 반환하는 함수
    public static int getNth(SinglyLinkedList list, int index) {
        Node curNode = list.head;
        int count = 0;

        while (curNode != null) {
            if (count == index) {
                return curNode.data;
            }
            count += 1;
            curNode = curNode.next;
        }

        return -1; // 존재하지 않는 요소를 찾는 경우 -1 반환
    }

    // 연결 리스트의 내용을 head부터 시작해서 출력하는 함수
    public static void printList(SinglyLinkedList list) {
        Node curNode = list.head;
        System.out.print("SinglyLinkedList: ");

        while (curNode != null) {
            System.out.print(curNode.data + " ");
            curNode = curNode.next;
        }
    }

    public static void main(String []args) {
        SinglyLinkedList list = new SinglyLinkedList();

        list = append(list, 6); // 6->null
        list = push(list, 7); // 7->6->null
        list = push(list, 1); // 1->7->6->null
        list = append(list, 4); // 1->7->6->4->null
        deleteNode(list, 6); // 1->7->4->null
        deleteNodeAt(list, 1); // 1->4->null
        list = insertAfter(list, list.head.next, 8); // 1->4->8->null
        System.out.println("Element at index 2 is : " + getNth(list, 2));

        printList(list);
    }
}