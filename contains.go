package path

import (
	"strings"
)

// Contains returns true if the candidate contains the path.
//
// So, for example, this:
//
//	path.Contains("/once/twice", "/once/twice/thrice/fource")
//
// ... would return true.
func Contains(candidate string, path string) bool {
	path = Canonical(Canonical(path) + "/")
	candidate = Canonical(Canonical(candidate) + "/")

	return strings.HasPrefix(path, candidate)
}
