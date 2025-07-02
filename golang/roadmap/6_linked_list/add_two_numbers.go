package linked_list

import datastruc "github.com/ViniNepo/algorithm/golang/date_structure/arrays_and_hashing"

/*
Add Two Numbers
You are given two non-empty linked lists, l1 and l2, where each represents a non-negative integer.

The digits are stored in reverse order, e.g. the number 123 is represented as 3 -> 2 -> 1 -> in the linked list.

Each of the nodes contains a single digit. You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Return the sum of the two numbers as a linked list.

Example 1:



Input: l1 = [1,2,3], l2 = [4,5,6]

Output: [5,7,9]

Explanation: 321 + 654 = 975.
Example 2:

Input: l1 = [9], l2 = [9]

Output: [8,1]
*/

func AddTwoNumbers(l1 *datastruc.ListNode, l2 *datastruc.ListNode) *datastruc.ListNode {
	dummy := &datastruc.ListNode{}
	curr := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		v1 := 0
		for l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		v2 := 0
		for l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry
		carry = sum / 10
		curr.Next = &datastruc.ListNode{Val: sum % 10}
		curr = curr.Next
	}

	return dummy.Next
}
