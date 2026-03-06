package main

import (
	"fmt"
	"log"
)

func add(a int, b int) int {
	return a + b
}

// Return multiple values, and named return
func divide(a, b int) (result int, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return
	}
	result = a / b
	return
}

// Function as parameter
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Closure. A function (A) returns a function (B), which when it's called (B), it modifies (A)'s local variable
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Variadic functions. The special parameter must be in the end, not before
func accumulate(op byte, nums ...int) int {
	total := 0
	switch op {
	case '*':
		total = 1
		for _, num := range nums {
			total *= num
		}
	case '+':
		for _, num := range nums {
			total += num
		}
	}
	return total
}

func main() {
	fmt.Println("Function as param tests")
	fmt.Println(add(10, 5))
	fmt.Println(operate(1, 5, add))

	fmt.Println("Closure tests")
	c := counter()
	fmt.Println(c())
	fmt.Println(c())

	fmt.Println("Variadic functions tests")
	fmt.Println(accumulate('*', 3, 12, 5, 6))
	fmt.Println(accumulate('+', 3, 12, 5, 6))
	nums := []int{1, 2, 10, 45}
	fmt.Println(accumulate('+', nums...)) // Must use trailing "..."

	// Type checking and printing value of a nil function
	var f func(int) int
	fmt.Printf("%T\n", f)
	fmt.Println(f)

	// Multi-return
	res, err := divide(10, 0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(res)
	}
}
