package main

// Every go program is made up of packages.
// programs start running in packages "main"

// these are packages being imported.  The package name is the same as the last element
// of the import path.  math/rand -> package rand

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is, ", rand.Intn(10))
}

// the environment is deterministic, so each time rand.Intn(10) will produce same number
