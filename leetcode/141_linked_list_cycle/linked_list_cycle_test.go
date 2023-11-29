package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	t.Parallel()

	// has a cycle
	a := buildList([]int{3, 2, 0, -4})
	// -4 points to 2
	a.Next.Next.Next = a.Next

	b := buildList([]int{3, 2, 0, -4})
	// -4 points to 2
	b.Next.Next = b.Next

	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{name: "empty", want: false, head: buildList([]int{})},
		{name: "1 item", want: false, head: buildList([]int{1})},
		{name: "2 items", want: false, head: buildList([]int{1, 2})},
		{name: "same values", want: false, head: buildList([]int{1, 1, 1, 1})},
		{name: "with cycle", want: true, head: a},
		{name: "mid-cycle", want: true, head: b},
	}

	for _, tt := range tests {
		result := hasCycle(tt.head)
		assert.Equal(t, tt.want, result)
	}
}
