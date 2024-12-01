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
