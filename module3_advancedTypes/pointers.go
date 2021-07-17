package main

import "fmt"

func main() {
	i, j := 42, 2701

	// * means "consider this an address and get value at address"
	// & means get address of variable

	p := &i         // point to i ( p is address )
	fmt.Println(*p) // read i ( go to address and get value )
	*p = 21         // set i through the pointer
	fmt.Println(i)

	p = &j       // point to j ( p is address of j )
	*p = *p / 37 // set value at p's address
	fmt.Println(j)

	// unlike C, go has no pointer arithmetic
}
