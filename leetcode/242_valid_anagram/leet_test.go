package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

// https://leetcode.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	if len(s) == 0 {
		return true
	}

	ss := strings.Split(s, "")
	tt := strings.Split(t, "")

	sort.Strings(ss)
	sort.Strings(tt)

	return strings.Compare(strings.Join(ss, ""), strings.Join(tt, "")) == 0
}

func TestIsAnagram(t *testing.T) {

	cases := []struct {
		S, T  string
		Valid bool
	}{
		{S: "anagram", T: "nagaram", Valid: true},
		{S: "rat", T: "car", Valid: false},
		{S: "", T: "", Valid: true},
		{S: "a", T: "a", Valid: true},
		{S: "a", T: "ab", Valid: false},
	}

	for _, c := range cases {
		got := isAnagram(c.S, c.T)
		if got != c.Valid {
			assert.Equal(t, c.Valid, got, "S: %s, T: %s", c.S, c.T)
		}
	}
}
