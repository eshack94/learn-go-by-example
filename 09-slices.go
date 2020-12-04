// * @Author: eshackelford
// * @Date:   2020-11-22 23:03:36
// * @Last Modified by:   eshackelford
// * @Last Modified time: 2020-11-27 18:33:21

//  See also:
// - https://gobyexample.com/slices
// - https://golang.org/doc/
// - https://tour.golang.org/moretypes/7
// - https://blog.golang.org/slices-intro

// Slices are typically used far more often than arrays in Go.

// Slices are a key data type in Go, giving a more powerful interface
// to sequence than arrays.

package main

import "fmt"

func main() {

	//  Unlike arrays, slices are typed only by the elements they contain (not
	// by the number of elements). To create an empty slice with non-zero length,
	// use the built-in make.

	// Here, we make a slice of strings of length 3 (initially zero-valued)

	s := make([]string, 3)
	fmt.Println("emp:", s)

	// We can set and get elements from slices in the same way we can from arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len returns the length of the slice, as it does with arrays.
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices support several additional
	// operations that make them more versatile than arrays.

	// One of these additional operations is the built-in append, which returns a slice
	// containing one or more new values.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copied using the copy built-in.
	// Below, we create an empty slice c of the same length as s, and then
	// copy the values from s into c.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy", c)

	// Slices also support the "slice" operator. The "slice" operator takes
	// the syntax 'slice[low:high]'. In the example below, this gets a slice of
	// the elements s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// This slices up from (and including) s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	// Append the letter 'j' to our declared string slice
	t = append(t, "j")
	fmt.Println("dcl_apd:", t)

	// Slices can be composed into multi-dimensional data structures.
	// Unlike multi-dimensional arrays, the length of the inner slices can vary.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Although slices are different types than arrays, they are rendered
	// similarly by fmt.Println.
}
