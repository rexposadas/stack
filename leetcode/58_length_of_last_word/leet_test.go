package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func lengthOfLastWord(s string) int {

	s = strings.TrimSpace(s)
	parts := strings.Split(s, " ")
	lastWord := strings.TrimSpace(parts[len(parts)-1])

	return len(lastWord)
}

func TestLengthOfLastWord(t *testing.T) {

	cases := []struct {
		Input  string
		Length int
	}{
		{Input: "Hello World", Length: 5},
		{Input: "   fly me   to   the moon  ", Length: 4},
		{Input: "luffy is still joyboy", Length: 6},
	}

	for _, c := range cases {
		got := lengthOfLastWord(c.Input)
		assert.Equal(t, c.Length, got, "input: %s, expected: %v, got: %v", c.Input, c.Length, got)
	}
}
