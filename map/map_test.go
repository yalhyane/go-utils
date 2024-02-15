package map_utils

import (
	"reflect"
	"sort"
	"testing"
)

func TestCopyMap(t *testing.T) {
	type args[KT comparable, T any] struct {
		m map[KT]T
	}
	type testCase[KT comparable, T any] struct {
		name string
		args args[KT, T]
		want map[KT]T
	}
	tests := []testCase[int, string]{
		{
			name: "Empty map",
			args: args[int, string]{
				m: map[int]string{},
			},
			want: map[int]string{},
		},
		{
			name: "None empty map",
			args: args[int, string]{
				m: map[int]string{
					1: "A",
					2: "B",
				},
			},
			want: map[int]string{
				1: "A",
				2: "B",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Copy(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Copy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMapFunc(t *testing.T) {
	type args[KT comparable, T any] struct {
		m map[KT]T
		f func(KT, T) bool
	}
	type testCase[KT comparable, T any] struct {
		name string
		args args[KT, T]
		want map[KT]T
	}
	tests := []testCase[int, string]{
		{
			name: "Empty map",
			args: args[int, string]{
				m: map[int]string{},
				f: func(k int, v string) bool {
					return true
				},
			},
			want: map[int]string{},
		},
		{
			name: "Even only map",
			args: args[int, string]{
				m: map[int]string{
					1: "A",
					2: "B",
					3: "C",
					4: "D",
				},
				f: func(k int, v string) bool {
					return k%2 == 0
				},
			},
			want: map[int]string{
				2: "B",
				4: "D",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterFunc(tt.args.m, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeMaps(t *testing.T) {
	type args[KT comparable, T any] struct {
		m  map[KT]T
		m2 []map[KT]T
	}
	type testCase[KT comparable, T any] struct {
		name string
		args args[KT, T]
		want map[KT]T
	}
	tests := []testCase[int, string]{
		{
			name: "merge 2 maps",
			args: args[int, string]{
				m: map[int]string{
					1: "A",
					3: "C",
				},
				m2: []map[int]string{
					{
						2: "B",
					},
				},
			},
			want: map[int]string{
				1: "A",
				2: "B",
				3: "C",
			},
		},
		{
			name: "merge multiple maps",
			args: args[int, string]{
				m: map[int]string{
					1: "A",
					3: "C",
				},
				m2: []map[int]string{
					{
						2: "B",
					},
					{
						4: "D",
						5: "E",
					},
					{
						6: "F",
					},
				},
			},
			want: map[int]string{
				1: "A",
				2: "B",
				3: "C",
				4: "D",
				5: "E",
				6: "F",
			},
		},
		{
			name: "later maps should not override first ones",
			args: args[int, string]{
				m: map[int]string{
					1: "A",
					3: "C",
				},
				m2: []map[int]string{
					{
						2: "B",
						3: "D",
					},
					{
						2: "E",
						4: "F",
					},
				},
			},
			want: map[int]string{
				1: "A",
				2: "B",
				3: "C",
				4: "F",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.m, tt.args.m2...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMapValues(t *testing.T) {
	var nilSlice []string
	type args[KT comparable, T any] struct {
		m map[KT]T
	}
	type testCase[KT comparable, T any] struct {
		name string
		args args[KT, T]
		want []T
	}
	tests := []testCase[int, string]{
		{
			name: "empty map",
			args: args[int, string]{
				m: map[int]string{},
			},
			want: nilSlice,
		},
		{
			name: "none empty map",
			args: args[int, string]{
				m: map[int]string{
					1: "A",
					2: "B",
				},
			},
			want: []string{
				"A",
				"B",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetValues(tt.args.m)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMapKeys(t *testing.T) {
	var nilSlice []int
	type args[KT comparable, T any] struct {
		m map[KT]T
	}
	type testCase[KT comparable, T any] struct {
		name string
		args args[KT, T]
		want []KT
	}
	tests := []testCase[int, string]{
		{
			name: "empty map",
			args: args[int, string]{
				m: map[int]string{},
			},
			want: nilSlice,
		},
		{
			name: "none empty map",
			args: args[int, string]{
				m: map[int]string{
					1: "A",
					2: "B",
				},
			},
			want: []int{
				1,
				2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetKeys(tt.args.m)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
