// * @Author: eshackelford
// * @Date:   2020-11-16 01:00:37
// * @Last Modified by:   eshackelford
// * @Last Modified time: 2020-11-27 17:45:53

//  See also:
// - https://gobyexample.com/for
// - https://golang.org/doc/
// - https://tour.golang.org/flowcontrol/1

// for is Go's only looping construct. Here are some basic types of for loops.

package main

import "fmt"

func main() {

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// This is a classic initial/condition/after for loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop repeatedly until you break out of the loop
	for {
		fmt.Println("loop")
		break
	}

	// You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
