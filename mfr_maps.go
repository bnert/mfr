package mfr

// Entry is a struct representation of an
// entry in a map.
type Entry[K any, V any] struct {
	Key   K
	Value V
}

// EntriesFromMap returns an array of entries, enabling maps to be used
// with Map/Reduce/Filter/ForEach
func EntriesFromMap[K comparable, V any](m map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], len(m))

	index := 0
	for k, v := range m {
		entries[index] = Entry[K, V]{Key: k, Value: v}
		index += 1
	}

	return entries
}

// EntriesToMap converts an array of entries back into a map
func EntriesToMap[K comparable, V any](entries []Entry[K, V]) map[K]V {
	m := make(map[K]V)
	ForEach(entries, func(ctx Ctx[Entry[K, V]]) {
		m[ctx.Item.Key] = ctx.Item.Value
	})
	return m
}
