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
	if head == nil {
		return head
	}

	if head.Next != nil && head.Next.Next == head {
		return head
	}

	if head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow := head.Next
	fast := head.Next.Next

	for slow != fast {
		if slow.Next == nil || fast.Next == nil || fast.Next.Next == nil {
			return nil
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
