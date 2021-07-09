package main

import (
	"fmt"
	"math"
)

func main() {
	// exported names have capital letters.  Like math.Pi
	fmt.Println(math.Pi)

	// this would fail
	// fmt.Println(math.pi)

	// Capitals are exported, and any unexported names are not accessible from outside the package.

}
