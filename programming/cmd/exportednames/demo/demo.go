package demo

// Exported constant
const Pi = 3.14

// unexported constant
const secret = "hidden"

// Exported variable
var Version = "1.0"

// unexported variable
var internalState = 42

// Exported function
func SayHello() string {
	return "Hello!"
}

// unexported function
func whisper() string {
	return "pssst"
}

// Exported struct
type Person struct {
	Name string // exported field
	age  int    // unexported field
}

// Exported method
func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

// unexported method
func (p Person) secretAge() int {
	return p.age
}
