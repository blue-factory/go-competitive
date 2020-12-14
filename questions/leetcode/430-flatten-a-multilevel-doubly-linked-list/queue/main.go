import "container/list"

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

	// define queue
	queue := list.New()
	queue.PushBack(root)

	tmpNodes := []*Node{}

	for queue.Len() > 0 {
		// get from node for queue
		node := queue.Front()
		nodeValue := node.Value.(*Node)

		// append value
		tmpNodes = append(tmpNodes, nodeValue)

		// push back child node in queue
		if nodeValue.Child != nil {
			queue.PushBack(nodeValue.Child)

			child := nodeValue.Child
			for child.Next != nil {
				child = child.Next
			}
			child.Next = nodeValue.Next
			if nodeValue.Next != nil {
				nodeValue.Next.Prev = child
			}
			nodeValue.Next = nil
		}

		// push back next node in queue
		if nodeValue.Next != nil {
			queue.PushBack(nodeValue.Next)
		}

		// remove last node
		queue.Remove(node)
	}

	// iterate temporal array and define linked list to response
	head := tmpNodes[0]
	tmp := head
	for i, node := range tmpNodes {
		if i != 0 {
			node.Prev = tmp
		}

		node.Child = nil
		tmp.Next = node
		tmp = tmp.Next
	}

	return head
}