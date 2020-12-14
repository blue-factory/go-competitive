package main

// https://leetcode.com/problems/linked-list-cycle/

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

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil || head.Next.Next == nil {
		return false
	}

	if head.Next != nil && head.Next.Next == head {
		return true
	}

	slow := head.Next
	fast := head.Next.Next

	for slow != fast {
		if slow.Next == nil || fast.Next == nil || fast.Next.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
