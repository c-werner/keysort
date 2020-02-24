package keysort

import (
	"fmt"
	"sort"
)


// Tuple sorts the provided slice given the provided key function
// which returns a tuple (a slice of interface{}).
//
// The sort is not guaranteed to be stable. For a stable sort, use TupleStable.
//
//  Basically, it wraps sort.Slice but uses the key function rather than a less function.
//
// The function panics if the provided interface is not a slice or key function
// returns non-comparable or non-ordered types.
func Tuple(slice interface{}, key func(i int) []interface{}) {
	sort.Slice(slice, func(i, j int) bool {
		return sliceLess(key(i), key(j))
	})
}

// TupleStable does a similar sort to Tuple but uses sort.SliceStable instead.
func TupleStable(slice interface{}, key func(i int) []interface{}) {
	sort.SliceStable(slice, func(i, j int) bool {
		return sliceLess(key(i), key(j))
	})
}

func TupleIsSorted(slice interface{}, key func(i int) []interface{}) bool {
	return sort.SliceIsSorted(slice, func(i, j int) bool {
		return sliceLess(key(i), key(j))
	})
}

func sliceLess(s1, s2 []interface{}) bool {
	if len(s1) != len(s2) {
		panic(fmt.Sprintf("slices provided by key function have mismatched lengths: %d vs %d", len(s1), len(s2)))
	}

	for i := range s1 {
		if interfaceLess(s1[i], s2[i]) {
			return true
		}
		if !interfaceEq(s1[i], s2[i]) {
			return false
		}
	}
	return false
}

// TODO is there a better way of doing this?
func interfaceLess(k1, k2 interface{}) bool {
	switch v := k1.(type) {
	case string:
		v2 := k2.(string)
		return v < v2
	case int:
		v2 := k2.(int)
		return v < v2
	case int8:
		v2 := k2.(int8)
		return v < v2
	case int16:
		v2 := k2.(int16)
		return v < v2
	case int32:
		v2 := k2.(int32)
		return v < v2
	case int64:
		v2 := k2.(int64)
		return v < v2
	case uint:
		v2 := k2.(uint)
		return v < v2
	case uint8:
		v2 := k2.(uint8)
		return v < v2
	case uint16:
		v2 := k2.(uint16)
		return v < v2
	case uint32:
		v2 := k2.(uint32)
		return v < v2
	case uint64:
		v2 := k2.(uint64)
		return v < v2
	case float32:
		v2 := k2.(float32)
		return v < v2
	case float64:
		v2 := k2.(float64)
		return v < v2
	default:
		panic(fmt.Sprintf("unsupported type %T provided by key function", v))
	}
}

// TODO is there a better way of doing this?
func interfaceEq(k1, k2 interface{}) bool {
	switch v := k1.(type) {
	case string:
		v2 := k2.(string)
		return v == v2
	case int:
		v2 := k2.(int)
		return v == v2
	case int8:
		v2 := k2.(int8)
		return v == v2
	case int16:
		v2 := k2.(int16)
		return v == v2
	case int32:
		v2 := k2.(int32)
		return v == v2
	case int64:
		v2 := k2.(int64)
		return v == v2
	case uint:
		v2 := k2.(uint)
		return v == v2
	case uint8:
		v2 := k2.(uint8)
		return v == v2
	case uint16:
		v2 := k2.(uint16)
		return v == v2
	case uint32:
		v2 := k2.(uint32)
		return v == v2
	case uint64:
		v2 := k2.(uint64)
		return v == v2
	case bool:
		v2 := k2.(bool)
		return v == v2
	case float32:
		v2 := k2.(float32)
		return v == v2
	case float64:
		v2 := k2.(float64)
		return v == v2
	default:
		panic(fmt.Sprintf("unsupported type %T provided by key function", v))
	}
}
