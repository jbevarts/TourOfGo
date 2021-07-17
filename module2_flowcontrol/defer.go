package main

import "fmt"

func main() {
	testDefer()
	testDefer2()
}

func testDefer() {
	// the defer statement is not executed until after
	// the calling frame has returned.
	// it is guaranteed to execute

	// the statement is evaluated immediately
	// but the call to the statement is not executed
	// until after the calling frame returns.

	i := 1

	defer fmt.Println(i)

	i++

	// this statement prints 2, but the defer prints 1
	// because i was evaluated at the time of reading
	// the defer statement
	fmt.Println(i)
}

func testDefer2() {
	// deferred calls are pushed onto a LIFO stack.
	// when the calling frme returns, the defer
	// statements are executed.

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	// should print done first, then print 9 > 0, descending
	// because of LIFO order
	fmt.Println("done")
}
