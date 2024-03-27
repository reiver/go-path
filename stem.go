package path

// Stem returns the file-name without the extension if it exists, and return an empty string otherwise.
//
// Note that if 'path' == "archive.tar.gz", then the stem would be "archive.tar" (rather than just "archive").
// But the stem of 'path' == "archive.tar" would be "archive".
//
func Stem(path string) string {
	if "" == path {
		return ""
	}

	var name string
	_, name = Split(path)
	if "" == name {
		return ""
	}

	var ext string = Ext(name)
	if "" == ext {
		return name
	}

	return name[:len(name)-len(ext)]
}
