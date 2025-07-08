package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var rules = []struct {
	multiple int
	word     string
}{
	{3, "Fizz"},
	{13, "Fezz"},
	{5, "Buzz"},
	{7, "Bang"},
	{11, "Bong"},
}

// this makes it easier to test individual outputs
func getOutput(i int) (output string) {
	word := []string{}
	for _, rule := range rules {
		if i%rule.multiple == 0 {
			word = append(word, rule.word)
		}
	}

	if i%11 == 0 {
		word = word[len(word)-1:]
	}

	if i%17 == 0 {
		slices.Reverse(word)
	}
	if len(word) > 0 {
		return strings.Join(word, "")
	} else {
		return strconv.Itoa(i)
	}
}

// This is our main function, this executes by default when we run the main package.
func main() {
	fmt.Println("Hello, World!")

	// Put your code here...
	for i := 1; i <= 100; i++ {
		fmt.Println(getOutput(i))
	}
}
