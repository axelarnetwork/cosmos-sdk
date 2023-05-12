package util

func Filter[T any](source []T, predicate func(T) bool) []T {
	var out []T

	for i := range source {
		if predicate(source[i]) {
			out = append(out, source[i])
		}
	}

	return out
}

func FilterIndex[T any](source []T, predicate func(T) bool) []int {
	var out []int

	for i := range source {
		if predicate(source[i]) {
			out = append(out, i)
		}
	}

	return out
}

func And[T any](predicate ...func(T) bool) func(T) bool {
	return func(t T) bool {
		for i := range predicate {
			if !predicate[i](t) {
				return false
			}
		}

		return true
	}
}

func Map[T, S any](source []T, f func(T) S) []S {
	out := make([]S, len(source))

	for i := range source {
		out[i] = f(source[i])

	}

	return out
}

// ForEach performs the given function on every element of the slice
func ForEach[T any](source []T, f func(T)) {
	for i := range source {
		f(source[i])
	}
}

// Reduce performs a reduction to a single value of the source slice according to the given function
func Reduce[T, S any](source []T, initial *S, f func(current *S, element T) *S) *S {
	v := initial

	for i := range source {
		v = f(v, source[i])
	}

	return v
}

func FirstMatch[T any](source []T, predicate func(T) bool) *T {
	for i := range source {
		if predicate(source[i]) {
			return &source[i]
		}
	}

	return nil
}

func Contains[T comparable](source []T, element T) bool {
	for i := range source {
		if source[i] == element {
			return true
		}
	}

	return false
}
