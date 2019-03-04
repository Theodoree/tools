package datatype

type Stack struct {
	node *StackNode
}
type StackNode struct {
	Next  *StackNode
	Value int
}

func NewStack() *Stack {
	return &Stack{}
}

func NewStackNode(value int) *StackNode {
	return &StackNode{Value: value}
}

func (s *Stack) Push(value int) {

	stackNode := NewStackNode(value)
	if s.node == nil {
		s.node = stackNode
		return
	}
	stackNode.Next = s.node
	s.node = stackNode
}

func (s *Stack) Pop() int {
	value := s.node.Value
	s.node = s.node.Next
	return value
}

func (s *Stack) IsEmpty() bool {
	return s.node != nil
}
