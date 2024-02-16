package slice_utils

import "math/rand"

// DeleteValue removes a specified value from a given slice.
func DeleteValue[T comparable](s []T, v T) []T {
	return FilterFunc(s, func(t T) bool {
		return t != v
	})
}

// FilterFunc allows filtering of a given slice with a given filter function.
func FilterFunc[T any](s []T, f func(T) bool) []T {
	return FilterFuncWithIndex(s, func(t T, i int) bool {
		return f(t)
	})
}

// FilterFuncWithIndex allows for a slice to be filtered based on index and value.
func FilterFuncWithIndex[T any](s []T, f func(T, int) bool) []T {
	var n []T
	for i, e := range s {
		if f(e, i) {
			n = append(n, e)
		}
	}

	return n
}

// MapFunc allows for the efficient mapping of a slice according to a given function.
func MapFunc[T, NT any](s []T, f func(T) NT) []NT {
	n := make([]NT, 0, len(s))
	for _, e := range s {
		n = append(n, f(e))
	}
	return n
}

// Copy makes and returns a copy of slice
func Copy[T any](s []T) []T {
	s1 := make([]T, len(s), cap(s))
	copy(s1, s)
	return s1
}

// RemoveByIndex remove a given index from slice. if index is out of range a clone of index will be returned
func RemoveByIndex[T any](s []T, i int) []T {
	if i >= len(s) {
		return Copy(s)
	}
	return append(s[:i], s[i+1:]...)
}

func Contains[T comparable](s []T, v T) bool {
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}

// WalkFunc can create a new slice by applying a function to each element of a given slice.
func WalkFunc[T any](s []T, f func(T) T) []T {
	res := make([]T, 0, len(s))
	for _, v := range s {
		res = append(res, f(v))
	}
	return res
}

// IndexOf returns the index of the entry in the given slice that matches the given item.
func IndexOf[T comparable](s []T, v T) int {
	return IndexOfFunc(s, func(t T, i int) bool {
		return t == v
	})
}

// IndexOfFunc finds the index of a given item in a slice.
func IndexOfFunc[T any](s []T, f func(T, int) bool) int {
	for i, e := range s {
		if f(e, i) {
			return i
		}
	}
	return -1
}

// RandomEntries randomly selects entries from a given slice and returns a new slice of the selections.
func RandomEntries[T any](s []T, n int) []T {
	l := len(s)
	var ts []T
	if n >= l {
		ts = append(ts, s...)
		return ts
	}
	keys := rand.Perm(l)

	for _, k := range keys[:n] {
		ts = append(ts, s[k])
	}

	return ts

}
