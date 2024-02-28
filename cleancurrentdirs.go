package path

import (
	"strings"
)

// CleanCurrentDirs returns the path with any extra current directory references removed.
func CleanCurrentDirs(path string) string {

	for {
		var before int = len(path)
		path = strings.ReplaceAll(path, "/./", "/")
		var after int = len(path)

		if before <= after {
			break
		}
	}

	{
		var length int = len(path)

		if  2 <= length && '.' == path[0] && '/' == path[1] {
			path = path[2:]
		}
	}

	{
		var length int = len(path)

		if 2 <= length && '/' == path[length-2] && '.' == path[length-1] {
			path = path[:length-1]
		}
	}

	if "" == path {
		return "."
	}

	return path
}
