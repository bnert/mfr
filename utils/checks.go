package utils

import (
	"github.com/bnert/mfr"
)

// Contains returns true if the item is in the array.
func Contains[T comparable](array []T, item T) bool {
	return mfr.Reduce[T, bool](array, false, func(ctx mfr.Ctx[T], acc bool) bool {
		return acc || (ctx.Item == item)
	})
}

// IsAll returns true if all the items in the array match the
// condition function.
func IsAll[T any](array []T, f func(T) bool) bool {
	filtered := mfr.Filter(array, func(ctx mfr.Ctx[T]) bool {
		return f(ctx.Item)
	})
	return len(array) == len(filtered)
}

// IsSome returns true if at least one item in the array matches the
// condition function.
func IsSome[T any](array []T, f func(T) bool) bool {
	filtered := mfr.Filter(array, func(ctx mfr.Ctx[T]) bool {
		return f(ctx.Item)
	})
	return len(filtered) > 0
}

// IsNone returns true if none of the items in the array match
// the condition function.
func IsNone[T any](array []T, f func(T) bool) bool {
	filtered := mfr.Filter(array, func(ctx mfr.Ctx[T]) bool {
		return f(ctx.Item)
	})
	return len(filtered) == 0
}

// Reverse returns the reverse of the original array.
func Reverse[T any](array []T) []T {
	return mfr.Map[T, T](array, func(ctx mfr.Ctx[T]) T {
		return ctx.Array[len(ctx.Array)-1-ctx.Index]
	})
}

// Distinct returns all distinc items in an array
func Distinct[T comparable](array []T) []T {
	if len(array) <= 0 {
		return array
	}

	m := map[T]struct{}{}

	return mfr.Filter(array, func(ctx mfr.Ctx[T]) bool {
		if _, ok := m[ctx.Item]; !ok {
			m[ctx.Item] = struct{}{}
			return true
		} else {
			return false
		}
	})
}

// DistinctBy returns the distinct members of the fn.
func DistinctBy[T any, K comparable](array []T, f func(T) K) []T {
	if len(array) <= 0 {
		return array
	}

	m := map[K]struct{}{}

	return mfr.Filter(array, func(ctx mfr.Ctx[T]) bool {
		k := f(ctx.Item)
		if _, ok := m[k]; !ok {
			m[k] = struct{}{}
			return true
		} else {
			return false
		}
	})
}

// IsDistinct returns true if all items in an array are distinct.
func IsDistinct[T comparable](array []T) bool {
	if len(array) <= 0 {
		return true
	}

	return len(Distinct(array)) == len(array)
}
