package _0_valid_parenthesis

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {

	// matching parenthesis
	parens := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	parts := strings.Split(s, "")

	var stack []string
	for _, part := range parts {
		if len(stack) == 0 {
			stack = append(stack, part)
			continue
		}

		// compare content to last item
		lastItem := stack[len(stack)-1]
		if parens[lastItem] == part {
			// if it's a match, pop
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, part)
	}

	// if we have nothing left then it's valid
	return len(stack) == 0
}

func Test_ValidParenthesis(t *testing.T) {

	cases := []struct {
		Input  string
		Output bool
	}{
		{Input: "()", Output: true},
		{Input: "()[]{}", Output: true},
		{Input: "(]", Output: false},
		{Input: "([]{}{}])", Output: false},
	}

	for _, c := range cases {
		got := isValid(c.Input)
		assert.Equal(t, c.Output, got, fmt.Sprintf("input: %s, expected: %v, got: %v", c.Input, c.Output, got))
	}
}
