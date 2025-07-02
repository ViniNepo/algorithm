package linked_list

import (
	"fmt"
	datastruc "github.com/ViniNepo/algorithm/golang/date_structure/arrays_and_hashing"
)

/*
Reorder Linked List
You are given the head of a singly linked-list.

The positions of a linked list of length = 7 for example, can intially be represented as:

[0, 1, 2, 3, 4, 5, 6]

Reorder the nodes of the linked list to be in the following order:

[0, 6, 1, 5, 2, 4, 3]

Notice that in the general case for a list of length = n the nodes are reordered to be in the following order:

[0, n-1, 1, n-2, 2, n-3, ...]

You may not modify the values in the list's nodes, but instead you must reorder the nodes themselves.

Example 1:

Input: head = [2,4,6,8]

Output: [2,8,4,6]
Example 2:

Input: head = [2,4,6,8,10]

Output: [2,10,4,8,6]
*/

func ReorderList(head *datastruc.ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil
	var prev *datastruc.ListNode
	for second != nil {
		temp := second.Next
		second.Next = prev
		prev = second
		second = temp
	}

	first := head
	second = prev

	for second != nil {
		temp1, temp2 := first.Next, second.Next
		first.Next = second
		second.Next = temp1
		first, second = temp1, temp2
	}

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
