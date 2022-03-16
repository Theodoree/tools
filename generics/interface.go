package generics

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64
}

func Abs[T Num](num T) T {
	if num < 0 {
		return -num
	}
	return num
}

func Add[T Num](a, b T) T {
	return a + b
}

func Sub[T Num](a, b T) T {
	return a - b
}

func Mut[T Num](a, b T) T {
	return a * b
}

func Div[T Num](a, b T) T {
	return a / b
}
