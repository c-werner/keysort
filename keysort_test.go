package keysort_test

import (
	"github.com/c-werner/keysort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	type Case struct {
		s string
		i int
	}

	var (
		slice = []Case{
			{
				s: "b",
				i: 2,
			},
			{
				s: "a",
				i: 1,
			},
			{
				s: "z",
				i: 4,
			},
			{
				s: "x",
				i: 3,
			},
		}

		byS = func(i int) string {
			return slice[i].s
		}

		expected = []Case{
			{
				s: "a",
				i: 1,
			},
			{
				s: "b",
				i: 2,
			},
			{
				s: "x",
				i: 3,
			},
			{
				s: "z",
				i: 4,
			},
		}
	)

	assert.False(t, keysort.StringIsSorted(slice, byS))
	keysort.String(slice, byS)
	assert.True(t, keysort.StringIsSorted(slice, byS))
	assert.Equal(t, expected, slice)
}

func TestTuple(t *testing.T) {
	type Case struct {
		s string
		i int
	}

	var (
		slice = []Case{
			{
				s: "b",
				i: 2,
			},
			{
				s: "b",
				i: 1,
			},
			{
				s: "z",
				i: 0,
			},
			{
				s: "x",
				i: 3,
			},
		}

		byStringThenInt = func(i int) []interface{} {
			return []interface{}{
				slice[i].s,
				slice[i].i,
			}
		}

		byIntThenString = func(i int) []interface{} {
			return []interface{}{
				slice[i].i,
				slice[i].s,
			}
		}

		strExpected = []Case{
			{
				s: "b",
				i: 1,
			},
			{
				s: "b",
				i: 2,
			},
			{
				s: "x",
				i: 3,
			},
			{
				s: "z",
				i: 0,
			},
		}

		intExpected = []Case{
			{
				s: "z",
				i: 0,
			},
			{
				s: "b",
				i: 1,
			},
			{
				s: "b",
				i: 2,
			},
			{
				s: "x",
				i: 3,
			},
		}
	)

	// Verify the initial slice isn't sorted at all
	assert.False(t, keysort.TupleIsSorted(slice, byStringThenInt))
	assert.False(t, keysort.TupleIsSorted(slice, byIntThenString))

	// Now sort the slice by (s, i) and verify it's appropriately sorted
	keysort.Tuple(slice, byStringThenInt)
	assert.Equal(t, strExpected, slice)
	assert.True(t, keysort.TupleIsSorted(slice, byStringThenInt))
	assert.False(t, keysort.TupleIsSorted(slice, byIntThenString))

	// Now sort the slice by (i, s) and verify it's appropriately sorted
	keysort.Tuple(slice, byIntThenString)
	assert.Equal(t, intExpected, slice)
	assert.True(t, keysort.TupleIsSorted(slice, byIntThenString))
	assert.False(t, keysort.TupleIsSorted(slice, byStringThenInt))
}
