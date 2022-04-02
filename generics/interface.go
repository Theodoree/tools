package generics

import (
	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Integer](num T) T {
	if num < 0 {
		return -num
	}
	return num
}

func Add[T constraints.Integer](a, b T) T {
	return a + b
}

func Sub[T constraints.Integer](a, b T) T {
	return a - b
}

func Mut[T constraints.Integer](a, b T) T {
	return a * b
}

func Div[T constraints.Integer](a, b T) T {
	return a / b
}
