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
	{2, "2"},
	{7, "Bang"},
	{21, "FizzBang"},
	{11, "Bong"},
	{33, "Bong"},
}

func TestFizzBuzz(t *testing.T) {
	for _, test := range tests {
		result := getOutput(test.input)
		if result != test.expected {
			t.Errorf("incorrect output %d --  got: %s, expected: %s", test.input, result, test.expected)
		}
	}
}
