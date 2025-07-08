package main

import (
	"testing"
)

var tests = []struct {
	input    int
	expected string
}{
	{3, "Fizz"},
	{5, "Buzz"},
	{11, "11"},
	{7, "Bang"},
	{21, "FizzBang"},
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range tests {
		result := getOutput(test.input)
		if result != test.expected {
			t.Errorf("incorrect output %d --  got: %s, expected: %s", test.input, result, test.expected)
		}
	}
}
