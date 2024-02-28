package path

// Join returns the canonical form of the parts connected together.
func Join(parts ...string) string {

	var buffer [256]byte
	var p []byte = buffer[0:0]

	for _, part := range parts {
		if "" == part {
			continue
		}

		length := len(p)

		if 0 < length && '/' != p[length-1]  {
			p = append(p, '/')
		}
		p = append(p, part...)
	}

	var str string = string(p)

	return Canonical(str)
}
