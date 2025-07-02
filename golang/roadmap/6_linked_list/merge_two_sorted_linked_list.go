package linked_list

import datastruc "github.com/ViniNepo/algorithm/golang/date_structure/arrays_and_hashing"

/*
Merge Two Sorted Linked Lists
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted linked list and return the head of the new sorted linked list.

The new list should be made up of nodes from list1 and list2.

Example 1:



Input: list1 = [1,2,4], list2 = [1,3,5]

Output: [1,1,2,3,4,5]
Example 2:

Input: list1 = [], list2 = [1,2]

Output: [1,2]
Example 3:

Input: list1 = [], list2 = []

Output: []
*/

func MergeTwoLists(list1 *datastruc.ListNode, list2 *datastruc.ListNode) *datastruc.ListNode {
	dummy := &datastruc.ListNode{}
	node := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Next = list1
			list1 = list1.Next
		} else {
			node.Next = list2
			list2 = list2.Next
		}
		node = node.Next
	}

	node.Next = list1
	if list1 == nil {
		node.Next = list2
	}

	return dummy.Next
}
