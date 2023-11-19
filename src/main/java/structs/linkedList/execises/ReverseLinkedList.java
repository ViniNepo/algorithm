package structs.linkedList.execises;

import org.w3c.dom.NodeList;
import structs.linkedList.LinkedList;
import structs.linkedList.ListNode;

public class ReverseLinkedList {

    public static void main(String[] args) {

        ListNode node1 = new ListNode(1);
        ListNode node2 = new ListNode(2);
        ListNode node3 = new ListNode(3);
        ListNode node4 = new ListNode(4);

        node1.next = node2;
        node2.next = node3;
        node3.next = node4;

        ListNode node = reverseList2(node1);
        System.out.println(node.val);
    }

    public static ListNode reverseList(ListNode head) {
        ListNode curr = head;
        ListNode prev = null;

        while (curr != null) {
            ListNode temp = curr.next;
            curr.next = prev;
            prev = curr;
            curr = temp;
        }

        return prev;
    }

    public static ListNode reverseList2(ListNode head) {

        if(head == null || head.next == null) {
            return head;
        }

        ListNode revList = reverseList2(head.next);
        ListNode revTail = revList;

        while(revTail.next != null) {
        revTail = revTail.next;
    }

    revTail.next = head;
    head.next = null;

        return revList;
}
}