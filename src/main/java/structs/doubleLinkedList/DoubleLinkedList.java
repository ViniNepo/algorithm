package structs.doubleLinkedList;

public class DoubleLinkedList {

    private DoubleListNode head;
    private DoubleListNode tail;

    public DoubleLinkedList() {
        head = new DoubleListNode(-1);
        tail = new DoubleListNode(-1);
        head.next = tail;
        tail.prev = head;
    }

    public void insertFront(int value) {
        DoubleListNode newNode = new DoubleListNode(value);
        newNode.prev = head;
        newNode.next = head.next;

        head.next.prev = newNode;
        head.next = newNode;
    }

    private void insertEnd(int value) {
        DoubleListNode newNode = new DoubleListNode(value);
        newNode.next = tail;
        newNode.prev = tail.prev;

        tail.prev.next = newNode;
        tail.prev = newNode;
    }

    public void removeFront() {
        head.next.next.prev = head;
        head.next = head.next.next;
    }

    public void removeEnd() {
        tail.prev.prev.next = tail;
        tail.prev = tail.prev.prev;
    }

    public void print() {
        DoubleListNode curr = head.next;
        while (curr != tail) {
            System.out.print(curr.value + " -> ");
            curr = curr.next;
        }
        System.out.println();
    }
}


