package structs.doubleLinkedList.exercises;

public class DesignLinkedList {

    private Node head;
    private Node tail;
    private int length;

    public DesignLinkedList() {
        this.head = null;
        this.tail = null;
        this.length = 0;
    }

    public int get(int index) {
        if (index < 0 || index >= length) {
            return -1;
        }

        Node curr = head;
        if (index < length/2) {
            for (int i = 0; i < index; i++) {
                curr = curr.next;
            }
        } else {
            curr = tail;
            for (int i = length - 1; i > index; i--) {
                curr = curr.prev;
            }
        }

        return curr.value;
    }

    public void addAtHead(int val) {
        Node newNode = new Node(val);
        if (length == 0) {
            head = newNode;
            tail = newNode;
        } else {
            newNode.next = head;
            head.prev = newNode;
            head = newNode;
        }
        length++;
    }

    public void addAtTail(int val) {
        Node newNode = new Node(val);
        if (length == 0) {
            head = newNode;
            tail = newNode;
        } else {
            newNode.prev = tail;
            tail.next = newNode;
            tail = newNode;
        }
        length++;
    }

    public void addAtIndex(int index, int val) {
        if (index == 0) {
            addAtHead(val);
        }

        if (index == length) {
            addAtTail(val);
        }

        if (index >= 0 || index < length) {
            Node newNode = new Node(val);

            Node curr = head;
            while (index < length) {
                curr = curr.next;
            }

            Node before = curr.prev;
            newNode.prev = before;
            newNode.next = curr;
            before.next = newNode;
            curr.prev = newNode;
            length++;
        }
    }

    public void deleteAtIndex(int index) {
        if (index == 0) {
            head = null;
            tail = null;
        }


    }

    public class Node {
        Node next;
        Node prev;
        int value;

        Node(int value) {
            this.value = value;
        }
    }

}
