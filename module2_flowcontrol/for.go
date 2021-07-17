package main

import "fmt"

func main() {

	// no paranthesis around anything
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// the init and post statements are optional
	sum1 := 1
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	// in fact, you can drop ; completely
	// essentially, for is a "while"
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// forever condition is possible very compactly
	for {
		break // this is needed here just to avoid the inf loop
	}

}
