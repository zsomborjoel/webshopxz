package common

func MergeMaps(source map[string]interface{}, target map[string]interface{}) {
	for k, v := range source {
		target[k] = v
	}
}
