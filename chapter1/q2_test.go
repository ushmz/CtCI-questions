package chapter1_test

import (
	"CtCI-questions/chapter1"
	"testing"
)

func TestIsPermutation(t *testing.T) {
	cases := []struct {
		input    []string
		expected bool
	}{
		{[]string{"asdf", "asdf"}, true},
		{[]string{"asdf", "fdsa"}, true},
		{[]string{"asdf", "wasd"}, false},
		{[]string{"", ""}, true},
	}

	for _, c := range cases {
		actual := chapter1.IsPermutation(c.input[0], c.input[1])
		if actual != c.expected {
			t.Fatalf("Input %v; Expected: %v, actual: %v\n", c.input, c.expected, actual)
		}
	}
}
