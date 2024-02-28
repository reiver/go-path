package path

import (
	"strings"
)

// Top returns the top of a path.
//
// Some examples:
//
//	"" → ""
//
//	"/" → "/"
//
//	"/doc.txt" → "/"
//
//	"doc.txt" → "doc.txt"
//
//	"/image.jpeg" → "/"
//
//	"image.jpeg" → "image.jpeg"
//
//	"/apple/banana/cherry.html" → "/"
//
//	"apple/banana/cherry.html" → "apple"
//
//	"/path/to/archive.tar.gz" → "/"
//
//	"path/to/archive.tar.gz" → "path"
//
//	"/once/twice/thrice/fource" → "/"
//
//	"once/twice/thrice/fource" → "once"
//
//	"/LICENSE" -> ""
//
//	"LICENSE" -> "LICENSE"
//
//	"/src/README.md" -> "/"
//
//	"src/README.md" -> "src"
//
//	"/ALLCAPS.HTML" -> "/"
//
//	"ALLCAPS.HTML" -> "ALLCAPS.HTML"
//
//	"/something/name.dir/filename" → "/"
//
//	"something/name.dir/filename" → "something"
func Top(path string) string {
	var length int = len(path)

	switch {
	case length <= 0:
		return ""
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
