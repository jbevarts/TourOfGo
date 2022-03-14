package main

import (
	"fmt"
)

// the types of x and y come after the variable name.
func add(x int, y int) int {
	return x + y
}

// two or more consecutive shared type function parameters can omit the type
func add2(x, y int) int {
	return x + y
}

// a function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

// named return values; "naked" returning with no arguments automatically returns the named values
// in declaration.  Only use naked returning in short methods.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(1, 1))
	a, b := swap("world", "hello")
	fmt.Println(a, b)
	fmt.Println(split(17))
}


func addFour (x int, y int, z int, k int) {
	fmt.Println(add(2,2))
}