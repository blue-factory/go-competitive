package main

// https://www.hackerrank.com/challenges/merge-two-sorted-linked-lists/problem

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     int data;
 *     SinglyLinkedListNode next;
 * }
 *
 */
function mergeLists(head1, head2) {
	if (head1 === null && head2 === null) {
		return null
	}
	if (head1 === null) {
		return head2
	}
	if (head2 === null) {
		return head1
	}

	let head
	let tail

	let node1 = head1
	let node2 = head2

	if (head1.data < head2.data) {
		head = node1
		node1 = node1.next
		tail = head
	} else {
		head = node2
		node2 = node2.next
		tail = head
	}

	while (node1 != null && node2 != null) {
		if (node1.data < node2.data) {
			tail.next = node1
			tail = tail.next

			node1 = node1.next
		} else {
			tail.next = node2
			tail = tail.next

			node2 = node2.next
		}
	}

	while (node1 != null) {
		tail.next = node1
		tail = tail.next

		node1 = node1.next
	}

	while (node2 != null) {
		tail.next = node2
		tail = tail.next

		node2 = node2.next
	}

	return head
}