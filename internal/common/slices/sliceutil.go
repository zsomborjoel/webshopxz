package slices

func IsEmptyList[T any](list []T) bool {
	return list == nil || len(list) == 0
}
