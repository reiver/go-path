package path

import (
	"strings"
)

// Split splits a path into the directory part and file part.
// It does NOT do any clean-up (such as removing the trailing slash on the directory).
//
// Note that whether there is a trailing separator (i.e., "/") character or not has an effect on the result.
//
// For example:
//
//	"/once/twice/thrice/fource"  -> dir: "/once/twice/thrice/",        file: "fource"
//
//	"/once/twice/thrice/fource/" -> dir: "/once/twice/thrice/fource/", file: ""
func Split(path string) (dir string, file string) {

	index := strings.LastIndexByte(path, '/')
	if index < 0 {
		return "", path
	}

	return path[:index+1], path[index+1:]
}
