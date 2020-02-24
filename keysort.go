package keysort

import (
	"sort"
)

// String sorts the provided slice given the provided key function.
//
// The function panics if the provided interface is not a slice.
func String(slice interface{}, key func(i int) string) {
	sort.Slice(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}

func StringStable(slice interface{}, key func(i int) string) {
	sort.SliceStable(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}

// StringIsSorted tests whether a slice is sorted.
//
// The function panics if the provided interface is not a slice.
func StringIsSorted(slice interface{}, key func(i int) string) bool {
	return sort.SliceIsSorted(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}

func Int(slice interface{}, key func(i int) int) {
	sort.Slice(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}

func IntStable(slice interface{}, key func(i int) int) {
	sort.SliceStable(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}

func IntIsSorted(slice interface{}, key func(i int) int) bool {
	return sort.SliceIsSorted(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}

func UInt(slice interface{}, key func(i int) uint) {
	sort.Slice(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}

func UIntStable(slice interface{}, key func(i int) uint) {
	sort.SliceStable(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}

func UIntIsSorted(slice interface{}, key func(i int) uint) bool {
	return sort.SliceIsSorted(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}


func Float64(slice interface{}, key func(i int) float64) {
	sort.Slice(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}

func Float64Stable(slice interface{}, key func(i int) float64) {
	sort.SliceStable(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}

func Float64IsSorted(slice interface{}, key func(i int) float64) bool {
	return sort.SliceIsSorted(slice, func(i, j int) bool {
		return key(i) < key(j)
	})
}


