package _25_valid_palindrome

import (
	"regexp"
	"strings"
	"testing"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z\d]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func isPalindrome(s string) bool {

	s = clearString(s)

	// empty string is a palindrome
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	// convert to lowercase
	s = strings.ToLower(s)

	// split into parts
	parts := strings.Split(s, "")
	for i := 0; i < len(parts)/2; i++ {
		left := parts[i]
		right := parts[len(parts)-1-i]
		if left != right {
			return false
		}
	}

	return true
}

func TestIsPalindrome(t *testing.T) {

	cases := []struct {
		Input  string
		Output bool
	}{
		{Input: "A man, a plan, a canal: Panama", Output: true},
		{Input: "race a car", Output: false},
		{Input: " ", Output: true},
	}

	for _, c := range cases {
		got := isPalindrome(c.Input)
		if got != c.Output {
			t.Errorf("input: %s, expected: %v, got: %v", c.Input, c.Output, got)
		}
	}
}
