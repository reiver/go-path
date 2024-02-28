package path

import (
	"strings"
)

// Canonical returns path in canonical form.
func Canonical(path string) string {
	path = CleanSeparators(path)
	path = CleanCurrentDirs(path)

	switch path {
	case "":
		return "."

	case "/":
		return path

	case ".":
		return path
	case "./":
		return "."

	case "..":
		return path
	case "../":
		return ".."
	}

	if !strings.Contains(path, "/..") {
		return path
	}

	var p []byte = make([]byte, 0, len(path))

	for 0 < len(path) {
		var top string = Top(path)
		if len(top) <= 0 {
			break
		}

		path = path[len(top):]
		var choppedSeparator bool
		if "/" != top && 0 < len(path) {
			path = path[1:]
			choppedSeparator = true
		}

		switch top {
		case "..":
			var length int = len(p)

			switch {
			// "" == p
			case length <= 0:
				p = append(p, ".."...)
			// "/" == p
			case 1 == length && '/' == p[0]:
				// Nothing here.
			// ".." == p
			case 2 == length && '.' == p[0] && '.' == p[1]:
				p = append(p, "/.."...)
			// HasSuffix(p, "/..")
			case 3 <= length && '/' == p[length-3] && '.' == p[length-2] && '.' == p[length-1]:
				p = append(p, "/.."...)
			default:
				up(&p)
			}
		default:
			if 0 < len(p) && !(1 == len(p) && '/' == p[0]) {
				p = append(p, '/')
			}
			p = append(p, top...)
			if "" == path && choppedSeparator {
				p = append(p, '/')
			}
		}

	}

	return string(p)
}
