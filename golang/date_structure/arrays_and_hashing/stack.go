package datastruc

import "fmt"

type Stack struct {
	stack []any
}

func (s *Stack) Push(n any) {
	s.stack = append(s.stack, n)
}

func (s *Stack) Pop() any {
	if len(s.stack) == 0 {
		fmt.Println("Error: Stack is empty")
		return 0
	}
	res := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return res
}

func (s *Stack) Size() any {
	return len(s.stack)
}
