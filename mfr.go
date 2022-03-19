package mfr

// Ctx is the context passed to Map/Filter/Reduce/ForEach
type Ctx[T any] struct {
	// Array is the
	Array   []T
	Index   int
	IsFirst bool
	IsLast  bool
	Item    T
}

// Filter returns an array of items for which the
// provided function returns true.
func Filter[T any](s []T, fn func(Ctx[T]) bool) []T {
	to := make([]T, 0)
	for i, v := range s {
		ok := fn(Ctx[T]{
			Array:   s,
			Index:   i,
			IsFirst: i == 0,
			IsLast:  i == (len(s) - 1),
			Item:    v,
		})
		if ok {
			to = append(to, v)
		}
	}
	return to
}

// ForEach iterates over each item of the provided array, calling the provided
// function for each item.
func ForEach[T any](s []T, fn func(Ctx[T])) {
	for i, v := range s {
		fn(Ctx[T]{
			Array:   s,
			Index:   i,
			IsFirst: i == 0,
			IsLast:  i == (len(s) - 1),
			Item:    v,
		})
	}
}

// Map maps one type to another, via the transform function provided.
func Map[From any, To any](s []From, fn func(Ctx[From]) To) []To {
	to := make([]To, len(s))
	for i, v := range s {
		to[i] = fn(Ctx[From]{
			Array:   s,
			Index:   i,
			IsFirst: i == 0,
			IsLast:  i == (len(s) - 1),
			Item:    v,
		})
	}
	return to
}

// Reduce reduces and/or collects from an array to a target type, from an
// initial value.
func Reduce[From any, To any](s []From, init To, fn func(Ctx[From], To) To) To {
	acc := init
	for i, v := range s {
		acc = fn(Ctx[From]{
			Array:   s,
			Index:   i,
			IsFirst: i == 0,
			IsLast:  i == (len(s) - 1),
			Item:    v,
		}, acc)
	}
	return acc
}
