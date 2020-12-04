// * @Author: eshackelford
// * @Date:   2020-11-22 22:28:32
// * @Last Modified by:   eshackelford
// * @Last Modified time: 2020-11-27 17:50:50

//  See also:
// - https://gobyexample.com/arrays
// - https://golang.org/doc/
// - https://tour.golang.org/moretypes/6
// - https://en.wikipedia.org/wiki/Associative_array

// An array in Go is a numbered sequence of elements of a specific length.

package main

import "fmt"

func main() {

	// Here we create an array a that will hold exactly 5 ints.
	// The type of element and length are both part of the array's type.
	// By default, an array is zero-valued (for ints, zero-value = 0).
	var a [5]int
	fmt.Println("emp:", a)

	// We can set a value at an index using the array[index] = value syntax,
	// We can get a value from the array with array[index] syntax.
	a[4] = 100 // 4 in the index is the 5th item since it starts at 0.
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The built-in len returns the length of an array.
	fmt.Println("len:", len(a))

	// One-liner to declare and initialize at the same time.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array types are one-dimensional, but you can compose types to
	// build multi-dimensional data structures:
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Note: when printed with fmt.Println, arrays take on the form [v1 v2 v3 v4 v5 ...]
}
