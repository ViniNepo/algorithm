package structs.linkedList;

public class ListNode {
    public int val;
    public ListNode next;

    // Constructor that sets 'next' to null by default
    public ListNode(int val) {
        this(val, null);
    }

    // Constructor that accepts both value and next node
    public ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }
}