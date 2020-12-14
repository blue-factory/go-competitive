package stack

// Stack ...
type Stack interface {
	Remove() *Node
	Add(int)
	Peek() *Node
	IsEmpty() bool
}

// Node ...
type Node struct {
	value int
	next  *Node
}

// S ...
type S struct {
	top *Node
}

// New ...
func New() Stack {
	return &S{
		top: nil,
	}
}

// Remove ...
func (s *S) Remove() *Node {
	top := s.top
	s.top = s.top.next
	return top
}

// Add ...
func (s *S) Add(number int) {
	if s.top == nil {
		s.top = &Node{number, nil}
		return
	}

	root := s.top
	for root.next != nil {
		root = root.next
	}

	root.next = &Node{number, nil}
}

// Peek ...
func (s *S) Peek() *Node {
	return s.top
}

// IsEmpty ...
func (s *S) IsEmpty() bool {
	if s.top == nil {
		return false
	}

	return true
}
