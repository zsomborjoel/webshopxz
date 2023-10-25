package common

import (
	"strings"
)

func GetLastPartUrlPath(url string) string {
	pts := strings.Split(url, "/")
	return pts[len(pts)-1]
}
