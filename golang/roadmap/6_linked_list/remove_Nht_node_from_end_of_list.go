package linked_list

import (
	datastruc "github.com/ViniNepo/algorithm/golang/date_structure/arrays_and_hashing"
)

/*
Remove Node From End of Linked List
You are given the beginning of a linked list head, and an integer n.

Remove the nth node from the end of the list and return the beginning of the list.

Example 1:

Input: head = [1,2,3,4], n = 2

Output: [1,2,4]
Example 2:

Input: head = [5], n = 1

Output: []
Example 3:

Input: head = [1,2], n = 2

Output: [2]
*/

func RemoveNhtFromEnd(head *datastruc.ListNode, n int) *datastruc.ListNode {
	dummy := &datastruc.ListNode{Next: head}
	left := dummy
	right := head

	for n > 0 {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next
	return dummy.Next
}
