package main

import "fmt"

// we capitalize Vertex, X, and Y if we want
// them to be visible outside of package
// ie. capitals export the values
// it's possible to only export certain underlying
// values of an otherwise exported struct.
// for instance, if x was lower case, only Y
// would be exported.

type Vertex struct {
	X, Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	// struct fields are accessed using
	// dot notation

	// this is a value, not a pointer
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println()

	x := 42
	// this is a pointer, not a value
	y := &Vertex{3, 4}

	// the below shows that ints and struct
	// pointers are printed differently
	// use %p when printing struct pointer
	fmt.Printf("struct pointer value(using v) %v\n", y)
	fmt.Printf("struct pointer value(using p) %p\n", y)
	fmt.Printf("struct value : %v\n", *y)
	fmt.Printf("int pointer (using v) : %v\n", &x)
	fmt.Printf("int pointer (using p) : %p\n", &x)

	// pointers to structs dont need to be
	// dereferenced prior to accessing member
	// (*p).X for example is not needed
	y.X = 7
	fmt.Println(*y)

	// struct literals
	// you can omit the creation of underlying
	// fields by using the "X:" format as shown below.
	// all other fields are set to their zero value.
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)

	fmt.Println(v1, p, v2, v3)

}
