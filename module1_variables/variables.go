package main

import (
	"fmt"
)

// the var statement declares a list of variables; as in function argument lists, the type is last
var c, python, java bool

// the type can be omitted if initializer is present.
var j, k = 1, 2

func main() {
	// a var statement can be at package or function level.  we both here.
	var i int
	fmt.Println(i, c, python, java)

	// if initializer is present, the type can be omitted; the varibles will take the type
	// of the initializer
	var c, python, java = true, false, "no!"
	fmt.Println(j, c, python, java)

	// inside of a function, the := short assignment statement can be used in place of a var
	// declaration with implicit type.
	// outside a function, every statement begins with a keyword (var, func, so on) and so the
	// := construct is not available.
	k := 3
	fmt.Println(k)
}
