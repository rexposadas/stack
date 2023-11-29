package _06_reversed_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_buildList(t *testing.T) {

	a := buildList([]int{1})
	assert.Equal(t, 1, a.Val)
	assert.Nil(t, a.Next)
	assert.Equal(t, 1, a.Count())

	b := buildList([]int{1, 2})
	assert.Equal(t, 1, b.Val)
	assert.NotNil(t, b.Next)
	assert.Equal(t, 2, b.Count())

	secondItem := b.Next
	assert.Equal(t, 2, secondItem.Val)
	assert.Nil(t, secondItem.Next)

}

func Test_reverseList(t *testing.T) {

	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{name: "1 item", head: buildList([]int{1}), want: buildList([]int{1})},
		{name: "2 items", head: buildList([]int{1, 2}), want: buildList([]int{2, 1})},
		{name: "3 items", head: buildList([]int{1, 2, 3}), want: buildList([]int{3, 2, 1})},
	}
	for _, tt := range tests {
		result := reverseList(tt.head)
		assert.Equal(t, result.Count(), tt.want.Count())
	}
}

func TestListNode_Count(t *testing.T) {
	one := buildList([]int{1})
	two := buildList([]int{1, 2})

	tests := []struct {
		name string
		list *ListNode
		want int
	}{
		{name: "1", list: one, want: 1},
		{name: "2", list: two, want: 2},
	}
	for _, tt := range tests {
		count := tt.list.Count()
		assert.Equal(t, tt.want, count)
	}
}
