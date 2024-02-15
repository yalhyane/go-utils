package math_utils

import (
	"testing"
)

func TestIntMin(t *testing.T) {
	// Define test cases.
	testCases := []struct {
		a, b, expected int
	}{
		{1, 2, 1},
		{-1, 1, -1},
		{10, -10, -10},
	}

	// Run each test case.
	for _, tc := range testCases {
		actual := IntMin(tc.a, tc.b)

		// Assert that the expected output is produced.
		if actual != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, actual)
		}
	}
}

func TestIntMax(t *testing.T) {
	// Define test cases.
	testCases := []struct {
		a, b, expected int
	}{
		{1, 2, 2},
		{-1, 1, 1},
		{10, -10, 10},
	}

	// Run each test case.
	for _, tc := range testCases {
		actual := IntMax(tc.a, tc.b)

		// Assert that the expected output is produced.
		if actual != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, actual)
		}
	}
}

func TestInt64Min(t *testing.T) {
	// Define test cases.
	testCases := []struct {
		a, b, expected int64
	}{
		{1, 2, 1},
		{-1, 1, -1},
		{10, -10, -10},
	}

	// Run each test case.
	for _, tc := range testCases {
		actual := Int64Min(tc.a, tc.b)

		// Assert that the expected output is produced.
		if actual != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, actual)
		}
	}
}

func TestInt64Max(t *testing.T) {
	// Define test cases.
	testCases := []struct {
		a, b, expected int64
	}{
		{1, 2, 2},
		{-1, 1, 1},
		{10, -10, 10},
	}

	// Run each test case.
	for _, tc := range testCases {
		actual := Int64Max(tc.a, tc.b)

		// Assert that the expected output is produced.
		if actual != tc.expected {
			t.Errorf("Expected %d, got %d", tc.expected, actual)
		}
	}
}
