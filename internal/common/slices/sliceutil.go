package slices

func IsEmpty[T any](list []T) bool {
	return list == nil || len(list) == 0
}
