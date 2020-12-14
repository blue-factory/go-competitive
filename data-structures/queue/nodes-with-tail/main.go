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
	top  *Node
	tail *Node
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
	node := &Node{number, nil}

	if s.top == nil {
		s.top = node
		s.tail = node
		return
	}

	s.tail.next = node
	s.tail = s.tail.next
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
