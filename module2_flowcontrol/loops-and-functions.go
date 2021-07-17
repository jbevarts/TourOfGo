package main

import (
	"fmt"
	"math"
)

func main() {
	var numberToRoot float64 = 15
	answer := Sqrt(numberToRoot)
	fmt.Printf("answer is : %g\n", answer)
	fmt.Printf("difference from built in sqrt function is %g", math.Abs(answer-math.Sqrt(numberToRoot)))

}

// loops over and comes closer
// z -= (z^2 - x)/(2z)
func Sqrt(x float64) float64 {
	var startingGuess float64 = 5.0
	fmt.Printf("Starting guess is %g\n", startingGuess)

	for math.Abs((startingGuess*startingGuess-x)/(2*startingGuess)) > 0.000000001 {
		startingGuess -= (startingGuess*startingGuess - x) / (2 * startingGuess)
		fmt.Printf("Guess has become %g\n", startingGuess)
	}
	return startingGuess
}
