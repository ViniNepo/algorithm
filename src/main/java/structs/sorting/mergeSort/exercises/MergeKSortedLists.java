package structs.sorting.mergeSort.exercises;

import java.util.Arrays;
import java.util.List;

public class MergeKSortedLists {

    public static class ListNode {
        int val;
        ListNode next;
        ListNode() {}
        ListNode(int val) { this.val = val; }
        ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }

    public static void main(String[] args) {
        ListNode n1 = new ListNode(1);
        ListNode n2 = new ListNode(4);
        ListNode n3 = new ListNode(5);
        ListNode n4 = new ListNode(1);
        ListNode n5 = new ListNode(3);
        ListNode n6 = new ListNode(4);
        ListNode n7 = new ListNode(2);
        ListNode n8 = new ListNode(6);

        n1.next = n2;
        n2.next = n3;
        n4.next = n5;
        n5.next = n6;
        n7.next = n8;

        ListNode[] listNodes = new ListNode[]{n1, n4, n7};

//        mergeKLists(listNodes);
    }

//    public static ListNode mergeKLists(ListNode[] lists) {
//
//    }
}
