package __99

import (
	"errors"
	"fmt"
)

// 20.有效的括号

type stack struct {
	stack []byte
}

func newStack(size int) *stack {
	return &stack{stack: make([]byte, 0, size)}
}

func (s *stack) push(b byte) {
	s.stack = append(s.stack, b)
}

func (s *stack) pop() (byte, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("empty")
	}

	b := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return b, nil
}

func isValid(s string) bool {

	stack := newStack(len(s))

	for _, v := range s {
		switch v {
		case '{':
			stack.push(byte('}'))
		case '[':
			stack.push(byte(']'))
		case '(':
			stack.push(byte(')'))
		default:
			b, err := stack.pop()
			if err != nil {
				return false
			}

			if b != byte(v) {
				return false
			}
		}
	}

	return !(len(stack.stack) > 0)

}

func main() {

	fmt.Println(isValid(`{[]}`))

}
