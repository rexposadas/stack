package main

import (
	"fmt"
	"strings"
	"testing"
)

func isPalindrome(x int) bool {

	s := fmt.Sprintf("%d", x)
	return isPalindromeString(s)
}

func isPalindromeString(s string) bool {

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
		Input  int
		Output bool
	}{
		{Input: 121, Output: true},
		{Input: -12321, Output: false},
		{Input: 0, Output: true},
	}

	for _, c := range cases {
		got := isPalindrome(c.Input)
		if got != c.Output {
			t.Errorf("input: %d, expected: %v, got: %v", c.Input, c.Output, got)
		}
	}
}
