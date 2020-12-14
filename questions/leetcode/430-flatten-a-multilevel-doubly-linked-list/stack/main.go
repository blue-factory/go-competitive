package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */
type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	// define stack
	stack := []*Node{root}

	var prev *Node
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack[len(stack)-1] = nil
		stack = stack[:len(stack)-1]

		// push back next node in queue
		if node.Next != nil {
			stack = append(stack, node.Next)
		}

		// push back child node in queue
		if node.Child != nil {
			stack = append(stack, node.Child)
		}

		if prev != nil {
			prev.Next = node
		}

		node.Child = nil
		node.Next = nil
		node.Prev = prev
		prev = node
	}

	return root
}
