package main

import "fmt"

var rules = []struct {
	multiple int
	word     string
}{
	{3, "Fizz"},
	{5, "Buzz"},
}

// This is our main function, this executes by default when we run the main package.
func main() {
	fmt.Println("Hello, World!")

	// Put your code here...

	for i := 1; i <= 100; i++ {
		word := ""
		for _, rule := range rules {
			if i%rule.multiple == 0 {
				word += rule.word
			}
		}
		if word != "" {
			fmt.Println(word)
		} else {
			fmt.Println(i)
		}
	}
}
