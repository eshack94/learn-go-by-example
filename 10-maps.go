// * @Author: eshackelford
// * @Date:   2020-11-22 22:51:13
// * @Last Modified by:   Elijah Shackelford
// * @Last Modified time: 2020-12-04 17:15:36

//  See also:
// - https://gobyexample.com/maps
// - https://golang.org/doc/
// - https://tour.golang.org/moretypes/19
// - https://golangdocs.com/blank-identifier-in-golang

// Maps are Go's built-in associative data type (hash maps/dicts in other languages).

package main

import "fmt"

func main() {

	// To create an empty map, use the make built-in:
	// e.g., make(map[key-type]val-type)
	m := make(map[string]int)

	// Key/value pairs can then be set using the typical [key] = value syntax.
	m["key1"] = 7
	m["key2"] = 13

	// Printing a map with fmt.Println will show all of its key-value pairs
	// in the form mapname[key:value key:value] sorted alphabetically by key name.
	fmt.Println(m)

	// One can get a value for a key with the mapname[keyname] syntax.
	v1 := m["key1"]
	fmt.Println("v1: ", v1)

	// The built-in len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "key2")
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates
	// if the key was present in the map. This can be used to disambiguate between
	// missing keys and keys with zero values like 0 or "". Here we didn't need the
	// value itself, so we ignored it with the blank identifier _.

	// We expect this to return false since we just deleted key2 above.
	_, key2_present := m["key2"]
	fmt.Println("present:", key2_present)

	// We expect this to return true since we didn't modify this key.
	_, key1_present := m["key1"]
	fmt.Println("present:", key1_present)

	// The following syntax will declare AND initialize a new map in the same line.
	n := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	fmt.Println("map:", n)
}
