package main

import (
	"fmt"

	"github.com/HenryNg101/golang-examples/cmd/exportednames/demo"
)

func main() {
	// These can print just fine, because they are functions/variables with capital letter in the beginning -> exported names
	fmt.Println(demo.Pi)
	fmt.Println(demo.Version)
	fmt.Println(demo.SayHello())

	p := demo.Person{Name: "Alice"}
	fmt.Println(p.Name)
	fmt.Println(p.Greet())

	// These will NOT compile:
	// fmt.Println(demo.secret)
	// fmt.Println(demo.internalState)
	// fmt.Println(demo.whisper())
	// fmt.Println(p.age) -> Even though the struct is exported, this specific field isn't
	// fmt.Println(p.secretAge())
}
