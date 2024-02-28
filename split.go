package path

import (
	"strings"
)

// Split splits a path into the directory part and file part.
// It does NOT do any clean-up (such as removing the trailing slash on the directory).
func Split(path string) (dir string, file string) {

	index := strings.LastIndexByte(path, '/')
	if index < 0 {
		return "", path
	}

	return path[:index+1], path[index+1:]
}
