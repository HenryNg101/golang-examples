package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	x := 10
	var p *int
	fmt.Println(p) // nil, because thats default value of pointers

	p = &x

	fmt.Println(x)  // 10
	fmt.Println(p)  // memory address
	fmt.Println(*p) // 10

	*p = 20
	fmt.Println(x) // 20

	// Pointer to struct example
	p2 := Person{"Alice", 25}
	ptr := &p2

	ptr.Age = 30                        // no need (*ptr).Age because of auto dereferencing
	fmt.Println("Alice age is", p2.Age) // 30
}
