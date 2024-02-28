package path

// CleanSeparators returns path with any extra consecutive separators removed.
func CleanSeparators(path string) string {

	var buffer [256]byte
	var p []byte = buffer[0:0]

	var length int = len(path)

	for i:=0; i<length; i++ {
		var b byte = path[i]

		var next int = i+1

		if '/' == path[i] && next < length && '/' == path[next] {
			continue
		}

		p = append(p, b)
	}

	return string(p)
}
