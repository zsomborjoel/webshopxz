package maps

func Merge(source map[string]interface{}, target map[string]interface{}) {
	for k, v := range source {
		target[k] = v
	}
}
