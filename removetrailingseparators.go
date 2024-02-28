package path

// RemoveTrailingSeparators returns the path with any trailing separator characters removed.
func RemoveTrailingSeparators(path string) string {
	if "" == path  {
		return ""
	}

	var length int = len(path)

	if 1 == length {
		return path
	}

	for  {
		length = len(path)
		if length <= 1 {
			break
		}
		if '/' != path[length-1] {
			break
		}

		path = path[:length-1]
	}

	return path
}
