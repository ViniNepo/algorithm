package structs.doubleLinkedList;

public class DoubleListNode {

    int value;
    DoubleListNode next;
    DoubleListNode prev;

    public DoubleListNode(int value) {
        this(value, null, null);
    }

    public DoubleListNode(int value, DoubleListNode next, DoubleListNode prev) {
        this.value = value;
        this.next = next;
        this.prev = prev;
    }

}
