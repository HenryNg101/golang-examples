package calculator

import (
	"errors"
)

var ErrDivideByZero = errors.New("cannot divide by zero")

func Add(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func MustPositive(n int) int {
	if n < 0 {
		panic("negative number not allowed")
	}
	return n
}
