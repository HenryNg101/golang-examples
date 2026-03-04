package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// =========================
	// Basic Numeric Types
	// =========================

	var i int = 42
	var i8 int8 = 8
	var i16 int16 = 16
	var i32 int32 = 32
	var i64 int64 = 64

	var u uint = 42
	var u8 uint8 = 8
	var u16 uint16 = 16
	var u32 uint32 = 32
	var u64 uint64 = 64
	var uintptrVal uintptr = uintptr(unsafe.Pointer(&i))

	var f32 float32 = 3.14
	var f64 float64 = 3.1415926535

	var c64 complex64 = complex(1, 2)
	var c128 complex128 = complex(3, 4)

	fmt.Println(i, i8, i16, i32, i64)
	fmt.Println(u, u8, u16, u32, u64, uintptrVal)
	fmt.Println(f32, f64)
	fmt.Println(c64, c128)

	// =========================
	// Boolean
	// =========================

	var b bool = true
	fmt.Println("bool:", b)

	// =========================
	// String
	// =========================

	var s string = "Hello"
	fmt.Println("string:", s)
	fmt.Println("length:", len(s))

	// =========================
	// Rune & Byte
	// =========================

	var r rune = 'A' // alias for int32
	var by byte = 65 // alias for uint8

	fmt.Println("rune:", r, string(r))
	fmt.Println("byte:", by)

	// =========================
	// Array (Fixed Size)
	// =========================

	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("array:", arr)

	// =========================
	// Slice (Dynamic Array)
	// =========================

	slice := []int{10, 20, 30}
	slice = append(slice, 40)
	fmt.Println("slice:", slice)

	// =========================
	// Map
	// =========================

	m := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	m["Charlie"] = 35
	fmt.Println("map:", m)

	// =========================
	// Struct
	// =========================

	type Person struct {
		Name string
		Age  int
	}

	p := Person{"Alice", 28}
	fmt.Println("struct:", p)

	// =========================
	// Pointer
	// =========================

	x := 100
	ptr := &x
	fmt.Println("pointer value:", *ptr)

	// =========================
	// Function Type
	// =========================

	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("function:", add(2, 3))

	// =========================
	// Interface
	// =========================

	var anyVal interface{} = "I can hold anything"
	fmt.Println("interface:", anyVal)

	// Type assertion
	strVal := anyVal.(string)
	fmt.Println("type assertion:", strVal)

	// =========================
	// Channel
	// =========================

	ch := make(chan int, 1)
	ch <- 10
	fmt.Println("channel:", <-ch)

	// =========================
	// Constants & iota
	// =========================

	const (
		Red = iota
		Green
		Blue
	)
	fmt.Println("iota:", Red, Green, Blue)

	// =========================
	// Type Alias
	// =========================

	type MyInt = int
	var mi MyInt = 50
	fmt.Println("type alias:", mi)

	// =========================
	// Custom Defined Type
	// =========================

	type Celsius float64
	var temp Celsius = 36.6
	fmt.Println("custom type:", temp)

	// =========================
	// Nil-able Types
	// =========================

	var pNil *int
	var sNil []int
	var mNil map[string]int
	var chNil chan int
	var fNil func()

	fmt.Println("nil pointer:", pNil)
	fmt.Println("nil slice:", sNil)
	fmt.Println("nil map:", mNil)
	fmt.Println("nil channel:", chNil)
	fmt.Println("nil func:", fNil)
}
