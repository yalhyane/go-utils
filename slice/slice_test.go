package slice_utils

import (
	"reflect"
	"testing"
)

func TestDeleteValueFromSlice(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	v := 3
	expected := []int{1, 2, 4, 5}
	result := DeleteValue(s, v)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v after deleting %v from %v, but got %v", expected, v, s, result)
	}
}

func TestFilterSliceFunc(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(x int) bool { return x%2 == 0 }
	expected := []int{2, 4}
	result := FilterFunc(s, f)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v after filtering with, but got %v", expected, result)
	}
}

func TestFilterSliceFuncWithIndex(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(x int, i int) bool { return i%2 == 0 }
	expected := []int{1, 3, 5}
	result := FilterFuncWithIndex(s, f)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v after filtering with, but got %v", expected, result)
	}
}

func TestSliceClone(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	result := Copy(s)

	if !reflect.DeepEqual(result, s) {
		t.Errorf("Expected a clone of %v, but got %v", s, result)
	}

	if &result == &s {
		t.Errorf("Expected a different slice reference, but got the same reference")
	}
}

func TestSliceRemoveByIndex(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	i := 2
	expected := []int{1, 2, 4, 5}
	result := RemoveByIndex(s, i)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v after removing index %d from %v, but got %v", expected, i, s, result)
	}
}

func TestSliceContains(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	v := 3
	if !Contains(s, v) {
		t.Errorf("Expected %v to be present in %v, but it's not", v, s)
	}

	v = 6
	if Contains(s, v) {
		t.Errorf("Expected %v to be absent in %v, but it's present", v, s)
	}
}

func TestWalkSliceFunc(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	f := func(x int) int { return x * 2 }
	expected := []int{2, 4, 6, 8, 10}
	result := WalkFunc(s, f)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v after applying to %v, but got %v", expected, s, result)
	}
}

func TestIndexOfSliceItem(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	v := 3
	expected := 2
	result := IndexOf(s, v)

	if result != expected {
		t.Errorf("Expected index %d of item %v in %v, but got %d", expected, v, s, result)
	}

	v = 6
	expected = -1
	result = IndexOf(s, v)

	if result != expected {
		t.Errorf("Expected index %d of item %v in %v, but got %d", expected, v, s, result)
	}
}

func TestSliceRandomEntries(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	n := 3
	result := RandomEntries(s, n)

	// all results should belong to s
	filtered := FilterFunc(result, func(i int) bool {
		return Contains(s, i)
	})

	if len(result) != n {
		t.Errorf("Expected %d random entries from %v, but got %v", n, s, result)
	}

	if !reflect.DeepEqual(result, filtered) {
		t.Errorf("Expected all random entries from %v to belong to %v", result, filtered)
	}
	if reflect.DeepEqual(result, s) {
		t.Errorf("Expected random entries, but got the same slice")
	}
}

func TestSliceMapFunc(t *testing.T) {
	type args[T any, NT any] struct {
		s []T
		f func(T) NT
	}
	type testCase[T any, NT any] struct {
		name string
		args args[T, NT]
		want []NT
	}

	type CustomT struct {
		name string
		age  int
	}
	tests := []testCase[CustomT, string]{
		{
			name: "simple test",
			args: args[CustomT, string]{
				s: []CustomT{
					{
						name: "a1",
						age:  12,
					},
					{
						name: "a2",
						age:  14,
					},
				},
				f: func(t CustomT) string {
					return t.name
				},
			},
			want: []string{
				"a1",
				"a2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapFunc(tt.args.s, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
