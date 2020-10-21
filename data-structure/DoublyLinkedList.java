import java.io.*;

public class DoublyLinkedList {
    Node head;

    static class Node {
        int data;
        Node prev;
        Node next;

        Node(int data, Node prev, Node next) {
            this.data = data;
            this.prev = prev;
            this.next = next;
        }
    }

    // 리스트의 맨 앞에 노드를 추가하는 함수
    public static DoublyLinkedList push(DoublyLinkedList list, int newData) {
        // 노드 할당 및 데이터 초기화
        // 새 노드의 next가 현재의 head를 가리키도록 설정
        Node newNode = new Node(newData, null, list.head);

        // head의 prev가 새 노드를 가리키도록 설정
        if (list.head != null) {
            list.head.prev = newNode;
        }

        // head가 새 노드를 가리키도록 설정
        list.head = newNode;
        return list;
    }

    // 이중 연결 리스트의 맨 뒤에 노드를 추가하는 함수
    public static DoublyLinkedList append(DoublyLinkedList list, int newData) {
        Node newNode = new Node(newData, null, null);

        if (list.head == null) {
            newNode.prev = null;
            list.head = newNode;
            return list;
        }

        Node last = list.head;
        while (last.next != null) {
            last = last.next;
        }

        last.next = newNode;
        newNode.prev = last;

        return list;
    }

    // 이중 연결 리스트에서 주어진 위치의 노드를 삭제하는 함수
    public static void deleteNodeAt(DoublyLinkedList list, int index) {
        // 연결 리스트가 비어 있거나 유효하지 않은 위치 값이 주어진 경우
        if (list.head == null || index < 0) {
            return;
        }

        Node curNode = list.head;
        int i = 0;

        // 처음부터 n번째 위치의 노드까지 순회
        while (curNode != null && i < index) {
            i += 1;
            curNode = curNode.next;
        }

        // 이중 연결 리스트의 노드 수보다 index가 큰 경우
        if (curNode == null) {
            return;
        }

        // 삭제하려는 노드가 head일 경우
        if (list.head == curNode) {
            list.head = curNode.next;
        }

        // 삭제할 노드가 마지막 노드가 아닌 경우에만 next를 변경
        if (curNode.next != null) {
            curNode.next.prev = curNode.prev;
        }

        // 삭제할 노드가 첫 번째 노드가 아닌 경우에만 prev를 변경
        if (curNode.prev != null) {
            curNode.prev.next = curNode.next;
        }
    }

    // 이중 연결 리스트의 내용을 출력하는 함수
    public static void printList(DoublyLinkedList list) {
        Node node = list.head;
        while (node != null) {
            System.out.print(node.data + " ");
            node = node.next;
        }
        System.out.println();
    }

    public static void main(String args[]) {
        DoublyLinkedList dll = new DoublyLinkedList();
        dll = append(dll, 6);   // 6
        printList(dll);
        dll = push(dll, 7); // 7<=>6
        printList(dll);
        dll = push(dll, 1); // 1<=>7<=>6
        printList(dll);
        dll = append(dll, 4); // 1<=>7<=>6<=>4
        printList(dll);
        deleteNodeAt(dll, 2); //  1<=>7<=>4
        printList(dll);
    }
}