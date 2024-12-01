package utils

func Filter[T any](collection []T, predicate func(value T) bool) (filtered []T) {
	for _, elem := range collection {
		if predicate(elem) {
			filtered = append(filtered, elem)
		}
	}
	return
}

func Map[T, R any](collection []T, mapping func(value T) R) (mapped []R) {
	for _, elem := range collection {
		mapped = append(mapped, mapping(elem))
	}
	return
}

func GroupBy[T any, K comparable](values []T, keyFunc func(value T) (key K)) (result map[K][]T) {
	result = make(map[K][]T)
	for _, value := range values {
		key := keyFunc(value)
		result[key] = append(result[key], value)
	}
	return
}

func Identity[T any](v T) T {
	return v
}
