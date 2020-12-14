package stack

import "container/list"

// Stack ...
type Stack interface {
	Remove() int
	Add(int)
	Peek() int
	IsEmpty() bool
}

// S ...
type S struct {
	list *list.List
}

// New ...
func New() Stack {
	return &S{
		list: list.New(),
	}
}

// Remove ...
func (s *S) Remove() int {
	top := s.list.Front()
	s.list.Remove(top)

	return top.Value.(int)
}

// Add ...
func (s *S) Add(number int) {
	s.list.PushFront(number)
}

// Peek ...
func (s *S) Peek() int {
	top := s.list.Front()

	return top.Value.(int)
}

// IsEmpty ...
func (s *S) IsEmpty() bool {
	if s.list.Len() > 0 {
		return true
	}

	return false
}
