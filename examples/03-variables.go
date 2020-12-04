// * @Author: eshackelford
// * @Date:   2020-11-15 23:28:02
// * @Last Modified by:   eshackelford
// * @Last Modified time: 2020-11-27 17:44:22

//  See also:
// - https://gobyexample.com/variables
// - https://golang.org/doc/
// - https://tour.golang.org/basics/8

// Variables in Go are explicitly declared and used by the compiler to
// check the type-correctness of function calls, etc.

package main

import "fmt"

func main() {

	// var declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type for initialized variables.
	var d = true
	fmt.Println(d)

	//  Variables declared without a corresponding initialization
	// are zero-valued. For example, the zero-value for an int is 0.
	var e int
	fmt.Println(e)

	var x = ""
	fmt.Println(x)

	//  The := syntax is shorthand for declaring and initializing a variable.
	// For example, the example below is interpreted as 'var f string = "apple"'
	f := "apple"
	fmt.Println(f)

	f2 := 10
	fmt.Println(f2)

	f3 := true
	fmt.Println(f3)
}
