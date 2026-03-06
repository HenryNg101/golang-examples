package main

import (
	"fmt"
)

///////////////////////////////////////////////////////////
// Basic Defer + LIFO
///////////////////////////////////////////////////////////

func deferDemo() {
	fmt.Println("deferDemo start")

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	fmt.Println("deferDemo end")
}

///////////////////////////////////////////////////////////
// Defer Arguments Evaluation Timing
///////////////////////////////////////////////////////////

func deferArgumentDemo() {
	fmt.Println("\ndeferArgumentDemo")

	x := 10

	defer fmt.Println("deferred value:", x)

	x = 20
	fmt.Println("current value:", x)
}

///////////////////////////////////////////////////////////
// Defer + Named Return (modifying return value)
///////////////////////////////////////////////////////////

func namedReturnDemo() (result int) {
	fmt.Println("\nnamedReturnDemo")

	defer func() {
		result += 10
	}()

	result = 5
	return
}

///////////////////////////////////////////////////////////
// Basic Panic
///////////////////////////////////////////////////////////

func panicDemo() {
	fmt.Println("\npanicDemo start")

	panic("something went terribly wrong")

	// never runs
	// fmt.Println("panicDemo end")
}

///////////////////////////////////////////////////////////
// Recover Properly
///////////////////////////////////////////////////////////

func recoverDemo() {
	fmt.Println("\nrecoverDemo start")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("boom!")

	// never runs
	// fmt.Println("recoverDemo end")
}

///////////////////////////////////////////////////////////
// Panic Without Recover (Program Crash)
///////////////////////////////////////////////////////////

func crashDemo() {
	fmt.Println("\ncrashDemo start")
	panic("this will crash the program")
}

///////////////////////////////////////////////////////////
// MAIN
///////////////////////////////////////////////////////////

func main() {

	// Defer LIFO
	deferDemo()

	// Defer argument timing
	deferArgumentDemo()

	// Named return modification
	result := namedReturnDemo()
	fmt.Println("namedReturnDemo result:", result)

	// Recover example
	recoverDemo()

	// Uncomment to see program crash
	// crashDemo()
}
