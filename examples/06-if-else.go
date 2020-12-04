// * @Author: eshackelford
// * @Date:   2020-11-17 11:30:14
// * @Last Modified by:   eshackelford
// * @Last Modified time: 2020-11-27 17:46:13

//  See also:
// - https://gobyexample.com/if-else
// - https://golang.org/doc/
// - https://tour.golang.org/flowcontrol/7

// Branching with if and else in Go is straight-forward.

package main

import "fmt"

func main() {

	// Here is a basic example.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can have an if statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals; any variables declared in this
	// statement are available in all branches of the conditional logic.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Note that you don't need parentheses around conditions in Go, but
// the braces are required.

// There is no ternary if operator (?:) in Go, so you'll need to use a full
// if statement even for basic conditions.
