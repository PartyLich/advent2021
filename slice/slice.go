// Package slice defines various functions useful with slices of any type.
// Unless otherwise specified, these functions all apply to the elements of a
// slice at index 0 <= i < len(s).
package slice

// Compare returns true if b contains all items in a, in the same order. Thus it
// will also return true if b contains additional items beyond the last index of
// a
func Compare[T comparable](a []T, b []T) bool {
	for i, v := range a {
		if i >= len(b) || v != b[i] {
			return false
		}
	}

	return true
}

// Map applies function fn to all items in list, returning a new slice
func Map[T any, O any](fn func(T) O) func(list []T) []O {
	return func(list []T) []O {
		out := make([]O, len(list))

		for i, t := range list {
			out[i] = fn(t)
		}

		return out
	}
}

// Reduce executes the supplied iterating function on each element of the array,
// in order, passing in the element and the return value from the calculation on
// the preceding element.
//
// The first time that the iterating function is called there is no “return
// value of the previous calculation”, the initial value is used in its place.
func Reduce[T any, O any](fn func(O, T) O) func(list []T, init O) O {
	return func(list []T, init O) O {
		out := init

		for _, t := range list {
			out = fn(out, t)
		}

		return out
	}
}

// Fold executes the supplied iterating function on each element of the array,
// in order, passing in the element and the return value from the calculation on
// the preceding element.
//
// The first time that the iterating function is called there is no “return
// value of the previous calculation”, the initial value is used in its place.
func Fold[T any, O any](fn func(O, T) O) func(list []T, init O) O {
	return Reduce(fn)
}

// ReduceRight executes the supplied iterating function on each element of the
// array, from end to start, passing in the element and the return value from
// the calculation on the preceding element.
//
// The first time that the iterating function is called there is no “return
// value of the previous calculation”, the initial value is used in its place.
func ReduceRight[T any, O any](fn func(O, T) O) func(list []T, init O) O {
	return func(list []T, init O) O {
		out := init

		for i := len(list) - 1; i >= 0; i-- {
			out = fn(out, list[i])
		}

		return out
	}
}

func Filter[T any](pred func(T) bool, list []T) []T {
	result := make([]T, 0, len(list))

	for _, v := range list {
		if pred(v) {
			result = append(result, v)
		}
	}

	return result
}
