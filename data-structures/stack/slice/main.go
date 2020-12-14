package stack

// Stack ...
type Stack interface {
	Pop() int
	Push(int)
	Peek() int
	IsEmpty() bool
}

// S ..
type S struct {
	s []int
}

// New ...
func New() Stack {
	return &S{
		s: []int{},
	}
}

// Pop ...
func (s *S) Pop() int {
	top := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]

	return top
}

// Push ...
func (s *S) Push(number int) {
	s.s = append(s.s, number)
}

// Peek ...
func (s *S) Peek() int {
	return s.s[len(s.s)-1]
}

// IsEmpty ...
func (s *S) IsEmpty() bool {
	if len(s.s) > 0 {
		return true
	}

	return false
}
