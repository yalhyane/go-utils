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

// Min returns the minimum value between the two given numbers
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}

	return b
}

// Max returns the greater value between the two given numbers.
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}

	return b
}

// IntMin returns the minimum value between the two given int numbers
func IntMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// IntMax returns the greater value between the two given int numbers.
func IntMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Int64Min returns the minimum value between the two given int64 numbers.
func Int64Min(a, b int64) int64 {
	if a < b {
		return a
	}

	return b
}

// Int64Max returns the greater value between the two given int64 numbers.
func Int64Max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}
