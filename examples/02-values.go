// * @Author: eshackelford
// * @Date:   2020-11-15 23:05:57
// * @Last Modified by:   eshackelford
// * @Last Modified time: 2020-11-27 17:44:42

//  See also:
// - https://gobyexample.com/values
// - https://golang.org/doc/
// - https://tour.golang.org/basics/11

// Go has various value types including strings, integers, floats, booleans, etc.

package main

import "fmt"

func main() {

	// Strings can be added together with +
	fmt.Println("go" + "lang")

	// Integers and floats, respectively
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// Booleans, which provide the true and false operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(!false)
	fmt.Println(!true && !false)
}
