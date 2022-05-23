package chapter1_test

import (
	"CtCI-questions/chapter1"
	"testing"
)

var (
	cases = []struct {
		inputStr    string
		inputLength int
		expected    string
	}{
		{"Mr John Smith", 13, "Mr%20John%20Smith"},
		{"Hello world", 11, "Hello%20world"},
	}
)

func TestURLify(t *testing.T) {
	for _, c := range cases {
		actual := chapter1.URLify(c.inputStr, c.inputLength)
		if actual != c.expected {
			t.Fatalf("Input %s; Expected: %v, actual: %v\n", c.inputStr, c.expected, actual)
		}
	}
}

func TestURLifyWithBuilder(t *testing.T) {
	for _, c := range cases {
		actual := chapter1.URLifyWithBuilder(c.inputStr, c.inputLength)
		if actual != c.expected {
			t.Fatalf("Input %s; Expected: %v, actual: %v\n", c.inputStr, c.expected, actual)
		}
	}
}
