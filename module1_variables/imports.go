package main

// factored import statements.  Otherwise, would each have a line saying "import x"
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
}
