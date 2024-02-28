package path

import (
	"strings"
)

// Contains returns true fs the candidate contains the path.
func Contains(candidate string, path string) bool {
	path = Canonical(Canonical(path) + "/")
	candidate = Canonical(Canonical(candidate) + "/")

	return strings.HasPrefix(path, candidate)
}
