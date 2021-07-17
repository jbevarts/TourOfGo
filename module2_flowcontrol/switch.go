package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	// go switches dont have to constants like C
	// go switches automatically break after success
	// don't need to explicitly call break

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X. ")
	case "linus":
		fmt.Println("Linux. ")
	default:
		// freebsd, openbsd
		// plan9, windows....
		fmt.Printf("%s.\n", os)
	}

	// switch cases evaluate from top to bottom
	// proceeding cases do not trigger if success is found
	// can call a function as a parameter of the case.

	fmt.Println("When is Satruday?")
	fmt.Printf("%T %T\n", time.Now(), time.Now().Weekday())

	today := time.Now().Weekday()
	// because Weekday is somewhat of a numerical constant under the hood
	// type Weekday int
	// it cannot be valued over 7, because that's how many days there are

	fmt.Println(today + 1)
	fmt.Println(today - 2)
	fmt.Println(today - 3)

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today. ")
	case today + 1:
		fmt.Println("Tomorrow. ")
	case today + 2:
		fmt.Println("In two days. ")
	default:
		fmt.Println("Too far away. ")
	}

	// a switch statement with no condition is the same
	// as writing switch true {}
	// which ultimately is a long strange of if/else statements
	// succeeding at the first "true" evaluation

	t := time.Now()
	fmt.Printf("hour is %v\n", t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good Evening")

	}
}
