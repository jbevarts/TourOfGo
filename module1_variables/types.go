package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Go's basic types are
// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
//      // represents a Unicode code point

// float32 float64

// complex64 complex128

//tldr; use int unless you need unsigned or sized

// variables can be blocked into one statement as below

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Numeric Constants
// these are high-precision values
// an untyped constant takes the type needed by its context!!!!
// try printing needInt(Big).
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// variables declard without an explicit initial value are given their "zero" value
	// zero strings are ""
	var i int
	var f float32
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// types can be casted, and must be done explicitly
	var x, y int = 3, 4
	var p float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(p)
	fmt.Println(x, y, z)

	// type inference
	// when the right-hand side of expression is typed, the new variable
	// is of the same type
	var m int
	n := m // this means now n is also a var int
	// however, untyped numeric constants are ambiguous with their type
	ab := 42
	bc := 3.142
	g := 0.867 + 0.5i
	fmt.Printf("%T %T %T %T \n", ab, bc, g, n)

	// constants
	// constants are declared like variables, but with the const keyword
	const World = "WORLD" // constants are characters, strings, booleans, or numeric types
	// cannot be declared using := syntax.

	// Numeric Constants
	// these are high-precision values
	// an untyped constant takes the type needed by its context!!!!
	// try printing needInt(Big).
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}
