package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var rules = []struct {
	multiple int
	word     string
	active   bool
}{
	{3, "Fizz", false},
	{13, "Fezz", false},
	{5, "Buzz", false},
	{7, "Bang", false},
	{11, "Bong", false},
	{17, "", false},
}

// this makes it easier to test individual outputs
func getOutput(i int) (output string) {
	word := []string{}
	for _, rule := range rules {
		if rule.active == true && i%rule.multiple == 0 {
			word = append(word, rule.word)
			switch rule.multiple {
			case 11:
				word = word[len(word)-1:]
			case 17:
				slices.Reverse(word)
			default:
			}
		}
	}

	if len(word) > 0 {
		return strings.Join(word, "")
	} else {
		return strconv.Itoa(i)
	}
}

func getInteger(message string) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(message)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		if number, error := strconv.Atoi(input); error == nil {
			return number
		}
	}
}

func getYesNo(message string) bool {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(message, " -- y/n: ")
		scanner.Scan()
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if input == "y" {
			return true
		} else if input == "n" {
			return false
		}
	}
}

/*
func getWordRule() (int, string) {
	number := getInteger("Choose a number to modify multiples of: ")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What word would you like to replace it with? ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	return number, input
}
*/

func main() {
	fmt.Println("Hello, World! \n This is FizzBuzz 2.0 -- you may toggle rules on and off.")
	max := getInteger("Enter a max value: ")
	for index, rule := range rules {
		message := "Activate the rule for number" + strconv.Itoa(rule.multiple) + "?"
		reply := getYesNo(message)
		rules[index].active = reply
	}

	// Put your code here...
	for i := 1; i <= max; i++ {
		fmt.Println(getOutput(i))
	}
	//fmt.Println(rules)
}
