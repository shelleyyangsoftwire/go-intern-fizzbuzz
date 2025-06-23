# Part 1

Write a program that prints the numbers from 1 to 100. If a number is a multiple of three, print "Fizz" instead of the number. If the number is a multiple of five print "Buzz" instead of the number. For numbers which are multiples of both three and five print "FizzBuzz" instead of the number.

Once you're happy with your program, show it to your trainer for some quick feedback.

# Part 2

FizzBuzz is pretty simple as programs go. But it's interesting to see what happens if you try adding new rules. Work through these in order, adding one at a time. How easy is it? How neat and tidy is the resulting code? Can you make changes to your program to make these sorts of enhancements easier, or cleaner?

- If a number is a multiple of 7, print "Bang" instead of the number. For numbers which are multiples of seven and three / five, append Bang to what you'd have printed anyway. (E.g. 3 * 7 = 21: "FizzBang").

- If a number is a multiple of 11, print "Bong" instead of the number. Do not print anything else in these cases. (E.g. 3 * 11 = 33: "Bong")

- If a number is a multiple of 13, print "Fezz" instead of the number. For multiples of most other numbers, the Fezz goes immediately in front of the first thing beginning with B, or at the end if there are none. (E.g. 5 * 13 = 65: "FezzBuzz", 3 * 5 * 13 = 195: "FizzFezzBuzz"). Note that Fezz should be printed even if Bong is also present (E.g. 11 * 13 = 143: "FezzBong")

- If a number is a multiple of 17, reverse the order in which any fizzes, buzzes, bangs etc. are printed. (E.g. 3 * 5 * 17 = 255: "BuzzFizz")

You will obviously need to display more than 100 numbers in order to test out some of these later cases.

How much of a mess has your code become - how can you make it clear what's supposed to be happening in the face of so many rules?

# Stretch Goals

Try these if you finish early or want to challenge yourself in your spare time.

## Prompt the user for a maximum number

Read a value in from the console, then print output up to that number.

## Allow the user to specify command-line options

Let the user pass in which rules to implement (e.g. any combination of 3, 5, 7, 11, 13, 17) as a command line parameter (or via some other means of your choice).

If you wanted to go wild and let the user define their own rules, how would you do that?...

## Write some unit tests

Talk to your trainer for help if you get to this point.
