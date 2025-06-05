package datastruc

type MinStack struct {
	stack    []int
	minStack []int
}

func NewMinStack() *MinStack {
	return &MinStack{
		stack: make([]int, 0),
	}
}

func (s *MinStack) Push(value int) {
	s.stack = append(s.stack, value)

	minVal := value
	if len(s.minStack) != 0 {
		if s.minStack[len(s.minStack)-1] < value {
			minVal = s.minStack[len(s.minStack)-1]
		}
	}

	s.minStack = append(s.minStack, minVal)
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}
