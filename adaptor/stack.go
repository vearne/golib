	package adaptor

import (
	"container/list"
)

type Stack struct {
	List *list.List
}

func NewStack() *Stack {
	q := Stack{}
	q.List = list.New()
	return &q
}

func (s *Stack) Empty() bool {
	return s.List.Len() == 0
}

func (s *Stack) Size() int {
	return s.List.Len()
}

func (s *Stack) Top() interface{} {
	return s.List.Back().Value
}

func (s *Stack) Push(v interface{}) {
	s.List.PushBack(v)
}

// from front
func (s *Stack) Pop() interface{} {
	e := s.List.Back()
	s.List.Remove(e)
	return e.Value
}
