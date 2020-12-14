package stack

// Stack ...
type Stack interface {
	Pop() *Node
	Push(int)
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

// Pop ...
func (s *S) Pop() *Node {
	top := s.top
	s.top = s.top.next

	return top
}

// Push ...
func (s *S) Push(number int) {
	if s.top == nil {
		s.top = &Node{number, nil}
		return
	}

	node := &Node{number, s.top}
	s.top = node
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
