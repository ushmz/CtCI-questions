package chapter1_test

import (
	"CtCI-questions/chapter1"
	"testing"
)

func TestIsAllUnique(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"abcd", true},
		{"abcc", false},
		{" ", true},
		{"  ", false},
		{"", true},
	}

	for _, c := range cases {
		actual := chapter1.IsAllUnique(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s; Expected: %v, actual: %v\n", c.input, c.expected, actual)
		}
	}
}
