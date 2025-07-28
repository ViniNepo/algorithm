package datastruc

import (
	"fmt"
)

type SinglyLinkedList struct {
	Head, Tail *ListNode
}

func NewSinglyLinkedList() *SinglyLinkedList {
	node := NewListNode(-1)
	return &SinglyLinkedList{
		Head: node,
		Tail: node,
	}
}

func (s *SinglyLinkedList) InsertEnd(val int) {
	s.Tail.Next = NewListNode(val)
	s.Tail = s.Tail.Next
}

func (s *SinglyLinkedList) Remove(index int) {
	i := 0
	curr := s.Head
	for i < index && curr != nil {
		i++
		curr = curr.Next
	}

	if curr != nil && curr.Next != nil {
		if curr.Next == s.Tail {
			s.Tail = curr
		}
		curr.Next = curr.Next.Next
	}
}

func (s *SinglyLinkedList) Print() {
	curr := s.Head.Next
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}

	fmt.Println()
}
