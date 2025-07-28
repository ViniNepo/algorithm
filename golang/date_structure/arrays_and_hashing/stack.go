package datastruc

import (
	"errors"
	"fmt"
)

type Stack struct {
	stack []int
}

func (s *Stack) Push(n int) {
	s.stack = append(s.stack, n)
}

func (s *Stack) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("Error: Stack is empty")

	}
	res := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return res, nil
}

func (s *Stack) Size() int {
	return len(s.stack)
}

func test() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	value, err := stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	value, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	fmt.Println(stack.Size())

	value, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	value, err = stack.Pop()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}
