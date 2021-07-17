package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

// like with for statements, if doesn't need paranthesis
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// like a for loop, the if statement can start with
// a short statement to execute before the condition
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// %g is default %v for floats
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
