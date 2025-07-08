package main

import (
	"fmt"
	"strconv"
)

var rules = []struct {
	multiple int
	word     string
}{
	{3, "Fizz"},
	{13, "Fezz"},
	{5, "Buzz"},
	{7, "Bang"},
}

// this makes it easier to test individual outputs
func getOutput(i int) (output string) {
	word := ""
	for _, rule := range rules {
		if i%rule.multiple == 0 {
			word += rule.word
		}
	}

	if i%11 == 0 {
		word = "Bong"
	}
	if word != "" {
		return word
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
