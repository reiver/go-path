package path

// Ext return the file-name extension if it exists, and return an empty string otherwise.
//
// Some examples:
//
//	"" → ""
//
//	"/" → ""
//
//	"/doc.txt" → ".txt"
//
//	"doc.txt" → ".txt"
//
//	"/image.jpeg" → ".jpeg"
//
//	"image.jpeg" → ".jpeg"
//
//	"/apple/banana/cherry.html" → ".html"
//
//	"apple/banana/cherry.html" → ".html"
//
//	"/path/to/archive.tar.gz" → ".gz"
//
//	"path/to/archive.tar.gz" → ".gz"
//
//	"/once/twice/thrice/fource" → ""
//
//	"once/twice/thrice/fource" → ""
//
//	"/LICENSE" → ""
//
//	"LICENSE" → ""
//
//	"/src/README.md" → ".md"
//
//	"src/README.md" → ".md"
//
//	"/ALLCAPS.HTML" → ".HTML"
//
//	"ALLCAPS.HTML" → ".HTML"
//
//	"/something/name.dir/filename" → ""
//
//	"something/name.dir/filename" → ""
func Ext(path string) string {

	var length int = len(path)

	for i:=length-1; i >= 0; i-- {
		var b byte = path[i]

		if '/' == b {
			return ""
		}
		if '.' == b {
			return path[i:]
		}
	}
	return ""
}
