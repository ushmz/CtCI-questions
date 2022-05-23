package chapter1_test

import (
	"CtCI-questions/chapter1"
	"testing"
)

func TestIsPermutatedBatch(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"Tact coa", true},
		{"Go dog", true},
		{"Boys come back to me", false},
		{"No it never propagates if I set a gap or prevention", true},
	}

	for _, c := range cases {
		actual := chapter1.IsPermutatedPalindrome(c.input)
		if actual != c.expected {
			t.Fatalf("Input %s; Expected: %v, actual: %v\n", c.input, c.expected, actual)
		}
	}
}
