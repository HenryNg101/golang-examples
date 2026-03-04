package main

import (
	"errors"
	"fmt"
)

///////////////////////////////////////////////////////////
// Sentinel Error
///////////////////////////////////////////////////////////

var ErrNotFound = errors.New("not found")

///////////////////////////////////////////////////////////
// Custom Error Type
///////////////////////////////////////////////////////////

type ValidationError struct {
	Field string
	Msg   string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", v.Field, v.Msg)
}

///////////////////////////////////////////////////////////
// Custom Error With Unwrap
///////////////////////////////////////////////////////////

type DatabaseError struct {
	Query string
	Err   error
}

func (d DatabaseError) Error() string {
	return fmt.Sprintf("database error on query '%s': %v", d.Query, d.Err)
}

func (d DatabaseError) Unwrap() error {
	return d.Err
}

///////////////////////////////////////////////////////////
// Function Returning Wrapped Errors
///////////////////////////////////////////////////////////

func findUser(id int) error {
	if id == 0 {
		return ErrNotFound
	}
	if id < 0 {
		return ValidationError{
			Field: "id",
			Msg:   "must be positive",
		}
	}
	return nil
}

func queryDatabase(id int) error {
	err := findUser(id)
	if err != nil {
		// wrapping error
		return fmt.Errorf("query failed: %w", err)
	}
	return nil
}

func serviceLayer(id int) error {
	err := queryDatabase(id)
	if err != nil {
		// wrap again
		return DatabaseError{
			Query: "SELECT * FROM users",
			Err:   err,
		}
	}
	return nil
}

///////////////////////////////////////////////////////////
// Custom Is Implementation
///////////////////////////////////////////////////////////

type PermissionError struct {
	Role string
}

func (p PermissionError) Error() string {
	return "permission denied"
}

// Custom comparison logic
func (p PermissionError) Is(target error) bool {
	_, ok := target.(PermissionError)
	return ok
}

///////////////////////////////////////////////////////////
// MAIN
///////////////////////////////////////////////////////////

func main() {

	fmt.Println("=== Sentinel Error ===")
	err := serviceLayer(0)
	fmt.Println("Error:", err)

	// Check sentinel using errors.Is
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Detected sentinel error: not found")
	}

	fmt.Println("\n=== Custom Type + errors.As ===")
	err = serviceLayer(-1)
	fmt.Println("Error:", err)

	var valErr ValidationError
	if errors.As(err, &valErr) {
		fmt.Println("Validation error on field:", valErr.Field)
	}

	fmt.Println("\n=== No Error Case ===")
	err = serviceLayer(10)
	fmt.Println("Error:", err)

	fmt.Println("\n=== Custom Is Implementation ===")
	perr := PermissionError{Role: "guest"}
	wrapped := fmt.Errorf("access failed: %w", perr)

	if errors.Is(wrapped, PermissionError{}) {
		fmt.Println("Permission error detected via custom Is()")
	}

	fmt.Println("\n=== Comparing errors.New ===")
	e1 := errors.New("same message")
	e2 := errors.New("same message")

	fmt.Println("e1 == e2:", e1 == e2)                   // false
	fmt.Println("errors.Is(e1, e2):", errors.Is(e1, e2)) // false
}
