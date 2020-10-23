package structures

import "fmt"

// Dynamic Stack

type Stack struct {
	head *Node
}

func (s *Stack) Push(value interface{}) {
	if s.head == nil {
		s.head = &Node{value, nil}
	} else {
		top := &Node{value, s.head}
		s.head = top
	}
}

func (s *Stack) Pop() interface{} {
	value := s.head.value
	s.head = s.head.next
	return value
}

func (s *Stack) Peek() interface{} {
	return s.head.value
}

func (s *Stack) Count() int {
	curr := s.head
	count := 0
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

func (s *Stack) String() string {
	curr := s.head
	output := ""
	for curr != nil {
		output += fmt.Sprintf("%v -> ", curr.value)
		curr = curr.next
	}
	return output
}
