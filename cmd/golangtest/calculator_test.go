package calculator

import (
	"errors"
	"testing"
)

///////////////////////////////////////////////////////////
// Basic Test
///////////////////////////////////////////////////////////

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

///////////////////////////////////////////////////////////
// Table-Driven Test
///////////////////////////////////////////////////////////

func TestAddTable(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive numbers", 1, 2, 3},
		{"zeros", 0, 0, 0},
		{"negative numbers", -1, -2, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d,%d) = %d; want %d",
					tt.a, tt.b, got, tt.want)
			}
		})
	}
}

///////////////////////////////////////////////////////////
// Testing Errors
///////////////////////////////////////////////////////////

func TestDivide_Error(t *testing.T) {
	_, err := Divide(10, 0)

	if !errors.Is(err, ErrDivideByZero) {
		t.Errorf("expected ErrDivideByZero, got %v", err)
	}
}

func TestDivide_Success(t *testing.T) {
	result, err := Divide(10, 2)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

///////////////////////////////////////////////////////////
// Testing Panic
///////////////////////////////////////////////////////////

func TestMustPositive_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic but did not panic")
		}
	}()

	MustPositive(-1)
}

func TestMustPositive_NoPanic(t *testing.T) {
	result := MustPositive(5)
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

///////////////////////////////////////////////////////////
// Benchmark
///////////////////////////////////////////////////////////

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}

///////////////////////////////////////////////////////////
// Example Test (Documentation Example)
///////////////////////////////////////////////////////////

func ExampleAdd() {
	result := Add(2, 3)
	println(result)
	// Output: 5
}
