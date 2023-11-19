package structs.doubleLinkedList.exercises;

public class DesignLinkedList {

    public static void main(String[] args) {
        DesignLinkedList list = new DesignLinkedList();

        list.addAtHead(7);
        list.addAtHead(2);
        list.addAtHead(1);
        list.addAtIndex(3, 0);
        list.deleteAtIndex(2);
        list.addAtHead(6);
        list.addAtTail(4);
        System.out.println(list.get(4));
        list.addAtHead(4);
        list.addAtIndex(5, 0);
        list.addAtHead(6);

        Node cur = list.head;
        int val = 0;
        while (cur.next != null) {
            System.out.println(list.get(val));

            val = val + 1;
            cur = cur.next;
        }
    }

    Node head;
    Node tail;

    public class Node {
        Node next;
        Node prev;
        int value;

        Node(int value) {
            this.value = value;
        }
    }

    public DesignLinkedList() {
        head = new Node(-1);
        tail = new Node(-1);
        head.next = tail;
        tail.prev = head;
    }

    public int get(int index) {
        Node curr = head.next;

        while (curr != null && index > 0) {
            curr = curr.next;
            index--;
        }

        if (curr != null && curr != tail && index == 0) {
            return curr.value;
        }

        return -1;
    }

    public void addAtHead(int val) {
        Node newNode = new Node(val);

        newNode.next = head.next;
        newNode.prev = head;
        head.next.prev = newNode;
        head.next = newNode;
    }

    public void addAtTail(int val) {
        Node newNode = new Node(val);

        newNode.next = tail;
        newNode.prev = tail.prev;
        tail.prev.next = newNode;
        tail.prev = newNode;
    }

    public void addAtIndex(int index, int val) {
        Node curr = head.next;

        while (curr != null && index > 0) {
            curr = curr.next;
            index--;
        }

        if (curr != null && index == 0) {
            Node newNode = new Node(val);

            newNode.next = curr;
            newNode.prev = curr.prev;
            curr.prev.next = newNode;
            curr.prev = newNode;
        }
    }

    public void deleteAtIndex(int index) {
        Node curr = head.next;

        while (curr != null && index > 0) {
            curr = curr.next;
            index--;
        }

        if (curr != null && curr != tail && index == 0) {
            curr.prev.next = curr.next;
            curr.next.prev = curr.prev;
        }
    }
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * MyLinkedList obj = new MyLinkedList();
 * int param_1 = obj.get(index);
 * obj.addAtHead(val);
 * obj.addAtTail(val);
 * obj.addAtIndex(index,val);
 * obj.deleteAtIndex(index);
 */