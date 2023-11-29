package _1_merge_two_lists

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var sortedList []int
	for {
		// nil cases

		if list1 == nil && list2 == nil {
			break
		}

		if list1 == nil {
			sortedList = append(sortedList, list2.Val)
			list2 = list2.Next

			continue
		}

		if list2 == nil {
			sortedList = append(sortedList, list1.Val)
			list1 = list1.Next
			continue
		}

		if list1.Val < list2.Val {
			sortedList = append(sortedList, list1.Val)
			list1 = list1.Next
		} else {
			sortedList = append(sortedList, list2.Val)
			list2 = list2.Next
		}
	}

	return buildList(sortedList)
}

func buildList(nums []int) *ListNode {

	var list *ListNode
	var head *ListNode

	for _, n := range nums {
		if list == nil {
			list = &ListNode{Val: n}
			head = list
			continue
		}
		list.Next = &ListNode{
			Val: n,
		}
		list = list.Next
	}

	return head
}

func CompareLists(List1, List2 *ListNode) bool {

	for {
		if List1 == nil && List2 == nil {
			break
		}

		if List1 == nil || List2 == nil {
			return false
		}

		if List1.Val != List2.Val {
			return false
		}

		List1 = List1.Next
		List2 = List2.Next
	}

	return true
}

func DisplayList(list *ListNode) []int {

	var result []int
	for {
		if list == nil {
			break
		}

		result = append(result, list.Val)
		list = list.Next
	}

	return result
}

// Test merge two lists
func TestMergeTwoLists(t *testing.T) {

	cases := []struct {
		List1, List2 *ListNode
		Output       *ListNode
	}{
		{List1: buildList([]int{1, 2, 4}), List2: buildList([]int{1, 3, 4}), Output: buildList([]int{1, 1, 2, 3, 4, 4})},
		{List1: buildList([]int{}), List2: buildList([]int{}), Output: buildList([]int{})},
		{List1: buildList([]int{}), List2: buildList([]int{2}), Output: buildList([]int{2})},
		{List1: buildList([]int{0, 1, 2}), List2: buildList([]int{-1, 2, 4}), Output: buildList([]int{-1, 0, 1, 2, 2, 4})},
	}

	for _, c := range cases {
		got := mergeTwoLists(c.List1, c.List2)
		assert.True(t, CompareLists(c.Output, got), fmt.Sprintf("expected %v, got %v", DisplayList(c.Output), DisplayList(got)))
	}

}
