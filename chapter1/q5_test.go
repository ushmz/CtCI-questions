package chapter1_test

import (
	"CtCI-questions/chapter1"
	"testing"
)

func TestIsOneStep(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bake", false},
		{"", "s", true},
		{"v", "", true},
	}

	for _, c := range cases {
		actual := chapter1.IsOneStep(c.input1, c.input2)
		if actual != c.expected {
			t.Fatalf("Input '%s', '%s'; Expected: %v, actual: %v\n", c.input1, c.input2, c.expected, actual)
		}
	}
}
