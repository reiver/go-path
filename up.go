package path

import (
	"bytes"
)

func up(path *[]byte) {
	if len(*path) <= 0 {
		*path = (*path)[0:0]
		return
	}

	// "." == path
	if 1 == len(*path) && '.' == (*path)[0] {
		*path = (*path)[0:0]
		return
	}

	// "/" == path
	if 1 == len(*path) && '/' == (*path)[0] {
		return
	}

	{
		if '/' == (*path)[len(*path)-1] {
			*path = (*path)[:len(*path)-1]
		}

		var lastindex int = bytes.LastIndexByte((*path), '/')

		switch {
		case lastindex < 0:
			*path = (*path)[0:0]
		default:
			*path = (*path)[:lastindex+1]
		}
	}
}
