package _51_reverse_words_in_string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// https://leetcode.com/problems/reverse-words-in-a-string/?envType=study-plan-v2&id=leetcode-75
func reverseWords(s string) string {
	parts := strings.Split(strings.TrimSpace(s), " ")

	var result string
	for i := len(parts) - 1; i >= 0; i-- {
		str := strings.TrimSpace(parts[i])

		if len(str) == 0 {
			continue
		}
		if len(result) == 0 {
			result = str
			continue
		}

		result = fmt.Sprintf("%s %s", result, str)
	}

	return result
}

func TestIsAnagram(t *testing.T) {

	cases := []struct {
		in, out string
	}{
		{in: "the sky is blue", out: "blue is sky the"},
		{in: "  hello world  ", out: "world hello"},
		{in: "a good   example", out: "example good a"},
	}

	for _, c := range cases {
		got := reverseWords(c.in)
		assert.Equalf(t, c.out, got, "Expected: %s, Got: %s", c.out, got)
	}
}
