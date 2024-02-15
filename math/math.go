package math_utils

type UIntNumber interface {
	uint | uint8 | uint16 | uint32 | uint64
}
type IntNumber interface {
	UIntNumber | int | int8 | int16 | int32 | int64
}

type FloatNumber interface {
	float32 | float64
}

type Number interface {
	IntNumber | FloatNumber
}

// Min returns min of two numbers
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}

	return b
}

// Max returns max of two numbers
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}

	return b
}

// IntMin returns min of two numbers
func IntMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// IntMax returns max of two numbers
func IntMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Int64Min returns min of two numbers
func Int64Min(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

// Int64Max returns max of two numbers
func Int64Max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}
