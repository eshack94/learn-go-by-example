// * @Author: eshackelford
// * @Date:   2020-11-22 22:51:13
// * @Last Modified by:   eshackelford
// * @Last Modified time: 2020-11-30 03:22:50

//  See also:
// - https://gobyexample.com/range
// - https://golang.org/doc/
// - https://tour.golang.org/moretypes/16
// - https://tour.golang.org/moretypes/17

// range iterates over elements in a variety of data structures.
// Here, we will see how range can be used with some of the data
// structures we have already learned.

package main

import "fmt"

func main() {

	// Here we use range to sum the numbers within a slice.
	// Arrays also work this way.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

}
