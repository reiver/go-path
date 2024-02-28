package path

// Ext returns the file-name extension if it exists, and return an empty string otherwise.
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
