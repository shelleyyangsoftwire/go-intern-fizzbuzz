package main

import (
	"bufio"
	"fmt"
	"os"
	//"regexp"
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

func getMax() int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a max value: ")
		scanner.Scan()
		//numeric := regexp.MustCompile(`[0-9]+$`)
		input := strings.TrimSpace(scanner.Text())
		if number, error := strconv.Atoi(input); error == nil {
			return number
		}
	}
}

// This is our main function, this executes by default when we run the main package.
func main() {
	fmt.Println("Hello, World!")
	max := getMax()
	// Put your code here...
	for i := 1; i <= max; i++ {
		fmt.Println(getOutput(i))
	}
}
