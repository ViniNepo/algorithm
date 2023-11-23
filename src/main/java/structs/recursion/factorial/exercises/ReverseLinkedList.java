package structs.recursion.factorial.exercises;


import structs.linkedList.LinkedList;
import structs.linkedList.ListNode;

public class ReverseLinkedList {

    public static void main(String[] args) {
        ListNode node1 = new ListNode(1);
        ListNode node2 = new ListNode(2);
        ListNode node3 = new ListNode(3);
        ListNode node4 = new ListNode(4);
        ListNode node5 = new ListNode(5);

        node1.next = node2;
        node2.next = node3;
        node3.next = node4;
        node4.next = node5;

        ListNode n = reverseList(node1);
    }



    public static ListNode reverseList2(ListNode head) {
        ListNode current = head;
        ListNode prev = null;

        while (current != null) {
            ListNode temp = current.next;
            current.next = prev;
            prev = current;
            current = temp;
        }

        return prev;
    }

    public static ListNode reverseList(ListNode head) {
        return rev(head, null);
    }

    public static ListNode rev(ListNode node, ListNode pre) {
        if (node == null) {
            return pre;
        }

        ListNode temp = node.next;
        node.next = pre;

        return rev(temp, node);
    }

}
