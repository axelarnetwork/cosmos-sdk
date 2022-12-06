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
