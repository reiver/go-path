package path

import (
	"strings"
)

// Parent returns the parent directory of path.
func Parent(path string) string {

	path = Canonical(path)

	switch path {
	case "/":
		return path
	case ".":
		return ".."
	case "..":
		return "../.."
	}

	if strings.HasSuffix(path, "/..") {
		return path + "/.."
	}

	{
		var lastindex int = len(path) - 1

		if '/' == path[lastindex] {
			path = path[:lastindex]
		}
	}

	{
		index := strings.LastIndexByte(path, '/')
		if index < 0 {
			return "."
		}
		if 0 == index {
			return "/"
		}

		var parent string = path[:index+1]

		if "../" == parent || strings.HasSuffix(parent, "/../") {
			parent = parent[:len(parent)-1]
		}

		return parent
	}
}
