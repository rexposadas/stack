package leetcode

import "fmt"

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

// https://leetcode.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	// only the key is important since that will be where we keep the pointer value.
	m := map[string]int{}

	steps := head
	for {
		if steps.Next == nil {
			return false
		}

		// Looks like this: i.e. 0xc0000592a0
		// A unique value per node.
		pos := fmt.Sprintf("%p", steps)

		if _, ok := m[pos]; ok { // If we came here before, then we have a cycle
			return true
		}
		m[pos] = steps.Val // Keep track of the pointer value, which is unique.

		// go to the next node
		steps = steps.Next
	}
}
