package path

import (
	"strings"
)

// Top returns the top of a path.
func Top(path string) string {

	for 2 <= len(path) && '.' == path[0] && '/' == path[1] {
		path = path[2:]
	}

	switch {
	case "" == path:
		return "."
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
