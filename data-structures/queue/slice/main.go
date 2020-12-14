package stack

// Stack ...
type Stack interface {
	Remove() int
	Add(int)
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

// Remove ...
func (s *S) Remove() int {
	top := s.s[0]
	s.s = s.s[1:]

	return top
}

// Add ...
func (s *S) Add(number int) {
	s.s = append(s.s, number)
}

// Peek ...
func (s *S) Peek() int {
	return s.s[0]
}

// IsEmpty ...
func (s *S) IsEmpty() bool {
	if len(s.s) > 0 {
		return true
	}

	return false
}
