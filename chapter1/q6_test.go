package chapter1_test

import (
	"CtCI-questions/chapter1"
	"testing"
)

func TestCompressText(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"addddssfffssd", "a1d4s2f3s2d1"},
		{"wwwwdassawwdss", "wwwwdassawwdss"},
		{"hjkl", "hjkl"},
	}

	for _, c := range cases {
		actual := chapter1.CompressText(c.input)
		if actual != c.expected {
			t.Fatalf("Input '%s'; Expected: %v, actual: %v\n", c.input, c.expected, actual)
		}
	}
}
