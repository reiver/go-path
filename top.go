package path

import (
	"strings"
)

// Top returns the top of a path.
func Top(path string) string {
	var length int = len(path)

	for 2 <= length && '.' == path[0] && '/' == path[1] {
		path = path[2:]
	}

	switch {
	case length <= 0:
		return "."
	case 1 == length:
		return path
	case '/' == path[0]:
		return "/"
	default:
		index := strings.Index(path, "/")
		if index < 0 {
			return path
		}
		return path[:index]
	}
}
