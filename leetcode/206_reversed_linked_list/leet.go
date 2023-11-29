package _06_reversed_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// Count .
// With the current node as head, count the number of nodes
// including the head.
func (n *ListNode) Count() int {
	return Count(n)
}

// Count .
// With the current node as head, count the number of nodes
// including the head.
func Count(a *ListNode) int {
	if a == nil {
		return 0
	}

	return 1 + Count(a.Next)
}

func buildList(list []int) *ListNode {
	var head *ListNode
	var traverse *ListNode
	for i, v := range list {
		if i == 0 {
			head = &ListNode{Val: v}
			traverse = head
			continue
		}
		newNode := &ListNode{Val: v}
		traverse.Next = newNode
		traverse = newNode
	}

	return head
}

func M(head, tail *ListNode) *ListNode {
	if tail == nil { // no more nodes to move
		return head
	}

	remainder := tail.Next
	tail.Next = nil // remove any trailing nodes before making it head
	tail.Next = head
	return M(tail, remainder)
}

// https://leetcode.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	newHead := &ListNode{
		Val: head.Val,
	}

	return M(newHead, head.Next)
}
