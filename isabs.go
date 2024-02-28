package path

// IsAbs returns if the path is an absoulte path.
//
// An absolute path begins with a "/".
func IsAbs(path string) bool {
	return 1 <= len(path) && '/' == path[0]
}
