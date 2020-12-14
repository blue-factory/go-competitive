package main

// https://leetcode.com/problems/linked-list-cycle-ii/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]*ListNode)

	node := head
	for node != nil {
		if _, exists := visited[node]; !exists {
			visited[node] = node
		} else {
			return node
		}

		node = node.Next
	}

	return nil
}
